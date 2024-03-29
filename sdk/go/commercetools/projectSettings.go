// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package commercetools

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-commercetools/sdk/go/commercetools/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ProjectSettings struct {
	pulumi.CustomResourceState

	// [Carts Configuration](https://docs.commercetools.com/api/projects/project#carts-configuration)
	Carts ProjectSettingsCartsPtrOutput `pulumi:"carts"`
	// A two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)
	Countries pulumi.StringArrayOutput `pulumi:"countries"`
	// A three-digit currency code as per [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217)
	Currencies pulumi.StringArrayOutput `pulumi:"currencies"`
	// Enable the Search Indexing of orders
	EnableSearchIndexOrders pulumi.BoolOutput `pulumi:"enableSearchIndexOrders"`
	// Enable the Search Indexing of products
	EnableSearchIndexProducts pulumi.BoolOutput `pulumi:"enableSearchIndexProducts"`
	// [External OAUTH](https://docs.commercetools.com/api/projects/project#externaloauth)
	ExternalOauth ProjectSettingsExternalOauthPtrOutput `pulumi:"externalOauth"`
	// The unique key of the project
	Key pulumi.StringOutput `pulumi:"key"`
	// [IETF Language Tag](https://en.wikipedia.org/wiki/IETF_language_tag)
	Languages pulumi.StringArrayOutput `pulumi:"languages"`
	// The change notifications subscribed to
	Messages ProjectSettingsMessagesPtrOutput `pulumi:"messages"`
	// The name of the project
	Name pulumi.StringOutput `pulumi:"name"`
	// If shipping_rate_input_type is set to CartClassification these values are used to create tiers . Only a key defined
	// inside the values array can be used to create a tier, or to set a value for the shippingRateInput on the cart. The keys
	// are checked for uniqueness and the request is rejected if keys are not unique
	ShippingRateCartClassificationValues ProjectSettingsShippingRateCartClassificationValueArrayOutput `pulumi:"shippingRateCartClassificationValues"`
	// Three ways to dynamically select a ShippingRatePriceTier exist. The CartValue type uses the sum of all line item prices,
	// whereas CartClassification and CartScore use the shippingRateInput field on the cart to select a tier
	ShippingRateInputType pulumi.StringPtrOutput `pulumi:"shippingRateInputType"`
	Version               pulumi.IntOutput       `pulumi:"version"`
}

