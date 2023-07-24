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

type Channel struct {
	pulumi.CustomResourceState

	Address ChannelAddressPtrOutput `pulumi:"address"`
	Custom  ChannelCustomPtrOutput  `pulumi:"custom"`
	// [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
	Description pulumi.MapOutput            `pulumi:"description"`
	Geolocation ChannelGeolocationPtrOutput `pulumi:"geolocation"`
	// Any arbitrary string key that uniquely identifies this channel within the project
	Key pulumi.StringOutput `pulumi:"key"`
	// [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
	Name pulumi.MapOutput `pulumi:"name"`
	// The [roles](https://docs.commercetools.com/api/projects/channels#channelroleenum) of this channel. Each channel must
	// have at least one role
	Roles   pulumi.StringArrayOutput `pulumi:"roles"`
	Version pulumi.IntOutput         `pulumi:"version"`
}

// NewChannel registers a new resource with the given unique name, arguments, and options.
func NewChannel(ctx *pulumi.Context,
	name string, args *ChannelArgs, opts ...pulumi.ResourceOption) (*Channel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Roles == nil {
		return nil, errors.New("invalid value for required argument 'Roles'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Channel
	err := ctx.RegisterResource("commercetools:index/channel:Channel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetChannel gets an existing Channel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ChannelState, opts ...pulumi.ResourceOption) (*Channel, error) {
	var resource Channel
	err := ctx.ReadResource("commercetools:index/channel:Channel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Channel resources.
type channelState struct {
	Address *ChannelAddress `pulumi:"address"`
	Custom  *ChannelCustom  `pulumi:"custom"`
	// [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
	Description map[string]interface{} `pulumi:"description"`
	Geolocation *ChannelGeolocation    `pulumi:"geolocation"`
	// Any arbitrary string key that uniquely identifies this channel within the project
	Key *string `pulumi:"key"`
	// [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
	Name map[string]interface{} `pulumi:"name"`
	// The [roles](https://docs.commercetools.com/api/projects/channels#channelroleenum) of this channel. Each channel must
	// have at least one role
	Roles   []string `pulumi:"roles"`
	Version *int     `pulumi:"version"`
}

type ChannelState struct {
	Address ChannelAddressPtrInput
	Custom  ChannelCustomPtrInput
	// [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
	Description pulumi.MapInput
	Geolocation ChannelGeolocationPtrInput
	// Any arbitrary string key that uniquely identifies this channel within the project
	Key pulumi.StringPtrInput
	// [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
	Name pulumi.MapInput
	// The [roles](https://docs.commercetools.com/api/projects/channels#channelroleenum) of this channel. Each channel must
	// have at least one role
	Roles   pulumi.StringArrayInput
	Version pulumi.IntPtrInput
}

func (ChannelState) ElementType() reflect.Type {
	return reflect.TypeOf((*channelState)(nil)).Elem()
}

type channelArgs struct {
	Address *ChannelAddress `pulumi:"address"`
	Custom  *ChannelCustom  `pulumi:"custom"`
	// [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
	Description map[string]interface{} `pulumi:"description"`
	Geolocation *ChannelGeolocation    `pulumi:"geolocation"`
	// Any arbitrary string key that uniquely identifies this channel within the project
	Key string `pulumi:"key"`
	// [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
	Name map[string]interface{} `pulumi:"name"`
	// The [roles](https://docs.commercetools.com/api/projects/channels#channelroleenum) of this channel. Each channel must
	// have at least one role
	Roles []string `pulumi:"roles"`
}

// The set of arguments for constructing a Channel resource.
type ChannelArgs struct {
	Address ChannelAddressPtrInput
	Custom  ChannelCustomPtrInput
	// [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
	Description pulumi.MapInput
	Geolocation ChannelGeolocationPtrInput
	// Any arbitrary string key that uniquely identifies this channel within the project
	Key pulumi.StringInput
	// [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
	Name pulumi.MapInput
	// The [roles](https://docs.commercetools.com/api/projects/channels#channelroleenum) of this channel. Each channel must
	// have at least one role
	Roles pulumi.StringArrayInput
}

func (ChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*channelArgs)(nil)).Elem()
}

type ChannelInput interface {
	pulumi.Input

	ToChannelOutput() ChannelOutput
	ToChannelOutputWithContext(ctx context.Context) ChannelOutput
}

func (*Channel) ElementType() reflect.Type {
	return reflect.TypeOf((**Channel)(nil)).Elem()
}

func (i *Channel) ToChannelOutput() ChannelOutput {
	return i.ToChannelOutputWithContext(context.Background())
}

func (i *Channel) ToChannelOutputWithContext(ctx context.Context) ChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChannelOutput)
}

// ChannelArrayInput is an input type that accepts ChannelArray and ChannelArrayOutput values.
// You can construct a concrete instance of `ChannelArrayInput` via:
//
//	ChannelArray{ ChannelArgs{...} }
type ChannelArrayInput interface {
	pulumi.Input

	ToChannelArrayOutput() ChannelArrayOutput
	ToChannelArrayOutputWithContext(context.Context) ChannelArrayOutput
}

type ChannelArray []ChannelInput

func (ChannelArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Channel)(nil)).Elem()
}

