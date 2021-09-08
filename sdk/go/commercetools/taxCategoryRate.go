// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package commercetools

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TaxCategoryRate struct {
	pulumi.CustomResourceState

	// Number Percentage in the range of [0..1]. The sum of the amounts of all subRates, if there are any
	Amount pulumi.Float64PtrOutput `pulumi:"amount"`
	// A two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)
	Country         pulumi.StringOutput `pulumi:"country"`
	IncludedInPrice pulumi.BoolOutput   `pulumi:"includedInPrice"`
	Name            pulumi.StringOutput `pulumi:"name"`
	// The state in the country
	State pulumi.StringPtrOutput `pulumi:"state"`
	// For countries (for example the US) where the total tax is a combination of multiple taxes (for example state and local
	// taxes)
	SubRates      TaxCategoryRateSubRateArrayOutput `pulumi:"subRates"`
	TaxCategoryId pulumi.StringOutput               `pulumi:"taxCategoryId"`
}

// NewTaxCategoryRate registers a new resource with the given unique name, arguments, and options.
func NewTaxCategoryRate(ctx *pulumi.Context,
	name string, args *TaxCategoryRateArgs, opts ...pulumi.ResourceOption) (*TaxCategoryRate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Country == nil {
		return nil, errors.New("invalid value for required argument 'Country'")
	}
	if args.IncludedInPrice == nil {
		return nil, errors.New("invalid value for required argument 'IncludedInPrice'")
	}
	if args.TaxCategoryId == nil {
		return nil, errors.New("invalid value for required argument 'TaxCategoryId'")
	}
	var resource TaxCategoryRate
	err := ctx.RegisterResource("commercetools:index/taxCategoryRate:TaxCategoryRate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTaxCategoryRate gets an existing TaxCategoryRate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTaxCategoryRate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TaxCategoryRateState, opts ...pulumi.ResourceOption) (*TaxCategoryRate, error) {
	var resource TaxCategoryRate
	err := ctx.ReadResource("commercetools:index/taxCategoryRate:TaxCategoryRate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TaxCategoryRate resources.
type taxCategoryRateState struct {
	// Number Percentage in the range of [0..1]. The sum of the amounts of all subRates, if there are any
	Amount *float64 `pulumi:"amount"`
	// A two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)
	Country         *string `pulumi:"country"`
	IncludedInPrice *bool   `pulumi:"includedInPrice"`
	Name            *string `pulumi:"name"`
	// The state in the country
	State *string `pulumi:"state"`
	// For countries (for example the US) where the total tax is a combination of multiple taxes (for example state and local
	// taxes)
	SubRates      []TaxCategoryRateSubRate `pulumi:"subRates"`
	TaxCategoryId *string                  `pulumi:"taxCategoryId"`
}

type TaxCategoryRateState struct {
	// Number Percentage in the range of [0..1]. The sum of the amounts of all subRates, if there are any
	Amount pulumi.Float64PtrInput
	// A two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)
	Country         pulumi.StringPtrInput
	IncludedInPrice pulumi.BoolPtrInput
	Name            pulumi.StringPtrInput
	// The state in the country
	State pulumi.StringPtrInput
	// For countries (for example the US) where the total tax is a combination of multiple taxes (for example state and local
	// taxes)
	SubRates      TaxCategoryRateSubRateArrayInput
	TaxCategoryId pulumi.StringPtrInput
}

func (TaxCategoryRateState) ElementType() reflect.Type {
	return reflect.TypeOf((*taxCategoryRateState)(nil)).Elem()
}

type taxCategoryRateArgs struct {
	// Number Percentage in the range of [0..1]. The sum of the amounts of all subRates, if there are any
	Amount *float64 `pulumi:"amount"`
	// A two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)
	Country         string  `pulumi:"country"`
	IncludedInPrice bool    `pulumi:"includedInPrice"`
	Name            *string `pulumi:"name"`
	// The state in the country
	State *string `pulumi:"state"`
	// For countries (for example the US) where the total tax is a combination of multiple taxes (for example state and local
	// taxes)
	SubRates      []TaxCategoryRateSubRate `pulumi:"subRates"`
	TaxCategoryId string                   `pulumi:"taxCategoryId"`
}

// The set of arguments for constructing a TaxCategoryRate resource.
type TaxCategoryRateArgs struct {
	// Number Percentage in the range of [0..1]. The sum of the amounts of all subRates, if there are any
	Amount pulumi.Float64PtrInput
	// A two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)
	Country         pulumi.StringInput
	IncludedInPrice pulumi.BoolInput
	Name            pulumi.StringPtrInput
	// The state in the country
	State pulumi.StringPtrInput
	// For countries (for example the US) where the total tax is a combination of multiple taxes (for example state and local
	// taxes)
	SubRates      TaxCategoryRateSubRateArrayInput
	TaxCategoryId pulumi.StringInput
}

func (TaxCategoryRateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*taxCategoryRateArgs)(nil)).Elem()
}

type TaxCategoryRateInput interface {
	pulumi.Input

	ToTaxCategoryRateOutput() TaxCategoryRateOutput
	ToTaxCategoryRateOutputWithContext(ctx context.Context) TaxCategoryRateOutput
}

func (*TaxCategoryRate) ElementType() reflect.Type {
	return reflect.TypeOf((*TaxCategoryRate)(nil))
}

func (i *TaxCategoryRate) ToTaxCategoryRateOutput() TaxCategoryRateOutput {
	return i.ToTaxCategoryRateOutputWithContext(context.Background())
}

func (i *TaxCategoryRate) ToTaxCategoryRateOutputWithContext(ctx context.Context) TaxCategoryRateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaxCategoryRateOutput)
}

