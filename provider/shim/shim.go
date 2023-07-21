package shim

import (
	pfschema "github.com/hashicorp/terraform-plugin-framework/provider"
	sdkschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	sdkprovider "github.com/labd/terraform-provider-commercetools/commercetools"
	pfprovider "github.com/labd/terraform-provider-commercetools/internal/provider"
)

func PFProvider() pfschema.Provider {
	prov := pfprovider.New("snapshot")

	return prov
}

func SDKProvider() *sdkschema.Provider {
	return sdkprovider.New("snapshot")()
}