// NewProjectSettings registers a new resource with the given unique name, arguments, and options.
func NewProjectSettings(ctx *pulumi.Context,
	name string, args *ProjectSettingsArgs, opts ...pulumi.ResourceOption) (*ProjectSettings, error) {
	if args == nil {
		args = &ProjectSettingsArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProjectSettings
	err := ctx.RegisterResource("commercetools:index/projectSettings:ProjectSettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectSettings gets an existing ProjectSettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectSettingsState, opts ...pulumi.ResourceOption) (*ProjectSettings, error) {
	var resource ProjectSettings
	err := ctx.ReadResource("commercetools:index/projectSettings:ProjectSettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectSettings resources.
type projectSettingsState struct {
	// [Carts Configuration](https://docs.commercetools.com/api/projects/project#carts-configuration)
	Carts *ProjectSettingsCarts `pulumi:"carts"`
	// A two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)
	Countries []string `pulumi:"countries"`
	// A three-digit currency code as per [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217)
	Currencies []string `pulumi:"currencies"`
	// Enable the Search Indexing of orders
	EnableSearchIndexOrders *bool `pulumi:"enableSearchIndexOrders"`
	// Enable the Search Indexing of products
	EnableSearchIndexProducts *bool `pulumi:"enableSearchIndexProducts"`
	// [External OAUTH](https://docs.commercetools.com/api/projects/project#externaloauth)
	ExternalOauth *ProjectSettingsExternalOauth `pulumi:"externalOauth"`
	// The unique key of the project
	Key *string `pulumi:"key"`
	// [IETF Language Tag](https://en.wikipedia.org/wiki/IETF_language_tag)
	Languages []string `pulumi:"languages"`
	// The change notifications subscribed to
	Messages *ProjectSettingsMessages `pulumi:"messages"`
	// The name of the project
	Name *string `pulumi:"name"`
	// If shipping_rate_input_type is set to CartClassification these values are used to create tiers . Only a key defined
	// inside the values array can be used to create a tier, or to set a value for the shippingRateInput on the cart. The keys
	// are checked for uniqueness and the request is rejected if keys are not unique
	ShippingRateCartClassificationValues []ProjectSettingsShippingRateCartClassificationValue `pulumi:"shippingRateCartClassificationValues"`
	// Three ways to dynamically select a ShippingRatePriceTier exist. The CartValue type uses the sum of all line item prices,
	// whereas CartClassification and CartScore use the shippingRateInput field on the cart to select a tier
	ShippingRateInputType *string `pulumi:"shippingRateInputType"`
	Version               *int    `pulumi:"version"`
}

type ProjectSettingsState struct {
	// [Carts Configuration](https://docs.commercetools.com/api/projects/project#carts-configuration)
	Carts ProjectSettingsCartsPtrInput
	// A two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)
	Countries pulumi.StringArrayInput
	// A three-digit currency code as per [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217)
	Currencies pulumi.StringArrayInput
	// Enable the Search Indexing of orders
	EnableSearchIndexOrders pulumi.BoolPtrInput
	// Enable the Search Indexing of products
	EnableSearchIndexProducts pulumi.BoolPtrInput
	// [External OAUTH](https://docs.commercetools.com/api/projects/project#externaloauth)
	ExternalOauth ProjectSettingsExternalOauthPtrInput
	// The unique key of the project
	Key pulumi.StringPtrInput
	// [IETF Language Tag](https://en.wikipedia.org/wiki/IETF_language_tag)
	Languages pulumi.StringArrayInput
	// The change notifications subscribed to
	Messages ProjectSettingsMessagesPtrInput
	// The name of the project
	Name pulumi.StringPtrInput
	// If shipping_rate_input_type is set to CartClassification these values are used to create tiers . Only a key defined
	// inside the values array can be used to create a tier, or to set a value for the shippingRateInput on the cart. The keys
	// are checked for uniqueness and the request is rejected if keys are not unique
	ShippingRateCartClassificationValues ProjectSettingsShippingRateCartClassificationValueArrayInput
	// Three ways to dynamically select a ShippingRatePriceTier exist. The CartValue type uses the sum of all line item prices,
	// whereas CartClassification and CartScore use the shippingRateInput field on the cart to select a tier
	ShippingRateInputType pulumi.StringPtrInput
	Version               pulumi.IntPtrInput
}

func (ProjectSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectSettingsState)(nil)).Elem()
}

type projectSettingsArgs struct {
	// [Carts Configuration](https://docs.commercetools.com/api/projects/project#carts-configuration)
	Carts *ProjectSettingsCarts `pulumi:"carts"`
	// A two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)
	Countries []string `pulumi:"countries"`
	// A three-digit currency code as per [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217)
	Currencies []string `pulumi:"currencies"`
	// Enable the Search Indexing of orders
	EnableSearchIndexOrders *bool `pulumi:"enableSearchIndexOrders"`
	// Enable the Search Indexing of products
	EnableSearchIndexProducts *bool `pulumi:"enableSearchIndexProducts"`
	// [External OAUTH](https://docs.commercetools.com/api/projects/project#externaloauth)
	ExternalOauth *ProjectSettingsExternalOauth `pulumi:"externalOauth"`
	// [IETF Language Tag](https://en.wikipedia.org/wiki/IETF_language_tag)
	Languages []string `pulumi:"languages"`
	// The change notifications subscribed to
	Messages *ProjectSettingsMessages `pulumi:"messages"`
	// The name of the project
	Name *string `pulumi:"name"`
	// If shipping_rate_input_type is set to CartClassification these values are used to create tiers . Only a key defined
	// inside the values array can be used to create a tier, or to set a value for the shippingRateInput on the cart. The keys
	// are checked for uniqueness and the request is rejected if keys are not unique
	ShippingRateCartClassificationValues []ProjectSettingsShippingRateCartClassificationValue `pulumi:"shippingRateCartClassificationValues"`
	// Three ways to dynamically select a ShippingRatePriceTier exist. The CartValue type uses the sum of all line item prices,
	// whereas CartClassification and CartScore use the shippingRateInput field on the cart to select a tier
	ShippingRateInputType *string `pulumi:"shippingRateInputType"`
}

// The set of arguments for constructing a ProjectSettings resource.
type ProjectSettingsArgs struct {
	// [Carts Configuration](https://docs.commercetools.com/api/projects/project#carts-configuration)
	Carts ProjectSettingsCartsPtrInput
	// A two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)
	Countries pulumi.StringArrayInput
	// A three-digit currency code as per [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217)
	Currencies pulumi.StringArrayInput
	// Enable the Search Indexing of orders
	EnableSearchIndexOrders pulumi.BoolPtrInput
	// Enable the Search Indexing of products
	EnableSearchIndexProducts pulumi.BoolPtrInput
	// [External OAUTH](https://docs.commercetools.com/api/projects/project#externaloauth)
	ExternalOauth ProjectSettingsExternalOauthPtrInput
	// [IETF Language Tag](https://en.wikipedia.org/wiki/IETF_language_tag)
	Languages pulumi.StringArrayInput
	// The change notifications subscribed to
	Messages ProjectSettingsMessagesPtrInput
	// The name of the project
	Name pulumi.StringPtrInput
	// If shipping_rate_input_type is set to CartClassification these values are used to create tiers . Only a key defined
	// inside the values array can be used to create a tier, or to set a value for the shippingRateInput on the cart. The keys
	// are checked for uniqueness and the request is rejected if keys are not unique
	ShippingRateCartClassificationValues ProjectSettingsShippingRateCartClassificationValueArrayInput
	// Three ways to dynamically select a ShippingRatePriceTier exist. The CartValue type uses the sum of all line item prices,
	// whereas CartClassification and CartScore use the shippingRateInput field on the cart to select a tier
	ShippingRateInputType pulumi.StringPtrInput
}

func (ProjectSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectSettingsArgs)(nil)).Elem()
}

