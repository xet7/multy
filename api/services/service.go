package services

import (
	"context"
	"fmt"
	"github.com/multycloud/multy/api/errors"
	"github.com/multycloud/multy/api/proto/commonpb"
	pberr "github.com/multycloud/multy/api/proto/errorspb"
	"github.com/multycloud/multy/api/proto/resourcespb"
	"github.com/multycloud/multy/api/service_context"
	"github.com/multycloud/multy/api/util"
	"github.com/multycloud/multy/db"
	"github.com/multycloud/multy/resources"
	"github.com/multycloud/multy/resources/types/metadata"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"log"
)

type CreateRequest[Arg proto.Message] interface {
	GetResource() Arg
	proto.Message
}

type UpdateRequest[Arg proto.Message] interface {
	GetResource() Arg
	WithResourceId
}

type WithResourceId interface {
	GetResourceId() string
	proto.Message
}

type Service[Arg proto.Message, OutT proto.Message] struct {
	ServiceContext *service_context.ResourceServiceContext
	ResourceName   string
}

func NewService[Arg proto.Message, OutT proto.Message](resourceName string, db *service_context.ResourceServiceContext) Service[Arg, OutT] {
	return Service[Arg, OutT]{ResourceName: resourceName, ServiceContext: db}
}

func (s Service[Arg, OutT]) updateErrorMetric(err error, method string) {
	if err != nil {
		go s.ServiceContext.AwsClient.UpdateErrorMetric(s.ResourceName, method, errors.ErrorCode(err))
	}
}

func (s Service[Arg, OutT]) Create(ctx context.Context, in CreateRequest[Arg]) (out OutT, err error) {
	defer func() { s.updateErrorMetric(err, "create") }()
	return errors.WrappingErrors(s.create)(ctx, in)
}

func (s Service[Arg, OutT]) create(ctx context.Context, in CreateRequest[Arg]) (out OutT, err error) {
	key, err := util.ExtractApiKey(ctx)
	if err != nil {
		return
	}

	userId, err := s.ServiceContext.GetUserId(ctx, key)
	if err != nil {
		return
	}
	configPrefix := getConfigPrefixForCreateReq(in.GetResource(), userId)

	go s.ServiceContext.AwsClient.UpdateQPSMetric(userId, s.ResourceName, "create")
	log.Printf("[INFO] user: %s. create %s", userId, s.ResourceName)

	lock, err := s.ServiceContext.LockConfig(ctx, userId, configPrefix)
	if err != nil {
		return
	}
	defer s.ServiceContext.UnlockConfig(ctx, lock)

	c, err := s.getConfig(ctx, userId, lock, configPrefix)
	if err != nil {
		return
	}

	resource, err := c.CreateResource(in.GetResource())
	if err != nil {
		return
	}

	defer func() {
		if err == nil {
			err = s.saveConfig(ctx, c, lock, configPrefix)
		} else {
			log.Println("[DEBUG] Something went wrong, not storing state")
		}
	}()

	log.Printf("[INFO] Deploying %s\n", resource.GetResourceId())
	rollbackFn, err := s.ServiceContext.DeploymentExecutor.Deploy(ctx, c, nil, resource, configPrefix)
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			rollbackFn()
		}
	}()

	return s.readFromConfig(ctx, c, &resourcespb.ReadVirtualNetworkRequest{ResourceId: resource.GetResourceId()}, configPrefix)
}

func (s Service[Arg, OutT]) getConfig(ctx context.Context, userId string, lock *db.ConfigLock, configPrefix string) (*resources.MultyConfig, error) {
	c, err := s.ServiceContext.LoadUserConfig(ctx, userId, configPrefix, lock)
	if err != nil {
		return nil, err
	}
	mconfig, err := resources.LoadConfig(c, metadata.Metadatas)
	if err != nil {
		return nil, err
	}
	return mconfig, err
}

func (s Service[Arg, OutT]) saveConfig(ctx context.Context, c *resources.MultyConfig, lock *db.ConfigLock, configPrefix string) error {
	exportedConfig, err := c.ExportConfig()
	if err != nil {
		return err
	}

	return s.ServiceContext.StoreUserConfig(ctx, exportedConfig, configPrefix, lock)
}

func (s Service[Arg, OutT]) Read(ctx context.Context, in WithResourceId) (out OutT, err error) {
	defer func() { s.updateErrorMetric(err, "read") }()
	return errors.WrappingErrors(s.read)(ctx, in)
}

func (s Service[Arg, OutT]) read(ctx context.Context, in WithResourceId) (OutT, error) {
	key, err := util.ExtractApiKey(ctx)
	if err != nil {
		return *new(OutT), err
	}

	userId, err := s.ServiceContext.GetUserId(ctx, key)
	if err != nil {
		return *new(OutT), err
	}
	configPrefix := GetConfigPrefix(in, userId)

	go s.ServiceContext.AwsClient.UpdateQPSMetric(userId, s.ResourceName, "read")
	log.Printf("[INFO] user: %s. read %s %s", userId, s.ResourceName, in.GetResourceId())

	lock, err := s.ServiceContext.LockConfig(ctx, userId, configPrefix)
	if err != nil {
		return *new(OutT), err
	}
	defer s.ServiceContext.UnlockConfig(ctx, lock)

	c, err := s.getConfig(ctx, userId, lock, configPrefix)
	if err != nil {
		return *new(OutT), err
	}

	_, err = s.ServiceContext.DeploymentExecutor.EncodeAndStoreTfFile(ctx, c, nil, nil, configPrefix)
	if err != nil {
		return *new(OutT), err
	}

	return s.readFromConfig(ctx, c, in, configPrefix)
}

