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

type Category struct {
	pulumi.CustomResourceState

	// Can be used to store images, icons or movies related to this category
	Assets      CategoryAssetArrayOutput `pulumi:"assets"`
	Custom      CategoryCustomPtrOutput  `pulumi:"custom"`
	Description pulumi.MapOutput         `pulumi:"description"`
	ExternalId  pulumi.StringPtrOutput   `pulumi:"externalId"`
	// Category-specific unique identifier. Must be unique across a project
	Key             pulumi.StringPtrOutput `pulumi:"key"`
	MetaDescription pulumi.MapOutput       `pulumi:"metaDescription"`
	MetaKeywords    pulumi.MapOutput       `pulumi:"metaKeywords"`
	MetaTitle       pulumi.MapOutput       `pulumi:"metaTitle"`
	Name            pulumi.MapOutput       `pulumi:"name"`
	// An attribute as base for a custom category order in one level, filled with random value when left empty
	OrderHint pulumi.StringPtrOutput `pulumi:"orderHint"`
	// A category that is the parent of this category in the category tree
	Parent pulumi.StringPtrOutput `pulumi:"parent"`
	// Human readable identifiers, needs to be unique
	Slug    pulumi.MapOutput `pulumi:"slug"`
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewCategory registers a new resource with the given unique name, arguments, and options.
func NewCategory(ctx *pulumi.Context,
	name string, args *CategoryArgs, opts ...pulumi.ResourceOption) (*Category, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.Slug == nil {
		return nil, errors.New("invalid value for required argument 'Slug'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Category
	err := ctx.RegisterResource("commercetools:index/category:Category", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCategory gets an existing Category resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCategory(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CategoryState, opts ...pulumi.ResourceOption) (*Category, error) {
	var resource Category
	err := ctx.ReadResource("commercetools:index/category:Category", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Category resources.
type categoryState struct {
	// Can be used to store images, icons or movies related to this category
	Assets      []CategoryAsset        `pulumi:"assets"`
	Custom      *CategoryCustom        `pulumi:"custom"`
	Description map[string]interface{} `pulumi:"description"`
	ExternalId  *string                `pulumi:"externalId"`
	// Category-specific unique identifier. Must be unique across a project
	Key             *string                `pulumi:"key"`
	MetaDescription map[string]interface{} `pulumi:"metaDescription"`
	MetaKeywords    map[string]interface{} `pulumi:"metaKeywords"`
	MetaTitle       map[string]interface{} `pulumi:"metaTitle"`
	Name            map[string]interface{} `pulumi:"name"`
	// An attribute as base for a custom category order in one level, filled with random value when left empty
	OrderHint *string `pulumi:"orderHint"`
	// A category that is the parent of this category in the category tree
	Parent *string `pulumi:"parent"`
	// Human readable identifiers, needs to be unique
	Slug    map[string]interface{} `pulumi:"slug"`
	Version *int                   `pulumi:"version"`
}

type CategoryState struct {
	// Can be used to store images, icons or movies related to this category
	Assets      CategoryAssetArrayInput
	Custom      CategoryCustomPtrInput
	Description pulumi.MapInput
	ExternalId  pulumi.StringPtrInput
	// Category-specific unique identifier. Must be unique across a project
	Key             pulumi.StringPtrInput
	MetaDescription pulumi.MapInput
	MetaKeywords    pulumi.MapInput
	MetaTitle       pulumi.MapInput
	Name            pulumi.MapInput
	// An attribute as base for a custom category order in one level, filled with random value when left empty
	OrderHint pulumi.StringPtrInput
	// A category that is the parent of this category in the category tree
	Parent pulumi.StringPtrInput
	// Human readable identifiers, needs to be unique
	Slug    pulumi.MapInput
	Version pulumi.IntPtrInput
}

func (CategoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*categoryState)(nil)).Elem()
}

type categoryArgs struct {
	// Can be used to store images, icons or movies related to this category
	Assets      []CategoryAsset        `pulumi:"assets"`
	Custom      *CategoryCustom        `pulumi:"custom"`
	Description map[string]interface{} `pulumi:"description"`
	ExternalId  *string                `pulumi:"externalId"`
	// Category-specific unique identifier. Must be unique across a project
	Key             *string                `pulumi:"key"`
	MetaDescription map[string]interface{} `pulumi:"metaDescription"`
	MetaKeywords    map[string]interface{} `pulumi:"metaKeywords"`
	MetaTitle       map[string]interface{} `pulumi:"metaTitle"`
	Name            map[string]interface{} `pulumi:"name"`
	// An attribute as base for a custom category order in one level, filled with random value when left empty
	OrderHint *string `pulumi:"orderHint"`
	// A category that is the parent of this category in the category tree
	Parent *string `pulumi:"parent"`
	// Human readable identifiers, needs to be unique
	Slug map[string]interface{} `pulumi:"slug"`
}

// The set of arguments for constructing a Category resource.
type CategoryArgs struct {
	// Can be used to store images, icons or movies related to this category
	Assets      CategoryAssetArrayInput
	Custom      CategoryCustomPtrInput
	Description pulumi.MapInput
	ExternalId  pulumi.StringPtrInput
	// Category-specific unique identifier. Must be unique across a project
	Key             pulumi.StringPtrInput
	MetaDescription pulumi.MapInput
	MetaKeywords    pulumi.MapInput
	MetaTitle       pulumi.MapInput
	Name            pulumi.MapInput
	// An attribute as base for a custom category order in one level, filled with random value when left empty
	OrderHint pulumi.StringPtrInput
	// A category that is the parent of this category in the category tree
	Parent pulumi.StringPtrInput
	// Human readable identifiers, needs to be unique
	Slug pulumi.MapInput
}

func (CategoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*categoryArgs)(nil)).Elem()
}

type CategoryInput interface {
	pulumi.Input

	ToCategoryOutput() CategoryOutput
	ToCategoryOutputWithContext(ctx context.Context) CategoryOutput
}

func (*Category) ElementType() reflect.Type {
	return reflect.TypeOf((**Category)(nil)).Elem()
}

func (i *Category) ToCategoryOutput() CategoryOutput {
	return i.ToCategoryOutputWithContext(context.Background())
}

func (i *Category) ToCategoryOutputWithContext(ctx context.Context) CategoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CategoryOutput)
}

// CategoryArrayInput is an input type that accepts CategoryArray and CategoryArrayOutput values.
// You can construct a concrete instance of `CategoryArrayInput` via:
//
//	CategoryArray{ CategoryArgs{...} }
type CategoryArrayInput interface {
	pulumi.Input

	ToCategoryArrayOutput() CategoryArrayOutput
	ToCategoryArrayOutputWithContext(context.Context) CategoryArrayOutput
}

type CategoryArray []CategoryInput

func (CategoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Category)(nil)).Elem()
}