type ProjectSettingsInput interface {
	pulumi.Input

	ToProjectSettingsOutput() ProjectSettingsOutput
	ToProjectSettingsOutputWithContext(ctx context.Context) ProjectSettingsOutput
}

func (*ProjectSettings) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectSettings)(nil)).Elem()
}

func (i *ProjectSettings) ToProjectSettingsOutput() ProjectSettingsOutput {
	return i.ToProjectSettingsOutputWithContext(context.Background())
}

func (i *ProjectSettings) ToProjectSettingsOutputWithContext(ctx context.Context) ProjectSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectSettingsOutput)
}

// ProjectSettingsArrayInput is an input type that accepts ProjectSettingsArray and ProjectSettingsArrayOutput values.
// You can construct a concrete instance of `ProjectSettingsArrayInput` via:
//
//	ProjectSettingsArray{ ProjectSettingsArgs{...} }
type ProjectSettingsArrayInput interface {
	pulumi.Input

	ToProjectSettingsArrayOutput() ProjectSettingsArrayOutput
	ToProjectSettingsArrayOutputWithContext(context.Context) ProjectSettingsArrayOutput
}

type ProjectSettingsArray []ProjectSettingsInput

func (ProjectSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectSettings)(nil)).Elem()
}

func (i ProjectSettingsArray) ToProjectSettingsArrayOutput() ProjectSettingsArrayOutput {
	return i.ToProjectSettingsArrayOutputWithContext(context.Background())
}

func (i ProjectSettingsArray) ToProjectSettingsArrayOutputWithContext(ctx context.Context) ProjectSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectSettingsArrayOutput)
}

// ProjectSettingsMapInput is an input type that accepts ProjectSettingsMap and ProjectSettingsMapOutput values.
// You can construct a concrete instance of `ProjectSettingsMapInput` via:
//
//	ProjectSettingsMap{ "key": ProjectSettingsArgs{...} }
type ProjectSettingsMapInput interface {
	pulumi.Input

	ToProjectSettingsMapOutput() ProjectSettingsMapOutput
	ToProjectSettingsMapOutputWithContext(context.Context) ProjectSettingsMapOutput
}

type ProjectSettingsMap map[string]ProjectSettingsInput

func (ProjectSettingsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectSettings)(nil)).Elem()
}

func (i ProjectSettingsMap) ToProjectSettingsMapOutput() ProjectSettingsMapOutput {
	return i.ToProjectSettingsMapOutputWithContext(context.Background())
}

func (i ProjectSettingsMap) ToProjectSettingsMapOutputWithContext(ctx context.Context) ProjectSettingsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectSettingsMapOutput)
}

type ProjectSettingsOutput struct{ *pulumi.OutputState }

func (ProjectSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectSettings)(nil)).Elem()
}

func (o ProjectSettingsOutput) ToProjectSettingsOutput() ProjectSettingsOutput {
	return o
}

func (o ProjectSettingsOutput) ToProjectSettingsOutputWithContext(ctx context.Context) ProjectSettingsOutput {
	return o
}

// [Carts Configuration](https://docs.commercetools.com/api/projects/project#carts-configuration)
func (o ProjectSettingsOutput) Carts() ProjectSettingsCartsPtrOutput {
	return o.ApplyT(func(v *ProjectSettings) ProjectSettingsCartsPtrOutput { return v.Carts }).(ProjectSettingsCartsPtrOutput)
}

// A two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)
func (o ProjectSettingsOutput) Countries() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ProjectSettings) pulumi.StringArrayOutput { return v.Countries }).(pulumi.StringArrayOutput)
}

// A three-digit currency code as per [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217)
func (o ProjectSettingsOutput) Currencies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ProjectSettings) pulumi.StringArrayOutput { return v.Currencies }).(pulumi.StringArrayOutput)
}

