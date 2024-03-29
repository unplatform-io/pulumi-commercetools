// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package commercetools

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-commercetools/sdk/go/commercetools/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CartDiscount struct {
	pulumi.CustomResourceState

	Custom CartDiscountCustomPtrOutput `pulumi:"custom"`
	// [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
	Description pulumi.MapOutput `pulumi:"description"`
	// Only active discount can be applied to the cart
	IsActive pulumi.BoolPtrOutput `pulumi:"isActive"`
	// User-specific unique identifier for a cart discount. Must be unique across a project
	Key pulumi.StringPtrOutput `pulumi:"key"`
	// [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
	Name pulumi.MapOutput `pulumi:"name"`
	// A valid [Cart Predicate](https://docs.commercetools.com/api/projects/predicates#cart-predicates)
	Predicate pulumi.StringOutput `pulumi:"predicate"`
	// States whether the discount can only be used in a connection with a
	// [DiscountCode](https://docs.commercetools.com/api/projects/discountCodes#discountcode)
	RequiresDiscountCode pulumi.BoolPtrOutput `pulumi:"requiresDiscountCode"`
	// The string must contain a number between 0 and 1. All matching cart discounts are applied to a cart in the order defined
	// by this field. A discount with greater sort order is prioritized higher than a discount with lower sort order. The sort
	// order is unambiguous among all cart discounts
	SortOrder pulumi.StringOutput `pulumi:"sortOrder"`
	// Specifies whether the application of this discount causes the following discounts to be ignored. Can be either Stacking
	// or StopAfterThisDiscount
	StackingMode pulumi.StringPtrOutput `pulumi:"stackingMode"`
	// Empty when the value has type giftLineItem, otherwise a
	// [CartDiscountTarget](https://docs.commercetools.com/api/projects/cartDiscounts#cartdiscounttarget)
	Target     CartDiscountTargetPtrOutput `pulumi:"target"`
	ValidFrom  pulumi.StringPtrOutput      `pulumi:"validFrom"`
	ValidUntil pulumi.StringPtrOutput      `pulumi:"validUntil"`
	// Defines the effect the discount will have.
	// [CartDiscountValue](https://docs.commercetools.com/api/projects/cartDiscounts#cartdiscountvalue)
	Value   CartDiscountValueOutput `pulumi:"value"`
	Version pulumi.IntOutput        `pulumi:"version"`
}

