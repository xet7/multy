package azure_resources

import (
	"fmt"
	"github.com/multycloud/multy/api/proto/commonpb"
	"github.com/multycloud/multy/api/proto/resourcespb"
	"github.com/multycloud/multy/flags"
	"github.com/multycloud/multy/resources"
	"github.com/multycloud/multy/resources/common"
	"github.com/multycloud/multy/resources/output"
	"github.com/multycloud/multy/resources/output/vault"
	"github.com/multycloud/multy/resources/types"
)

type AzureVault struct {
	*types.Vault
}

func InitVault(vn *types.Vault) resources.ResourceTranslator[*resourcespb.VaultResource] {
	return AzureVault{vn}
}

type AzureClientConfig struct {
	*output.TerraformDataSource `hcl:",squash" default:"name=azurerm_client_config"`
}

func (r AzureVault) FromState(state *output.TfState) (*resourcespb.VaultResource, error) {
	out := &resourcespb.VaultResource{
		CommonParameters: &commonpb.CommonResourceParameters{
			ResourceId:      r.ResourceId,
			ResourceGroupId: r.Args.CommonParameters.ResourceGroupId,
			Location:        r.Args.CommonParameters.Location,
			CloudProvider:   r.Args.CommonParameters.CloudProvider,
			NeedsUpdate:     false,
		},
		Name:        r.Args.Name,
		GcpOverride: r.Args.GcpOverride,
	}

	if flags.DryRun {
		return out, nil
	}
	statuses := map[string]commonpb.ResourceStatus_Status{}
	if stateResource, exists, err := output.MaybeGetParsedById[vault.AzureKeyVault](state, r.ResourceId); exists {
		if err != nil {
			return nil, err
		}
		out.Name = stateResource.Name
		out.AzureOutputs = &resourcespb.VaultAzureOutputs{KeyVaultId: stateResource.ResourceId}
	} else {
		statuses["azure_vault"] = commonpb.ResourceStatus_NEEDS_CREATE
	}

	if len(statuses) > 0 {
		out.CommonParameters.ResourceStatus = &commonpb.ResourceStatus{Statuses: statuses}
	}
	return out, nil

}

func (r AzureVault) Translate(resources.MultyContext) ([]output.TfBlock, error) {
	return []output.TfBlock{
		AzureClientConfig{TerraformDataSource: &output.TerraformDataSource{ResourceId: r.ResourceId}},
		vault.AzureKeyVault{
			AzResource: &common.AzResource{
				TerraformResource: output.TerraformResource{ResourceId: r.ResourceId},
				Name:              r.Args.Name,
				ResourceGroupName: GetResourceGroupName(r.Args.CommonParameters.ResourceGroupId),
				Location:          r.GetCloudSpecificLocation(),
			},
			Sku:      "standard",
			TenantId: fmt.Sprintf("data.azurerm_client_config.%s.tenant_id", r.ResourceId),
			AccessPolicy: []vault.AzureKeyVaultAccessPolicyInline{{
				TenantId: fmt.Sprintf(
					"data.azurerm_client_config.%s.tenant_id", r.ResourceId,
				),
				ObjectId: fmt.Sprintf(
					"data.azurerm_client_config.%s.object_id", r.ResourceId,
				),
				AzureKeyVaultPermissions: &vault.AzureKeyVaultPermissions{
					CertificatePermissions: []string{},
					KeyPermissions:         []string{},
					SecretPermissions:      []string{"List", "Get", "Set", "Delete", "Recover", "Backup", "Restore", "Purge"},
				},
			}},
		}}, nil
}

func (r AzureVault) GetMainResourceName() (string, error) {
	return vault.AzureResourceName, nil
}