func (i *TaxCategoryRate) ToTaxCategoryRatePtrOutput() TaxCategoryRatePtrOutput {
	return i.ToTaxCategoryRatePtrOutputWithContext(context.Background())
}

func (i *TaxCategoryRate) ToTaxCategoryRatePtrOutputWithContext(ctx context.Context) TaxCategoryRatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaxCategoryRatePtrOutput)
}

type TaxCategoryRatePtrInput interface {
	pulumi.Input

	ToTaxCategoryRatePtrOutput() TaxCategoryRatePtrOutput
	ToTaxCategoryRatePtrOutputWithContext(ctx context.Context) TaxCategoryRatePtrOutput
}

type taxCategoryRatePtrType TaxCategoryRateArgs

func (*taxCategoryRatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TaxCategoryRate)(nil))
}

func (i *taxCategoryRatePtrType) ToTaxCategoryRatePtrOutput() TaxCategoryRatePtrOutput {
	return i.ToTaxCategoryRatePtrOutputWithContext(context.Background())
}

func (i *taxCategoryRatePtrType) ToTaxCategoryRatePtrOutputWithContext(ctx context.Context) TaxCategoryRatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaxCategoryRatePtrOutput)
}

// TaxCategoryRateArrayInput is an input type that accepts TaxCategoryRateArray and TaxCategoryRateArrayOutput values.
// You can construct a concrete instance of `TaxCategoryRateArrayInput` via:
//
//          TaxCategoryRateArray{ TaxCategoryRateArgs{...} }
type TaxCategoryRateArrayInput interface {
	pulumi.Input

	ToTaxCategoryRateArrayOutput() TaxCategoryRateArrayOutput
	ToTaxCategoryRateArrayOutputWithContext(context.Context) TaxCategoryRateArrayOutput
}

type TaxCategoryRateArray []TaxCategoryRateInput

func (TaxCategoryRateArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*TaxCategoryRate)(nil))
}

func (i TaxCategoryRateArray) ToTaxCategoryRateArrayOutput() TaxCategoryRateArrayOutput {
	return i.ToTaxCategoryRateArrayOutputWithContext(context.Background())
}

func (i TaxCategoryRateArray) ToTaxCategoryRateArrayOutputWithContext(ctx context.Context) TaxCategoryRateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaxCategoryRateArrayOutput)
}