// Enable the Search Indexing of orders
func (o ProjectSettingsOutput) EnableSearchIndexOrders() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectSettings) pulumi.BoolOutput { return v.EnableSearchIndexOrders }).(pulumi.BoolOutput)
}

// Enable the Search Indexing of products
func (o ProjectSettingsOutput) EnableSearchIndexProducts() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectSettings) pulumi.BoolOutput { return v.EnableSearchIndexProducts }).(pulumi.BoolOutput)
}

// [External OAUTH](https://docs.commercetools.com/api/projects/project#externaloauth)
func (o ProjectSettingsOutput) ExternalOauth() ProjectSettingsExternalOauthPtrOutput {
	return o.ApplyT(func(v *ProjectSettings) ProjectSettingsExternalOauthPtrOutput { return v.ExternalOauth }).(ProjectSettingsExternalOauthPtrOutput)
}

// The unique key of the project
func (o ProjectSettingsOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectSettings) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// [IETF Language Tag](https://en.wikipedia.org/wiki/IETF_language_tag)
func (o ProjectSettingsOutput) Languages() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ProjectSettings) pulumi.StringArrayOutput { return v.Languages }).(pulumi.StringArrayOutput)
}

// The change notifications subscribed to
func (o ProjectSettingsOutput) Messages() ProjectSettingsMessagesPtrOutput {
	return o.ApplyT(func(v *ProjectSettings) ProjectSettingsMessagesPtrOutput { return v.Messages }).(ProjectSettingsMessagesPtrOutput)
}

// The name of the project
func (o ProjectSettingsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectSettings) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// If shipping_rate_input_type is set to CartClassification these values are used to create tiers . Only a key defined
// inside the values array can be used to create a tier, or to set a value for the shippingRateInput on the cart. The keys
// are checked for uniqueness and the request is rejected if keys are not unique
func (o ProjectSettingsOutput) ShippingRateCartClassificationValues() ProjectSettingsShippingRateCartClassificationValueArrayOutput {
	return o.ApplyT(func(v *ProjectSettings) ProjectSettingsShippingRateCartClassificationValueArrayOutput {
		return v.ShippingRateCartClassificationValues
	}).(ProjectSettingsShippingRateCartClassificationValueArrayOutput)
}

// Three ways to dynamically select a ShippingRatePriceTier exist. The CartValue type uses the sum of all line item prices,
// whereas CartClassification and CartScore use the shippingRateInput field on the cart to select a tier
func (o ProjectSettingsOutput) ShippingRateInputType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectSettings) pulumi.StringPtrOutput { return v.ShippingRateInputType }).(pulumi.StringPtrOutput)
}

func (o ProjectSettingsOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v *ProjectSettings) pulumi.IntOutput { return v.Version }).(pulumi.IntOutput)
}

type ProjectSettingsArrayOutput struct{ *pulumi.OutputState }

func (ProjectSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectSettings)(nil)).Elem()
}

func (o ProjectSettingsArrayOutput) ToProjectSettingsArrayOutput() ProjectSettingsArrayOutput {
	return o
}

func (o ProjectSettingsArrayOutput) ToProjectSettingsArrayOutputWithContext(ctx context.Context) ProjectSettingsArrayOutput {
	return o
}

func (o ProjectSettingsArrayOutput) Index(i pulumi.IntInput) ProjectSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProjectSettings {
		return vs[0].([]*ProjectSettings)[vs[1].(int)]
	}).(ProjectSettingsOutput)
}

type ProjectSettingsMapOutput struct{ *pulumi.OutputState }

func (ProjectSettingsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectSettings)(nil)).Elem()
}

func (o ProjectSettingsMapOutput) ToProjectSettingsMapOutput() ProjectSettingsMapOutput {
	return o
}

func (o ProjectSettingsMapOutput) ToProjectSettingsMapOutputWithContext(ctx context.Context) ProjectSettingsMapOutput {
	return o
}

func (o ProjectSettingsMapOutput) MapIndex(k pulumi.StringInput) ProjectSettingsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProjectSettings {
		return vs[0].(map[string]*ProjectSettings)[vs[1].(string)]
	}).(ProjectSettingsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectSettingsInput)(nil)).Elem(), &ProjectSettings{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectSettingsArrayInput)(nil)).Elem(), ProjectSettingsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectSettingsMapInput)(nil)).Elem(), ProjectSettingsMap{})
	pulumi.RegisterOutputType(ProjectSettingsOutput{})
	pulumi.RegisterOutputType(ProjectSettingsArrayOutput{})
	pulumi.RegisterOutputType(ProjectSettingsMapOutput{})
}