func (i ChannelArray) ToChannelArrayOutput() ChannelArrayOutput {
	return i.ToChannelArrayOutputWithContext(context.Background())
}

func (i ChannelArray) ToChannelArrayOutputWithContext(ctx context.Context) ChannelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChannelArrayOutput)
}

// ChannelMapInput is an input type that accepts ChannelMap and ChannelMapOutput values.
// You can construct a concrete instance of `ChannelMapInput` via:
//
//	ChannelMap{ "key": ChannelArgs{...} }
type ChannelMapInput interface {
	pulumi.Input

	ToChannelMapOutput() ChannelMapOutput
	ToChannelMapOutputWithContext(context.Context) ChannelMapOutput
}

type ChannelMap map[string]ChannelInput

func (ChannelMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Channel)(nil)).Elem()
}

func (i ChannelMap) ToChannelMapOutput() ChannelMapOutput {
	return i.ToChannelMapOutputWithContext(context.Background())
}

func (i ChannelMap) ToChannelMapOutputWithContext(ctx context.Context) ChannelMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChannelMapOutput)
}

type ChannelOutput struct{ *pulumi.OutputState }

func (ChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Channel)(nil)).Elem()
}

func (o ChannelOutput) ToChannelOutput() ChannelOutput {
	return o
}

func (o ChannelOutput) ToChannelOutputWithContext(ctx context.Context) ChannelOutput {
	return o
}

func (o ChannelOutput) Address() ChannelAddressPtrOutput {
	return o.ApplyT(func(v *Channel) ChannelAddressPtrOutput { return v.Address }).(ChannelAddressPtrOutput)
}

func (o ChannelOutput) Custom() ChannelCustomPtrOutput {
	return o.ApplyT(func(v *Channel) ChannelCustomPtrOutput { return v.Custom }).(ChannelCustomPtrOutput)
}

// [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
func (o ChannelOutput) Description() pulumi.MapOutput {
	return o.ApplyT(func(v *Channel) pulumi.MapOutput { return v.Description }).(pulumi.MapOutput)
}

func (o ChannelOutput) Geolocation() ChannelGeolocationPtrOutput {
	return o.ApplyT(func(v *Channel) ChannelGeolocationPtrOutput { return v.Geolocation }).(ChannelGeolocationPtrOutput)
}

// Any arbitrary string key that uniquely identifies this channel within the project
func (o ChannelOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
func (o ChannelOutput) Name() pulumi.MapOutput {
	return o.ApplyT(func(v *Channel) pulumi.MapOutput { return v.Name }).(pulumi.MapOutput)
}

// The [roles](https://docs.commercetools.com/api/projects/channels#channelroleenum) of this channel. Each channel must
// have at least one role
func (o ChannelOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringArrayOutput { return v.Roles }).(pulumi.StringArrayOutput)
}

func (o ChannelOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v *Channel) pulumi.IntOutput { return v.Version }).(pulumi.IntOutput)
}

type ChannelArrayOutput struct{ *pulumi.OutputState }

func (ChannelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Channel)(nil)).Elem()
}

func (o ChannelArrayOutput) ToChannelArrayOutput() ChannelArrayOutput {
	return o
}

func (o ChannelArrayOutput) ToChannelArrayOutputWithContext(ctx context.Context) ChannelArrayOutput {
	return o
}

func (o ChannelArrayOutput) Index(i pulumi.IntInput) ChannelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Channel {
		return vs[0].([]*Channel)[vs[1].(int)]
	}).(ChannelOutput)
}

type ChannelMapOutput struct{ *pulumi.OutputState }

func (ChannelMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Channel)(nil)).Elem()
}

func (o ChannelMapOutput) ToChannelMapOutput() ChannelMapOutput {
	return o
}

func (o ChannelMapOutput) ToChannelMapOutputWithContext(ctx context.Context) ChannelMapOutput {
	return o
}

func (o ChannelMapOutput) MapIndex(k pulumi.StringInput) ChannelOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Channel {
		return vs[0].(map[string]*Channel)[vs[1].(string)]
	}).(ChannelOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ChannelInput)(nil)).Elem(), &Channel{})
	pulumi.RegisterInputType(reflect.TypeOf((*ChannelArrayInput)(nil)).Elem(), ChannelArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ChannelMapInput)(nil)).Elem(), ChannelMap{})
	pulumi.RegisterOutputType(ChannelOutput{})
	pulumi.RegisterOutputType(ChannelArrayOutput{})
	pulumi.RegisterOutputType(ChannelMapOutput{})
}