// TaxCategoryRateMapInput is an input type that accepts TaxCategoryRateMap and TaxCategoryRateMapOutput values.
// You can construct a concrete instance of `TaxCategoryRateMapInput` via:
//
//          TaxCategoryRateMap{ "key": TaxCategoryRateArgs{...} }
type TaxCategoryRateMapInput interface {
	pulumi.Input

	ToTaxCategoryRateMapOutput() TaxCategoryRateMapOutput
	ToTaxCategoryRateMapOutputWithContext(context.Context) TaxCategoryRateMapOutput
}

type TaxCategoryRateMap map[string]TaxCategoryRateInput

func (TaxCategoryRateMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*TaxCategoryRate)(nil))
}

func (i TaxCategoryRateMap) ToTaxCategoryRateMapOutput() TaxCategoryRateMapOutput {
	return i.ToTaxCategoryRateMapOutputWithContext(context.Background())
}

func (i TaxCategoryRateMap) ToTaxCategoryRateMapOutputWithContext(ctx context.Context) TaxCategoryRateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaxCategoryRateMapOutput)
}

type TaxCategoryRateOutput struct {
	*pulumi.OutputState
}

func (TaxCategoryRateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TaxCategoryRate)(nil))
}

func (o TaxCategoryRateOutput) ToTaxCategoryRateOutput() TaxCategoryRateOutput {
	return o
}

func (o TaxCategoryRateOutput) ToTaxCategoryRateOutputWithContext(ctx context.Context) TaxCategoryRateOutput {
	return o
}

func (o TaxCategoryRateOutput) ToTaxCategoryRatePtrOutput() TaxCategoryRatePtrOutput {
	return o.ToTaxCategoryRatePtrOutputWithContext(context.Background())
}

func (o TaxCategoryRateOutput) ToTaxCategoryRatePtrOutputWithContext(ctx context.Context) TaxCategoryRatePtrOutput {
	return o.ApplyT(func(v TaxCategoryRate) *TaxCategoryRate {
		return &v
	}).(TaxCategoryRatePtrOutput)
}

type TaxCategoryRatePtrOutput struct {
	*pulumi.OutputState
}

func (TaxCategoryRatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TaxCategoryRate)(nil))
}

func (o TaxCategoryRatePtrOutput) ToTaxCategoryRatePtrOutput() TaxCategoryRatePtrOutput {
	return o
}

func (o TaxCategoryRatePtrOutput) ToTaxCategoryRatePtrOutputWithContext(ctx context.Context) TaxCategoryRatePtrOutput {
	return o
}

type TaxCategoryRateArrayOutput struct{ *pulumi.OutputState }

func (TaxCategoryRateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TaxCategoryRate)(nil))
}

func (o TaxCategoryRateArrayOutput) ToTaxCategoryRateArrayOutput() TaxCategoryRateArrayOutput {
	return o
}

func (o TaxCategoryRateArrayOutput) ToTaxCategoryRateArrayOutputWithContext(ctx context.Context) TaxCategoryRateArrayOutput {
	return o
}

func (o TaxCategoryRateArrayOutput) Index(i pulumi.IntInput) TaxCategoryRateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TaxCategoryRate {
		return vs[0].([]TaxCategoryRate)[vs[1].(int)]
	}).(TaxCategoryRateOutput)
}

type TaxCategoryRateMapOutput struct{ *pulumi.OutputState }

func (TaxCategoryRateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]TaxCategoryRate)(nil))
}

func (o TaxCategoryRateMapOutput) ToTaxCategoryRateMapOutput() TaxCategoryRateMapOutput {
	return o
}

func (o TaxCategoryRateMapOutput) ToTaxCategoryRateMapOutputWithContext(ctx context.Context) TaxCategoryRateMapOutput {
	return o
}

func (o TaxCategoryRateMapOutput) MapIndex(k pulumi.StringInput) TaxCategoryRateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) TaxCategoryRate {
		return vs[0].(map[string]TaxCategoryRate)[vs[1].(string)]
	}).(TaxCategoryRateOutput)
}

func init() {
	pulumi.RegisterOutputType(TaxCategoryRateOutput{})
	pulumi.RegisterOutputType(TaxCategoryRatePtrOutput{})
	pulumi.RegisterOutputType(TaxCategoryRateArrayOutput{})
	pulumi.RegisterOutputType(TaxCategoryRateMapOutput{})
}