func (i CategoryArray) ToCategoryArrayOutput() CategoryArrayOutput {
	return i.ToCategoryArrayOutputWithContext(context.Background())
}

func (i CategoryArray) ToCategoryArrayOutputWithContext(ctx context.Context) CategoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CategoryArrayOutput)
}

// CategoryMapInput is an input type that accepts CategoryMap and CategoryMapOutput values.
// You can construct a concrete instance of `CategoryMapInput` via:
//
//	CategoryMap{ "key": CategoryArgs{...} }
type CategoryMapInput interface {
	pulumi.Input

	ToCategoryMapOutput() CategoryMapOutput
	ToCategoryMapOutputWithContext(context.Context) CategoryMapOutput
}

type CategoryMap map[string]CategoryInput

func (CategoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Category)(nil)).Elem()
}

func (i CategoryMap) ToCategoryMapOutput() CategoryMapOutput {
	return i.ToCategoryMapOutputWithContext(context.Background())
}

func (i CategoryMap) ToCategoryMapOutputWithContext(ctx context.Context) CategoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CategoryMapOutput)
}

type CategoryOutput struct{ *pulumi.OutputState }

func (CategoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Category)(nil)).Elem()
}

func (o CategoryOutput) ToCategoryOutput() CategoryOutput {
	return o
}

func (o CategoryOutput) ToCategoryOutputWithContext(ctx context.Context) CategoryOutput {
	return o
}

