// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package commercetools

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type CustomObject struct {
	pulumi.CustomResourceState

	Container pulumi.StringOutput `pulumi:"container"`
	Key       pulumi.StringOutput `pulumi:"key"`
	Value     pulumi.StringOutput `pulumi:"value"`
	Version   pulumi.IntOutput    `pulumi:"version"`
}

// NewCustomObject registers a new resource with the given unique name, arguments, and options.
func NewCustomObject(ctx *pulumi.Context,
	name string, args *CustomObjectArgs, opts ...pulumi.ResourceOption) (*CustomObject, error) {
	if args == nil || args.Container == nil {
		return nil, errors.New("missing required argument 'Container'")
	}
	if args == nil || args.Key == nil {
		return nil, errors.New("missing required argument 'Key'")
	}
	if args == nil || args.Value == nil {
		return nil, errors.New("missing required argument 'Value'")
	}
	if args == nil {
		args = &CustomObjectArgs{}
	}
	var resource CustomObject
	err := ctx.RegisterResource("commercetools:index/customObject:CustomObject", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomObject gets an existing CustomObject resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomObject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomObjectState, opts ...pulumi.ResourceOption) (*CustomObject, error) {
	var resource CustomObject
	err := ctx.ReadResource("commercetools:index/customObject:CustomObject", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomObject resources.
type customObjectState struct {
	Container *string `pulumi:"container"`
	Key       *string `pulumi:"key"`
	Value     *string `pulumi:"value"`
	Version   *int    `pulumi:"version"`
}

type CustomObjectState struct {
	Container pulumi.StringPtrInput
	Key       pulumi.StringPtrInput
	Value     pulumi.StringPtrInput
	Version   pulumi.IntPtrInput
}

func (CustomObjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*customObjectState)(nil)).Elem()
}

type customObjectArgs struct {
	Container string `pulumi:"container"`
	Key       string `pulumi:"key"`
	Value     string `pulumi:"value"`
}

// The set of arguments for constructing a CustomObject resource.
type CustomObjectArgs struct {
	Container pulumi.StringInput
	Key       pulumi.StringInput
	Value     pulumi.StringInput
}

func (CustomObjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customObjectArgs)(nil)).Elem()
}
