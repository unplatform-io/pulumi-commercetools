// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package commercetools

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type Subscription struct {
	pulumi.CustomResourceState

	Changes     SubscriptionChangeArrayOutput    `pulumi:"changes"`
	Destination SubscriptionDestinationPtrOutput `pulumi:"destination"`
	Format      SubscriptionFormatPtrOutput      `pulumi:"format"`
	Key         pulumi.StringPtrOutput           `pulumi:"key"`
	Messages    SubscriptionMessageArrayOutput   `pulumi:"messages"`
	Version     pulumi.IntOutput                 `pulumi:"version"`
}

// NewSubscription registers a new resource with the given unique name, arguments, and options.
func NewSubscription(ctx *pulumi.Context,
	name string, args *SubscriptionArgs, opts ...pulumi.ResourceOption) (*Subscription, error) {
	if args == nil {
		args = &SubscriptionArgs{}
	}
	var resource Subscription
	err := ctx.RegisterResource("commercetools:index/subscription:Subscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubscription gets an existing Subscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubscriptionState, opts ...pulumi.ResourceOption) (*Subscription, error) {
	var resource Subscription
	err := ctx.ReadResource("commercetools:index/subscription:Subscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Subscription resources.
type subscriptionState struct {
	Changes     []SubscriptionChange     `pulumi:"changes"`
	Destination *SubscriptionDestination `pulumi:"destination"`
	Format      *SubscriptionFormat      `pulumi:"format"`
	Key         *string                  `pulumi:"key"`
	Messages    []SubscriptionMessage    `pulumi:"messages"`
	Version     *int                     `pulumi:"version"`
}

type SubscriptionState struct {
	Changes     SubscriptionChangeArrayInput
	Destination SubscriptionDestinationPtrInput
	Format      SubscriptionFormatPtrInput
	Key         pulumi.StringPtrInput
	Messages    SubscriptionMessageArrayInput
	Version     pulumi.IntPtrInput
}

func (SubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionState)(nil)).Elem()
}

type subscriptionArgs struct {
	Changes     []SubscriptionChange     `pulumi:"changes"`
	Destination *SubscriptionDestination `pulumi:"destination"`
	Format      *SubscriptionFormat      `pulumi:"format"`
	Key         *string                  `pulumi:"key"`
	Messages    []SubscriptionMessage    `pulumi:"messages"`
}

// The set of arguments for constructing a Subscription resource.
type SubscriptionArgs struct {
	Changes     SubscriptionChangeArrayInput
	Destination SubscriptionDestinationPtrInput
	Format      SubscriptionFormatPtrInput
	Key         pulumi.StringPtrInput
	Messages    SubscriptionMessageArrayInput
}

func (SubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionArgs)(nil)).Elem()
}