// Can be used to store images, icons or movies related to this category
func (o CategoryOutput) Assets() CategoryAssetArrayOutput {
	return o.ApplyT(func(v *Category) CategoryAssetArrayOutput { return v.Assets }).(CategoryAssetArrayOutput)
}

func (o CategoryOutput) Custom() CategoryCustomPtrOutput {
	return o.ApplyT(func(v *Category) CategoryCustomPtrOutput { return v.Custom }).(CategoryCustomPtrOutput)
}

func (o CategoryOutput) Description() pulumi.MapOutput {
	return o.ApplyT(func(v *Category) pulumi.MapOutput { return v.Description }).(pulumi.MapOutput)
}

func (o CategoryOutput) ExternalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Category) pulumi.StringPtrOutput { return v.ExternalId }).(pulumi.StringPtrOutput)
}

// Category-specific unique identifier. Must be unique across a project
func (o CategoryOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Category) pulumi.StringPtrOutput { return v.Key }).(pulumi.StringPtrOutput)
}

func (o CategoryOutput) MetaDescription() pulumi.MapOutput {
	return o.ApplyT(func(v *Category) pulumi.MapOutput { return v.MetaDescription }).(pulumi.MapOutput)
}

func (o CategoryOutput) MetaKeywords() pulumi.MapOutput {
	return o.ApplyT(func(v *Category) pulumi.MapOutput { return v.MetaKeywords }).(pulumi.MapOutput)
}

func (o CategoryOutput) MetaTitle() pulumi.MapOutput {
	return o.ApplyT(func(v *Category) pulumi.MapOutput { return v.MetaTitle }).(pulumi.MapOutput)
}

func (o CategoryOutput) Name() pulumi.MapOutput {
	return o.ApplyT(func(v *Category) pulumi.MapOutput { return v.Name }).(pulumi.MapOutput)
}

// An attribute as base for a custom category order in one level, filled with random value when left empty
func (o CategoryOutput) OrderHint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Category) pulumi.StringPtrOutput { return v.OrderHint }).(pulumi.StringPtrOutput)
}

// A category that is the parent of this category in the category tree
func (o CategoryOutput) Parent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Category) pulumi.StringPtrOutput { return v.Parent }).(pulumi.StringPtrOutput)
}

// Human readable identifiers, needs to be unique
func (o CategoryOutput) Slug() pulumi.MapOutput {
	return o.ApplyT(func(v *Category) pulumi.MapOutput { return v.Slug }).(pulumi.MapOutput)
}

func (o CategoryOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v *Category) pulumi.IntOutput { return v.Version }).(pulumi.IntOutput)
}

type CategoryArrayOutput struct{ *pulumi.OutputState }

func (CategoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Category)(nil)).Elem()
}

func (o CategoryArrayOutput) ToCategoryArrayOutput() CategoryArrayOutput {
	return o
}

func (o CategoryArrayOutput) ToCategoryArrayOutputWithContext(ctx context.Context) CategoryArrayOutput {
	return o
}

func (o CategoryArrayOutput) Index(i pulumi.IntInput) CategoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Category {
		return vs[0].([]*Category)[vs[1].(int)]
	}).(CategoryOutput)
}

type CategoryMapOutput struct{ *pulumi.OutputState }

func (CategoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Category)(nil)).Elem()
}

func (o CategoryMapOutput) ToCategoryMapOutput() CategoryMapOutput {
	return o
}

func (o CategoryMapOutput) ToCategoryMapOutputWithContext(ctx context.Context) CategoryMapOutput {
	return o
}

func (o CategoryMapOutput) MapIndex(k pulumi.StringInput) CategoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Category {
		return vs[0].(map[string]*Category)[vs[1].(string)]
	}).(CategoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CategoryInput)(nil)).Elem(), &Category{})
	pulumi.RegisterInputType(reflect.TypeOf((*CategoryArrayInput)(nil)).Elem(), CategoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CategoryMapInput)(nil)).Elem(), CategoryMap{})
	pulumi.RegisterOutputType(CategoryOutput{})
	pulumi.RegisterOutputType(CategoryArrayOutput{})
	pulumi.RegisterOutputType(CategoryMapOutput{})
}