// NewCartDiscount registers a new resource with the given unique name, arguments, and options.
func NewCartDiscount(ctx *pulumi.Context,
	name string, args *CartDiscountArgs, opts ...pulumi.ResourceOption) (*CartDiscount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.Predicate == nil {
		return nil, errors.New("invalid value for required argument 'Predicate'")
	}
	if args.SortOrder == nil {
		return nil, errors.New("invalid value for required argument 'SortOrder'")
	}
	if args.Value == nil {
		return nil, errors.New("invalid value for required argument 'Value'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CartDiscount
	err := ctx.RegisterResource("commercetools:index/cartDiscount:CartDiscount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCartDiscount gets an existing CartDiscount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCartDiscount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CartDiscountState, opts ...pulumi.ResourceOption) (*CartDiscount, error) {
	var resource CartDiscount
	err := ctx.ReadResource("commercetools:index/cartDiscount:CartDiscount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CartDiscount resources.
type cartDiscountState struct {
	Custom *CartDiscountCustom `pulumi:"custom"`
	// [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
	Description map[string]interface{} `pulumi:"description"`
	// Only active discount can be applied to the cart
	IsActive *bool `pulumi:"isActive"`
	// User-specific unique identifier for a cart discount. Must be unique across a project
	Key *string `pulumi:"key"`
	// [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
	Name map[string]interface{} `pulumi:"name"`
	// A valid [Cart Predicate](https://docs.commercetools.com/api/projects/predicates#cart-predicates)
	Predicate *string `pulumi:"predicate"`
	// States whether the discount can only be used in a connection with a
	// [DiscountCode](https://docs.commercetools.com/api/projects/discountCodes#discountcode)
	RequiresDiscountCode *bool `pulumi:"requiresDiscountCode"`
	// The string must contain a number between 0 and 1. All matching cart discounts are applied to a cart in the order defined
	// by this field. A discount with greater sort order is prioritized higher than a discount with lower sort order. The sort
	// order is unambiguous among all cart discounts
	SortOrder *string `pulumi:"sortOrder"`
	// Specifies whether the application of this discount causes the following discounts to be ignored. Can be either Stacking
	// or StopAfterThisDiscount
	StackingMode *string `pulumi:"stackingMode"`
	// Empty when the value has type giftLineItem, otherwise a
	// [CartDiscountTarget](https://docs.commercetools.com/api/projects/cartDiscounts#cartdiscounttarget)
	Target     *CartDiscountTarget `pulumi:"target"`
	ValidFrom  *string             `pulumi:"validFrom"`
	ValidUntil *string             `pulumi:"validUntil"`
	// Defines the effect the discount will have.
	// [CartDiscountValue](https://docs.commercetools.com/api/projects/cartDiscounts#cartdiscountvalue)
	Value   *CartDiscountValue `pulumi:"value"`
	Version *int               `pulumi:"version"`
}

type CartDiscountState struct {
	Custom CartDiscountCustomPtrInput
	// [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
	Description pulumi.MapInput
	// Only active discount can be applied to the cart
	IsActive pulumi.BoolPtrInput
	// User-specific unique identifier for a cart discount. Must be unique across a project
	Key pulumi.StringPtrInput
	// [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
	Name pulumi.MapInput
	// A valid [Cart Predicate](https://docs.commercetools.com/api/projects/predicates#cart-predicates)
	Predicate pulumi.StringPtrInput
	// States whether the discount can only be used in a connection with a
	// [DiscountCode](https://docs.commercetools.com/api/projects/discountCodes#discountcode)
	RequiresDiscountCode pulumi.BoolPtrInput
	// The string must contain a number between 0 and 1. All matching cart discounts are applied to a cart in the order defined
	// by this field. A discount with greater sort order is prioritized higher than a discount with lower sort order. The sort
	// order is unambiguous among all cart discounts
	SortOrder pulumi.StringPtrInput
	// Specifies whether the application of this discount causes the following discounts to be ignored. Can be either Stacking
	// or StopAfterThisDiscount
	StackingMode pulumi.StringPtrInput
	// Empty when the value has type giftLineItem, otherwise a
	// [CartDiscountTarget](https://docs.commercetools.com/api/projects/cartDiscounts#cartdiscounttarget)
	Target     CartDiscountTargetPtrInput
	ValidFrom  pulumi.StringPtrInput
	ValidUntil pulumi.StringPtrInput
	// Defines the effect the discount will have.
	// [CartDiscountValue](https://docs.commercetools.com/api/projects/cartDiscounts#cartdiscountvalue)
	Value   CartDiscountValuePtrInput
	Version pulumi.IntPtrInput
}

func (CartDiscountState) ElementType() reflect.Type {
	return reflect.TypeOf((*cartDiscountState)(nil)).Elem()
}

type cartDiscountArgs struct {
	Custom *CartDiscountCustom `pulumi:"custom"`
	// [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
	Description map[string]interface{} `pulumi:"description"`
	// Only active discount can be applied to the cart
	IsActive *bool `pulumi:"isActive"`
	// User-specific unique identifier for a cart discount. Must be unique across a project
	Key *string `pulumi:"key"`
	// [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
	Name map[string]interface{} `pulumi:"name"`
	// A valid [Cart Predicate](https://docs.commercetools.com/api/projects/predicates#cart-predicates)
	Predicate string `pulumi:"predicate"`
	// States whether the discount can only be used in a connection with a
	// [DiscountCode](https://docs.commercetools.com/api/projects/discountCodes#discountcode)
	RequiresDiscountCode *bool `pulumi:"requiresDiscountCode"`
	// The string must contain a number between 0 and 1. All matching cart discounts are applied to a cart in the order defined
	// by this field. A discount with greater sort order is prioritized higher than a discount with lower sort order. The sort
	// order is unambiguous among all cart discounts
	SortOrder string `pulumi:"sortOrder"`
	// Specifies whether the application of this discount causes the following discounts to be ignored. Can be either Stacking
	// or StopAfterThisDiscount
	StackingMode *string `pulumi:"stackingMode"`
	// Empty when the value has type giftLineItem, otherwise a
	// [CartDiscountTarget](https://docs.commercetools.com/api/projects/cartDiscounts#cartdiscounttarget)
	Target     *CartDiscountTarget `pulumi:"target"`
	ValidFrom  *string             `pulumi:"validFrom"`
	ValidUntil *string             `pulumi:"validUntil"`
	// Defines the effect the discount will have.
	// [CartDiscountValue](https://docs.commercetools.com/api/projects/cartDiscounts#cartdiscountvalue)
	Value CartDiscountValue `pulumi:"value"`
}

// The set of arguments for constructing a CartDiscount resource.
type CartDiscountArgs struct {
	Custom CartDiscountCustomPtrInput
	// [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
	Description pulumi.MapInput
	// Only active discount can be applied to the cart
	IsActive pulumi.BoolPtrInput
	// User-specific unique identifier for a cart discount. Must be unique across a project
	Key pulumi.StringPtrInput
	// [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
	Name pulumi.MapInput
	// A valid [Cart Predicate](https://docs.commercetools.com/api/projects/predicates#cart-predicates)
	Predicate pulumi.StringInput
	// States whether the discount can only be used in a connection with a
	// [DiscountCode](https://docs.commercetools.com/api/projects/discountCodes#discountcode)
	RequiresDiscountCode pulumi.BoolPtrInput
	// The string must contain a number between 0 and 1. All matching cart discounts are applied to a cart in the order defined
	// by this field. A discount with greater sort order is prioritized higher than a discount with lower sort order. The sort
	// order is unambiguous among all cart discounts
	SortOrder pulumi.StringInput
	// Specifies whether the application of this discount causes the following discounts to be ignored. Can be either Stacking
	// or StopAfterThisDiscount
	StackingMode pulumi.StringPtrInput
	// Empty when the value has type giftLineItem, otherwise a
	// [CartDiscountTarget](https://docs.commercetools.com/api/projects/cartDiscounts#cartdiscounttarget)
	Target     CartDiscountTargetPtrInput
	ValidFrom  pulumi.StringPtrInput
	ValidUntil pulumi.StringPtrInput
	// Defines the effect the discount will have.
	// [CartDiscountValue](https://docs.commercetools.com/api/projects/cartDiscounts#cartdiscountvalue)
	Value CartDiscountValueInput
}

func (CartDiscountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cartDiscountArgs)(nil)).Elem()
}

type CartDiscountInput interface {
	pulumi.Input

	ToCartDiscountOutput() CartDiscountOutput
	ToCartDiscountOutputWithContext(ctx context.Context) CartDiscountOutput
}

func (*CartDiscount) ElementType() reflect.Type {
	return reflect.TypeOf((**CartDiscount)(nil)).Elem()
}

func (i *CartDiscount) ToCartDiscountOutput() CartDiscountOutput {
	return i.ToCartDiscountOutputWithContext(context.Background())
}

func (i *CartDiscount) ToCartDiscountOutputWithContext(ctx context.Context) CartDiscountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CartDiscountOutput)
}

// CartDiscountArrayInput is an input type that accepts CartDiscountArray and CartDiscountArrayOutput values.
// You can construct a concrete instance of `CartDiscountArrayInput` via:
//
//	CartDiscountArray{ CartDiscountArgs{...} }
type CartDiscountArrayInput interface {
	pulumi.Input

	ToCartDiscountArrayOutput() CartDiscountArrayOutput
	ToCartDiscountArrayOutputWithContext(context.Context) CartDiscountArrayOutput
}

type CartDiscountArray []CartDiscountInput

func (CartDiscountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CartDiscount)(nil)).Elem()
}

