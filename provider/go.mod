module github.com/unplatform-io/pulumi-commercetools/provider

go 1.15

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/hashicorp/terraform-exec => github.com/hashicorp/terraform-exec v0.1.1
)

require (
	github.com/hashicorp/terraform-plugin-sdk v1.15.0
	github.com/labd/terraform-provider-commercetools v0.25.1
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.11.0
	github.com/pulumi/pulumi/sdk/v2 v2.12.0
)
