// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package commercetools

import (
	"fmt"
	"path/filepath"
	"context"
	_ "embed"	
	provShim "github.com/labd/terraform-provider-commercetools/shim"
    pf "github.com/pulumi/pulumi-terraform-bridge/pf/tfbridge"
    "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"

	"github.com/unplatform-io/pulumi-commercetools/provider/pkg/version"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

// all of the token components used below.
const (
	// packages:
	mainPkg = "commercetools"
	// modules:
	mainMod = "index" // the y module
)

//go:embed cmd/pulumi-resource-commercetools/bridge-metadata.json
var bridgeMetadata []byte

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	// p := shimv2.NewProvider(commercetools.New(version.Version)())
	p := pf.MuxShimWithPF(context.Background(),
		shimv2.NewProvider(provShim.SDKProvider()),
		provShim.PFProvider(),
	)
	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:                 p,
		Name:              "commercetools",
		DisplayName:       "commercetools",
		Publisher:         "Aviva Solutions",
		LogoURL:           "",
		PluginDownloadURL: "",
		Description:       "A Pulumi package for creating and managing commercetools cloud resources.",
		Keywords:          []string{"pulumi", "commercetools"},
		License:           "Apache-2.0",
		Homepage:          "https://pulumi.io",
		Repository:        "https://github.com/unplatform-io/pulumi-commercetools",
		GitHubOrg:         "unplatform-io",
		Version:		   version.Version,
		Config:            map[string]*tfbridge.SchemaInfo{
			// Add any required configuration here, or remove the example below if
			// no additional points are required.
			// "region": {
			// 	Type: makeType("region", "Region"),
			// 	Default: &tfbridge.DefaultInfo{
			// 		EnvVars: []string{"AWS_REGION", "AWS_DEFAULT_REGION"},
			// 	},
			// },
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			"commercetools_subscription":  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Subscription")},
			"commercetools_api_client":    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ApiClient")},
			"commercetools_api_extension": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ApiExtension")},
			"commercetools_cart_discount": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "CartDiscount")},
			"commercetools_category": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Category")},
			"commercetools_channel":       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Channel")},
			"commercetools_custom_object": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "CustomObject")},
			"commercetools_customer_group": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "CustomerGroup")},
			"commercetools_discount_code": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "DiscountCode")},
			"commercetools_product_discount": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ProductDiscount")},
			"commercetools_product_type": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "ProductType"),
				// Rename ElementType, because a property with that name already exists in the GO SDK
				Fields: map[string]*tfbridge.SchemaInfo{
					"attribute": {
						Elem: &tfbridge.SchemaInfo{
							Fields: map[string]*tfbridge.SchemaInfo{
								"type": {
									Elem: &tfbridge.SchemaInfo{
										Fields: map[string]*tfbridge.SchemaInfo{
											"element_type": {
												Name: "ElementType2",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"commercetools_project_settings":   {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ProjectSettings")},
			"commercetools_shipping_method":    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ShippingMethod")},
			"commercetools_shipping_zone":      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ShippingZone")},
			"commercetools_shipping_zone_rate": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ShippingZoneRate")},
			"commercetools_state":              {Tok: tfbridge.MakeResource(mainPkg, mainMod, "State")},
			"commercetools_state_transitions":  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "StateTransitions")},
			"commercetools_store":              {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Store")},
			"commercetools_tax_category":       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "TaxCategory")},
			"commercetools_tax_category_rate":  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "TaxCategoryRate")},
			"commercetools_type": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "Type"),
				// // Rename ElementType, because a property with that name already exists in the GO SDK
				Fields: map[string]*tfbridge.SchemaInfo{
					"field": {
						Elem: &tfbridge.SchemaInfo{
							Fields: map[string]*tfbridge.SchemaInfo{
								"type": {
									Elem: &tfbridge.SchemaInfo{
										Fields: map[string]*tfbridge.SchemaInfo{
											"element_type": {
												Name: "ElementType2",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			// Map each resource in the Terraform provider to a Pulumi type. Two examples
			// are below - the single line form is the common case. The multi-line form is
			// needed only if you wish to override types or other default options.
			//
			// "aws_iam_role": {Tok: makeResource(mainMod, "IamRole")}
			//
			// "aws_acm_certificate": {
			// 	Tok: makeResource(mainMod, "Certificate"),
			// 	Fields: map[string]*tfbridge.SchemaInfo{
			// 		"tags": {Type: makeType(mainPkg, "Tags")},
			// 	},
			// },
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			// Map each resource in the Terraform provider to a Pulumi function. An example
			// is below.
			// "aws_ami": {Tok: makeDataSource(mainMod, "getAmi")},
			"commercetools_type": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getType")},
		},
		MetadataInfo: tfbridge.NewProviderMetadata(bridgeMetadata),
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			PackageName: "@unplatform/commercetools",
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}
	prov.SetAutonaming(255, "-")

	return prov
}