func (s Service[Arg, OutT]) readFromConfig(ctx context.Context, c *resources.MultyConfig, in WithResourceId, configPrefix string) (OutT, error) {
	for _, r := range c.Resources.GetAll() {
		if r.GetResourceId() == in.GetResourceId() {
			err := s.ServiceContext.DeploymentExecutor.MaybeInit(ctx, configPrefix)
			if err != nil {
				return *new(OutT), err
			}
			state, err := s.ServiceContext.DeploymentExecutor.GetState(ctx, configPrefix, s.ServiceContext.Database)
			if err != nil {
				return *new(OutT), err
			}
			m, err := r.GetMetadata(metadata.Metadatas)
			if err != nil {
				return *new(OutT), err
			}
			out, err := m.ReadFromState(r, state)
			if err != nil {
				return *new(OutT), errors.InternalServerErrorWithMessage(fmt.Sprintf("Unable to retrieve data for resource %s.", r.GetResourceId()), err)
			}
			return out.(OutT), err
		}
	}

	return *new(OutT), errors.ResourceNotFound(in.GetResourceId())
}

func (s Service[Arg, OutT]) Update(ctx context.Context, in UpdateRequest[Arg]) (out OutT, err error) {
	defer func() { s.updateErrorMetric(err, "update") }()
	return errors.WrappingErrors(s.update)(ctx, in)
}

func (s Service[Arg, OutT]) update(ctx context.Context, in UpdateRequest[Arg]) (out OutT, err error) {
	key, err := util.ExtractApiKey(ctx)
	if err != nil {
		return
	}
	userId, err := s.ServiceContext.GetUserId(ctx, key)
	if err != nil {
		return
	}
	configPrefix := GetConfigPrefix(in, userId)

	go s.ServiceContext.AwsClient.UpdateQPSMetric(userId, s.ResourceName, "update")
	log.Printf("[INFO] user: %s. update %s %s", userId, s.ResourceName, in.GetResourceId())
	lock, err := s.ServiceContext.LockConfig(ctx, userId, configPrefix)
	if err != nil {
		return
	}
	defer s.ServiceContext.UnlockConfig(ctx, lock)

	c, err := s.getConfig(ctx, userId, lock, configPrefix)
	if err != nil {
		return
	}

	r, err := c.UpdateResource(in.GetResourceId(), in.GetResource())
	if err != nil {
		return
	}

	defer func() {
		if err == nil {
			err = s.saveConfig(ctx, c, lock, configPrefix)
		} else {
			log.Println("[DEBUG] Something went wrong, not storing state")
		}
	}()

	rollbackFn, err := s.ServiceContext.DeploymentExecutor.Deploy(ctx, c, r, r, configPrefix)
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			rollbackFn()
		}
	}()
	return s.readFromConfig(ctx, c, in, configPrefix)
}

func (s Service[Arg, OutT]) Delete(ctx context.Context, in WithResourceId) (_ *commonpb.Empty, err error) {
	defer func() { s.updateErrorMetric(err, "delete") }()
	return errors.WrappingErrors(s.delete)(ctx, in)
}

func (s Service[Arg, OutT]) delete(ctx context.Context, in WithResourceId) (out *commonpb.Empty, err error) {
	key, err := util.ExtractApiKey(ctx)
	if err != nil {
		return
	}
	userId, err := s.ServiceContext.GetUserId(ctx, key)
	if err != nil {
		return
	}
	configPrefix := GetConfigPrefix(in, userId)
	go s.ServiceContext.AwsClient.UpdateQPSMetric(userId, s.ResourceName, "delete")
	log.Printf("[INFO] user: %s. delete %s %s", userId, s.ResourceName, in.GetResourceId())
	lock, err := s.ServiceContext.LockConfig(ctx, userId, configPrefix)
	if err != nil {
		return
	}
	defer s.ServiceContext.UnlockConfig(ctx, lock)
	c, err := s.getConfig(ctx, userId, lock, configPrefix)
	if err != nil {
		return
	}
	previousResource, err := c.DeleteResource(in.GetResourceId())
	if err != nil {
		return
	}

	defer func() {
		if err == nil {
			err = s.saveConfig(ctx, c, lock, configPrefix)
		} else {
			log.Println("[DEBUG] Something went wrong, not storing state")
		}
	}()

	_, err = s.ServiceContext.DeploymentExecutor.Deploy(ctx, c, previousResource, nil, configPrefix)
	if err != nil {
		if s, ok := status.FromError(err); ok && s.Code() == codes.InvalidArgument {
			for _, details := range s.Details() {
				v := details.(*pberr.ResourceValidationError)
				if v.GetNotFoundDetails() != nil && v.GetNotFoundDetails().ResourceId == in.GetResourceId() {
					return nil, errors.ResourceInUseError(in.GetResourceId(), v.ResourceId)
				}
			}
		}
		return nil, err
	}
	return &commonpb.Empty{}, nil
}