func (i CartDiscountArray) ToCartDiscountArrayOutput() CartDiscountArrayOutput {
	return i.ToCartDiscountArrayOutputWithContext(context.Background())
}

func (i CartDiscountArray) ToCartDiscountArrayOutputWithContext(ctx context.Context) CartDiscountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CartDiscountArrayOutput)
}

// CartDiscountMapInput is an input type that accepts CartDiscountMap and CartDiscountMapOutput values.
// You can construct a concrete instance of `CartDiscountMapInput` via:
//
//	CartDiscountMap{ "key": CartDiscountArgs{...} }
type CartDiscountMapInput interface {
	pulumi.Input

	ToCartDiscountMapOutput() CartDiscountMapOutput
	ToCartDiscountMapOutputWithContext(context.Context) CartDiscountMapOutput
}

type CartDiscountMap map[string]CartDiscountInput

func (CartDiscountMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CartDiscount)(nil)).Elem()
}

func (i CartDiscountMap) ToCartDiscountMapOutput() CartDiscountMapOutput {
	return i.ToCartDiscountMapOutputWithContext(context.Background())
}

func (i CartDiscountMap) ToCartDiscountMapOutputWithContext(ctx context.Context) CartDiscountMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CartDiscountMapOutput)
}

type CartDiscountOutput struct{ *pulumi.OutputState }

