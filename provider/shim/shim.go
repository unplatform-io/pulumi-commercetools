package shim

import (
	sdkschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pfschema "github.com/hashicorp/terraform-plugin-framework/provider"
	pfprovider "github.com/labd/terraform-provider-commercetools/internal/provider"
	sdkprovider "github.com/labd/terraform-provider-commercetools/commercetools"
)

func PFProvider() pfschema.Provider {
	prov := pfprovider.New("snapshot")
	
	return prov
}

func SDKProvider() *sdkschema.Provider {
	return sdkprovider.New("snapshot")()
}