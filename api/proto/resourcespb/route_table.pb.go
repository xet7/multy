// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: api/proto/resourcespb/route_table.proto

package resourcespb

import (
	commonpb "github.com/multycloud/multy/api/proto/commonpb"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RouteDestination int32

const (
	RouteDestination_UNKNOWN_DESTINATION RouteDestination = 0
	RouteDestination_INTERNET            RouteDestination = 1
)

// Enum value maps for RouteDestination.
var (
	RouteDestination_name = map[int32]string{
		0: "UNKNOWN_DESTINATION",
		1: "INTERNET",
	}
	RouteDestination_value = map[string]int32{
		"UNKNOWN_DESTINATION": 0,
		"INTERNET":            1,
	}
)

func (x RouteDestination) Enum() *RouteDestination {
	p := new(RouteDestination)
	*p = x
	return p
}

func (x RouteDestination) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RouteDestination) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_resourcespb_route_table_proto_enumTypes[0].Descriptor()
}

func (RouteDestination) Type() protoreflect.EnumType {
	return &file_api_proto_resourcespb_route_table_proto_enumTypes[0]
}

func (x RouteDestination) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RouteDestination.Descriptor instead.
func (RouteDestination) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_resourcespb_route_table_proto_rawDescGZIP(), []int{0}
}

type CreateRouteTableRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resource *RouteTableArgs `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *CreateRouteTableRequest) Reset() {
	*x = CreateRouteTableRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_resourcespb_route_table_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRouteTableRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRouteTableRequest) ProtoMessage() {}

func (x *CreateRouteTableRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_resourcespb_route_table_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRouteTableRequest.ProtoReflect.Descriptor instead.
func (*CreateRouteTableRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_resourcespb_route_table_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRouteTableRequest) GetResource() *RouteTableArgs {
	if x != nil {
		return x.Resource
	}
	return nil
}

type ReadRouteTableRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
}

func (x *ReadRouteTableRequest) Reset() {
	*x = ReadRouteTableRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_resourcespb_route_table_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadRouteTableRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRouteTableRequest) ProtoMessage() {}

func (x *ReadRouteTableRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_resourcespb_route_table_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRouteTableRequest.ProtoReflect.Descriptor instead.
func (*ReadRouteTableRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_resourcespb_route_table_proto_rawDescGZIP(), []int{1}
}

func (x *ReadRouteTableRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

type UpdateRouteTableRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId string          `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	Resource   *RouteTableArgs `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *UpdateRouteTableRequest) Reset() {
	*x = UpdateRouteTableRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_resourcespb_route_table_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRouteTableRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRouteTableRequest) ProtoMessage() {}

func (x *UpdateRouteTableRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_resourcespb_route_table_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRouteTableRequest.ProtoReflect.Descriptor instead.
func (*UpdateRouteTableRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_resourcespb_route_table_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateRouteTableRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *UpdateRouteTableRequest) GetResource() *RouteTableArgs {
	if x != nil {
		return x.Resource
	}
	return nil
}

type DeleteRouteTableRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
}

func (x *DeleteRouteTableRequest) Reset() {
	*x = DeleteRouteTableRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_resourcespb_route_table_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRouteTableRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRouteTableRequest) ProtoMessage() {}

func (x *DeleteRouteTableRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_resourcespb_route_table_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRouteTableRequest.ProtoReflect.Descriptor instead.
func (*DeleteRouteTableRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_resourcespb_route_table_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteRouteTableRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

type Route struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CidrBlock   string           `protobuf:"bytes,1,opt,name=cidr_block,json=cidrBlock,proto3" json:"cidr_block,omitempty"`
	Destination RouteDestination `protobuf:"varint,2,opt,name=destination,proto3,enum=dev.multy.resources.RouteDestination" json:"destination,omitempty"`
}

func (x *Route) Reset() {
	*x = Route{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_resourcespb_route_table_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Route) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Route) ProtoMessage() {}

func (x *Route) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_resourcespb_route_table_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Route.ProtoReflect.Descriptor instead.
func (*Route) Descriptor() ([]byte, []int) {
	return file_api_proto_resourcespb_route_table_proto_rawDescGZIP(), []int{4}
}

func (x *Route) GetCidrBlock() string {
	if x != nil {
		return x.CidrBlock
	}
	return ""
}

func (x *Route) GetDestination() RouteDestination {
	if x != nil {
		return x.Destination
	}
	return RouteDestination_UNKNOWN_DESTINATION
}

type RouteTableArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonParameters *commonpb.ChildResourceCommonArgs `protobuf:"bytes,1,opt,name=common_parameters,json=commonParameters,proto3" json:"common_parameters,omitempty"`
	Name             string                            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	VirtualNetworkId string                            `protobuf:"bytes,3,opt,name=virtual_network_id,json=virtualNetworkId,proto3" json:"virtual_network_id,omitempty"`
	Routes           []*Route                          `protobuf:"bytes,4,rep,name=routes,proto3" json:"routes,omitempty"`
}

func (x *RouteTableArgs) Reset() {
	*x = RouteTableArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_resourcespb_route_table_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteTableArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteTableArgs) ProtoMessage() {}

func (x *RouteTableArgs) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_resourcespb_route_table_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteTableArgs.ProtoReflect.Descriptor instead.
func (*RouteTableArgs) Descriptor() ([]byte, []int) {
	return file_api_proto_resourcespb_route_table_proto_rawDescGZIP(), []int{5}
}

func (x *RouteTableArgs) GetCommonParameters() *commonpb.ChildResourceCommonArgs {
	if x != nil {
		return x.CommonParameters
	}
	return nil
}

func (x *RouteTableArgs) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RouteTableArgs) GetVirtualNetworkId() string {
	if x != nil {
		return x.VirtualNetworkId
	}
	return ""
}

func (x *RouteTableArgs) GetRoutes() []*Route {
	if x != nil {
		return x.Routes
	}
	return nil
}

type RouteTableAwsOutputs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RouteTableId string `protobuf:"bytes,1,opt,name=route_table_id,json=routeTableId,proto3" json:"route_table_id,omitempty"`
}

func (x *RouteTableAwsOutputs) Reset() {
	*x = RouteTableAwsOutputs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_resourcespb_route_table_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteTableAwsOutputs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteTableAwsOutputs) ProtoMessage() {}

func (x *RouteTableAwsOutputs) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_resourcespb_route_table_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteTableAwsOutputs.ProtoReflect.Descriptor instead.
func (*RouteTableAwsOutputs) Descriptor() ([]byte, []int) {
	return file_api_proto_resourcespb_route_table_proto_rawDescGZIP(), []int{6}
}

func (x *RouteTableAwsOutputs) GetRouteTableId() string {
	if x != nil {
		return x.RouteTableId
	}
	return ""
}

type RouteTableAzureOutputs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RouteTableId string `protobuf:"bytes,1,opt,name=route_table_id,json=routeTableId,proto3" json:"route_table_id,omitempty"`
}

func (x *RouteTableAzureOutputs) Reset() {
	*x = RouteTableAzureOutputs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_resourcespb_route_table_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteTableAzureOutputs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteTableAzureOutputs) ProtoMessage() {}

func (x *RouteTableAzureOutputs) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_resourcespb_route_table_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteTableAzureOutputs.ProtoReflect.Descriptor instead.
func (*RouteTableAzureOutputs) Descriptor() ([]byte, []int) {
	return file_api_proto_resourcespb_route_table_proto_rawDescGZIP(), []int{7}
}

func (x *RouteTableAzureOutputs) GetRouteTableId() string {
	if x != nil {
		return x.RouteTableId
	}
	return ""
}

type RouteTableGcpOutputs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ComputeRouteId []string `protobuf:"bytes,1,rep,name=compute_route_id,json=computeRouteId,proto3" json:"compute_route_id,omitempty"`
}

func (x *RouteTableGcpOutputs) Reset() {
	*x = RouteTableGcpOutputs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_resourcespb_route_table_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteTableGcpOutputs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteTableGcpOutputs) ProtoMessage() {}

func (x *RouteTableGcpOutputs) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_resourcespb_route_table_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteTableGcpOutputs.ProtoReflect.Descriptor instead.
func (*RouteTableGcpOutputs) Descriptor() ([]byte, []int) {
	return file_api_proto_resourcespb_route_table_proto_rawDescGZIP(), []int{8}
}

func (x *RouteTableGcpOutputs) GetComputeRouteId() []string {
	if x != nil {
		return x.ComputeRouteId
	}
	return nil
}

type RouteTableResource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonParameters *commonpb.CommonChildResourceParameters `protobuf:"bytes,1,opt,name=common_parameters,json=commonParameters,proto3" json:"common_parameters,omitempty"`
	Name             string                                  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	VirtualNetworkId string                                  `protobuf:"bytes,3,opt,name=virtual_network_id,json=virtualNetworkId,proto3" json:"virtual_network_id,omitempty"`
	Routes           []*Route                                `protobuf:"bytes,4,rep,name=routes,proto3" json:"routes,omitempty"`
	AwsOutputs       *RouteTableAwsOutputs                   `protobuf:"bytes,5,opt,name=aws_outputs,json=awsOutputs,proto3" json:"aws_outputs,omitempty"`
	AzureOutputs     *RouteTableAzureOutputs                 `protobuf:"bytes,6,opt,name=azure_outputs,json=azureOutputs,proto3" json:"azure_outputs,omitempty"`
	GcpOutputs       *RouteTableGcpOutputs                   `protobuf:"bytes,7,opt,name=gcp_outputs,json=gcpOutputs,proto3" json:"gcp_outputs,omitempty"`
}

func (x *RouteTableResource) Reset() {
	*x = RouteTableResource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_resourcespb_route_table_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteTableResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteTableResource) ProtoMessage() {}

func (x *RouteTableResource) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_resourcespb_route_table_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteTableResource.ProtoReflect.Descriptor instead.
func (*RouteTableResource) Descriptor() ([]byte, []int) {
	return file_api_proto_resourcespb_route_table_proto_rawDescGZIP(), []int{9}
}

func (x *RouteTableResource) GetCommonParameters() *commonpb.CommonChildResourceParameters {
	if x != nil {
		return x.CommonParameters
	}
	return nil
}

func (x *RouteTableResource) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RouteTableResource) GetVirtualNetworkId() string {
	if x != nil {
		return x.VirtualNetworkId
	}
	return ""
}

func (x *RouteTableResource) GetRoutes() []*Route {
	if x != nil {
		return x.Routes
	}
	return nil
}

func (x *RouteTableResource) GetAwsOutputs() *RouteTableAwsOutputs {
	if x != nil {
		return x.AwsOutputs
	}
	return nil
}

func (x *RouteTableResource) GetAzureOutputs() *RouteTableAzureOutputs {
	if x != nil {
		return x.AzureOutputs
	}
	return nil
}

func (x *RouteTableResource) GetGcpOutputs() *RouteTableGcpOutputs {
	if x != nil {
		return x.GcpOutputs
	}
	return nil
}

var File_api_proto_resourcespb_route_table_proto protoreflect.FileDescriptor

var file_api_proto_resourcespb_route_table_proto_rawDesc = []byte{
	0x0a, 0x27, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x70, 0x62, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x64, 0x65, 0x76, 0x2e, 0x6d,
	0x75, 0x6c, 0x74, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x1f,
	0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x70, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x5a, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x08, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x64,
	0x65, 0x76, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x72, 0x67,
	0x73, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x38, 0x0a, 0x15, 0x52,
	0x65, 0x61, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x49, 0x64, 0x22, 0x7b, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49,
	0x64, 0x12, 0x3f, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x79, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x41, 0x72, 0x67, 0x73, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x22, 0x3a, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x22, 0x6f,
	0x0a, 0x05, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x69, 0x64, 0x72, 0x5f,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x69, 0x64,
	0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x47, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x64, 0x65,
	0x76, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0xde, 0x01, 0x0a, 0x0e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x72,
	0x67, 0x73, 0x12, 0x56, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e,
	0x64, 0x65, 0x76, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x41, 0x72, 0x67, 0x73, 0x52, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2c,
	0x0a, 0x12, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x76, 0x69, 0x72, 0x74,
	0x75, 0x61, 0x6c, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x06,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x64,
	0x65, 0x76, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73,
	0x22, 0x3c, 0x0a, 0x14, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x77,
	0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x12, 0x24, 0x0a, 0x0e, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x3e,
	0x0a, 0x16, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x7a, 0x75, 0x72,
	0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x12, 0x24, 0x0a, 0x0e, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x40,
	0x0a, 0x14, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x47, 0x63, 0x70, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74,
	0x65, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x49, 0x64,
	0x22, 0xd2, 0x03, 0x0a, 0x12, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x5c, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x79, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x43, 0x68, 0x69, 0x6c,
	0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x73, 0x52, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x76, 0x69, 0x72,
	0x74, 0x75, 0x61, 0x6c, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6d, 0x75,
	0x6c, 0x74, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x4a, 0x0a, 0x0b, 0x61,
	0x77, 0x73, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x29, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x79, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x41, 0x77, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x52, 0x0a, 0x61, 0x77, 0x73,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x12, 0x50, 0x0a, 0x0d, 0x61, 0x7a, 0x75, 0x72, 0x65,
	0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b,
	0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x41,
	0x7a, 0x75, 0x72, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x52, 0x0c, 0x61, 0x7a, 0x75,
	0x72, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x12, 0x4a, 0x0a, 0x0b, 0x67, 0x63, 0x70,
	0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x47,
	0x63, 0x70, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x52, 0x0a, 0x67, 0x63, 0x70, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x73, 0x2a, 0x39, 0x0a, 0x10, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x44, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x13, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x44, 0x45, 0x53, 0x54, 0x49, 0x4e, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x45, 0x54, 0x10, 0x01,
	0x42, 0x5e, 0x0a, 0x17, 0x64, 0x65, 0x76, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x79, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x0e, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x31, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x79, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_resourcespb_route_table_proto_rawDescOnce sync.Once
	file_api_proto_resourcespb_route_table_proto_rawDescData = file_api_proto_resourcespb_route_table_proto_rawDesc
)

func file_api_proto_resourcespb_route_table_proto_rawDescGZIP() []byte {
	file_api_proto_resourcespb_route_table_proto_rawDescOnce.Do(func() {
		file_api_proto_resourcespb_route_table_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_resourcespb_route_table_proto_rawDescData)
	})
	return file_api_proto_resourcespb_route_table_proto_rawDescData
}

var file_api_proto_resourcespb_route_table_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_proto_resourcespb_route_table_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_proto_resourcespb_route_table_proto_goTypes = []interface{}{
	(RouteDestination)(0),                          // 0: dev.multy.resources.RouteDestination
	(*CreateRouteTableRequest)(nil),                // 1: dev.multy.resources.CreateRouteTableRequest
	(*ReadRouteTableRequest)(nil),                  // 2: dev.multy.resources.ReadRouteTableRequest
	(*UpdateRouteTableRequest)(nil),                // 3: dev.multy.resources.UpdateRouteTableRequest
	(*DeleteRouteTableRequest)(nil),                // 4: dev.multy.resources.DeleteRouteTableRequest
	(*Route)(nil),                                  // 5: dev.multy.resources.Route
	(*RouteTableArgs)(nil),                         // 6: dev.multy.resources.RouteTableArgs
	(*RouteTableAwsOutputs)(nil),                   // 7: dev.multy.resources.RouteTableAwsOutputs
	(*RouteTableAzureOutputs)(nil),                 // 8: dev.multy.resources.RouteTableAzureOutputs
	(*RouteTableGcpOutputs)(nil),                   // 9: dev.multy.resources.RouteTableGcpOutputs
	(*RouteTableResource)(nil),                     // 10: dev.multy.resources.RouteTableResource
	(*commonpb.ChildResourceCommonArgs)(nil),       // 11: dev.multy.common.ChildResourceCommonArgs
	(*commonpb.CommonChildResourceParameters)(nil), // 12: dev.multy.common.CommonChildResourceParameters
}
var file_api_proto_resourcespb_route_table_proto_depIdxs = []int32{
	6,  // 0: dev.multy.resources.CreateRouteTableRequest.resource:type_name -> dev.multy.resources.RouteTableArgs
	6,  // 1: dev.multy.resources.UpdateRouteTableRequest.resource:type_name -> dev.multy.resources.RouteTableArgs
	0,  // 2: dev.multy.resources.Route.destination:type_name -> dev.multy.resources.RouteDestination
	11, // 3: dev.multy.resources.RouteTableArgs.common_parameters:type_name -> dev.multy.common.ChildResourceCommonArgs
	5,  // 4: dev.multy.resources.RouteTableArgs.routes:type_name -> dev.multy.resources.Route
	12, // 5: dev.multy.resources.RouteTableResource.common_parameters:type_name -> dev.multy.common.CommonChildResourceParameters
	5,  // 6: dev.multy.resources.RouteTableResource.routes:type_name -> dev.multy.resources.Route
	7,  // 7: dev.multy.resources.RouteTableResource.aws_outputs:type_name -> dev.multy.resources.RouteTableAwsOutputs
	8,  // 8: dev.multy.resources.RouteTableResource.azure_outputs:type_name -> dev.multy.resources.RouteTableAzureOutputs
	9,  // 9: dev.multy.resources.RouteTableResource.gcp_outputs:type_name -> dev.multy.resources.RouteTableGcpOutputs
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_api_proto_resourcespb_route_table_proto_init() }
func file_api_proto_resourcespb_route_table_proto_init() {
	if File_api_proto_resourcespb_route_table_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_resourcespb_route_table_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRouteTableRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_resourcespb_route_table_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadRouteTableRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_resourcespb_route_table_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRouteTableRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_resourcespb_route_table_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRouteTableRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_resourcespb_route_table_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Route); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_resourcespb_route_table_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteTableArgs); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_resourcespb_route_table_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteTableAwsOutputs); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_resourcespb_route_table_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteTableAzureOutputs); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_resourcespb_route_table_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteTableGcpOutputs); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_resourcespb_route_table_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteTableResource); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_resourcespb_route_table_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_proto_resourcespb_route_table_proto_goTypes,
		DependencyIndexes: file_api_proto_resourcespb_route_table_proto_depIdxs,
		EnumInfos:         file_api_proto_resourcespb_route_table_proto_enumTypes,
		MessageInfos:      file_api_proto_resourcespb_route_table_proto_msgTypes,
	}.Build()
	File_api_proto_resourcespb_route_table_proto = out.File
	file_api_proto_resourcespb_route_table_proto_rawDesc = nil
	file_api_proto_resourcespb_route_table_proto_goTypes = nil
	file_api_proto_resourcespb_route_table_proto_depIdxs = nil
}