func (CartDiscountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CartDiscount)(nil)).Elem()
}

func (o CartDiscountOutput) ToCartDiscountOutput() CartDiscountOutput {
	return o
}

func (o CartDiscountOutput) ToCartDiscountOutputWithContext(ctx context.Context) CartDiscountOutput {
	return o
}

func (o CartDiscountOutput) Custom() CartDiscountCustomPtrOutput {
	return o.ApplyT(func(v *CartDiscount) CartDiscountCustomPtrOutput { return v.Custom }).(CartDiscountCustomPtrOutput)
}

// [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
func (o CartDiscountOutput) Description() pulumi.MapOutput {
	return o.ApplyT(func(v *CartDiscount) pulumi.MapOutput { return v.Description }).(pulumi.MapOutput)
}

// Only active discount can be applied to the cart
func (o CartDiscountOutput) IsActive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CartDiscount) pulumi.BoolPtrOutput { return v.IsActive }).(pulumi.BoolPtrOutput)
}

// User-specific unique identifier for a cart discount. Must be unique across a project
func (o CartDiscountOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CartDiscount) pulumi.StringPtrOutput { return v.Key }).(pulumi.StringPtrOutput)
}

// [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
func (o CartDiscountOutput) Name() pulumi.MapOutput {
	return o.ApplyT(func(v *CartDiscount) pulumi.MapOutput { return v.Name }).(pulumi.MapOutput)
}

// A valid [Cart Predicate](https://docs.commercetools.com/api/projects/predicates#cart-predicates)
func (o CartDiscountOutput) Predicate() pulumi.StringOutput {
	return o.ApplyT(func(v *CartDiscount) pulumi.StringOutput { return v.Predicate }).(pulumi.StringOutput)
}

// States whether the discount can only be used in a connection with a
// [DiscountCode](https://docs.commercetools.com/api/projects/discountCodes#discountcode)
func (o CartDiscountOutput) RequiresDiscountCode() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CartDiscount) pulumi.BoolPtrOutput { return v.RequiresDiscountCode }).(pulumi.BoolPtrOutput)
}

// The string must contain a number between 0 and 1. All matching cart discounts are applied to a cart in the order defined
// by this field. A discount with greater sort order is prioritized higher than a discount with lower sort order. The sort
// order is unambiguous among all cart discounts
func (o CartDiscountOutput) SortOrder() pulumi.StringOutput {
	return o.ApplyT(func(v *CartDiscount) pulumi.StringOutput { return v.SortOrder }).(pulumi.StringOutput)
}

// Specifies whether the application of this discount causes the following discounts to be ignored. Can be either Stacking
// or StopAfterThisDiscount
func (o CartDiscountOutput) StackingMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CartDiscount) pulumi.StringPtrOutput { return v.StackingMode }).(pulumi.StringPtrOutput)
}

// Empty when the value has type giftLineItem, otherwise a
// [CartDiscountTarget](https://docs.commercetools.com/api/projects/cartDiscounts#cartdiscounttarget)
func (o CartDiscountOutput) Target() CartDiscountTargetPtrOutput {
	return o.ApplyT(func(v *CartDiscount) CartDiscountTargetPtrOutput { return v.Target }).(CartDiscountTargetPtrOutput)
}

func (o CartDiscountOutput) ValidFrom() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CartDiscount) pulumi.StringPtrOutput { return v.ValidFrom }).(pulumi.StringPtrOutput)
}

func (o CartDiscountOutput) ValidUntil() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CartDiscount) pulumi.StringPtrOutput { return v.ValidUntil }).(pulumi.StringPtrOutput)
}

// Defines the effect the discount will have.
// [CartDiscountValue](https://docs.commercetools.com/api/projects/cartDiscounts#cartdiscountvalue)
func (o CartDiscountOutput) Value() CartDiscountValueOutput {
	return o.ApplyT(func(v *CartDiscount) CartDiscountValueOutput { return v.Value }).(CartDiscountValueOutput)
}

func (o CartDiscountOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v *CartDiscount) pulumi.IntOutput { return v.Version }).(pulumi.IntOutput)
}

type CartDiscountArrayOutput struct{ *pulumi.OutputState }

func (CartDiscountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CartDiscount)(nil)).Elem()
}

func (o CartDiscountArrayOutput) ToCartDiscountArrayOutput() CartDiscountArrayOutput {
	return o
}

func (o CartDiscountArrayOutput) ToCartDiscountArrayOutputWithContext(ctx context.Context) CartDiscountArrayOutput {
	return o
}

func (o CartDiscountArrayOutput) Index(i pulumi.IntInput) CartDiscountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CartDiscount {
		return vs[0].([]*CartDiscount)[vs[1].(int)]
	}).(CartDiscountOutput)
}

type CartDiscountMapOutput struct{ *pulumi.OutputState }

func (CartDiscountMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CartDiscount)(nil)).Elem()
}

func (o CartDiscountMapOutput) ToCartDiscountMapOutput() CartDiscountMapOutput {
	return o
}

func (o CartDiscountMapOutput) ToCartDiscountMapOutputWithContext(ctx context.Context) CartDiscountMapOutput {
	return o
}

func (o CartDiscountMapOutput) MapIndex(k pulumi.StringInput) CartDiscountOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CartDiscount {
		return vs[0].(map[string]*CartDiscount)[vs[1].(string)]
	}).(CartDiscountOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CartDiscountInput)(nil)).Elem(), &CartDiscount{})
	pulumi.RegisterInputType(reflect.TypeOf((*CartDiscountArrayInput)(nil)).Elem(), CartDiscountArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CartDiscountMapInput)(nil)).Elem(), CartDiscountMap{})
	pulumi.RegisterOutputType(CartDiscountOutput{})
	pulumi.RegisterOutputType(CartDiscountArrayOutput{})
	pulumi.RegisterOutputType(CartDiscountMapOutput{})
}
