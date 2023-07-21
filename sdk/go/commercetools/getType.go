// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package commercetools

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-commercetools/sdk/go/commercetools/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupType(ctx *pulumi.Context, args *LookupTypeArgs, opts ...pulumi.InvokeOption) (*LookupTypeResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupTypeResult
	err := ctx.Invoke("commercetools:index/getType:getType", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getType.
type LookupTypeArgs struct {
	Key string `pulumi:"key"`
}

// A collection of values returned by getType.
type LookupTypeResult struct {
	Id  string `pulumi:"id"`
	Key string `pulumi:"key"`
}

func LookupTypeOutput(ctx *pulumi.Context, args LookupTypeOutputArgs, opts ...pulumi.InvokeOption) LookupTypeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTypeResult, error) {
			args := v.(LookupTypeArgs)
			r, err := LookupType(ctx, &args, opts...)
			var s LookupTypeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTypeResultOutput)
}

// A collection of arguments for invoking getType.
type LookupTypeOutputArgs struct {
	Key pulumi.StringInput `pulumi:"key"`
}

func (LookupTypeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTypeArgs)(nil)).Elem()
}

// A collection of values returned by getType.
type LookupTypeResultOutput struct{ *pulumi.OutputState }

func (LookupTypeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTypeResult)(nil)).Elem()
}

func (o LookupTypeResultOutput) ToLookupTypeResultOutput() LookupTypeResultOutput {
	return o
}

func (o LookupTypeResultOutput) ToLookupTypeResultOutputWithContext(ctx context.Context) LookupTypeResultOutput {
	return o
}

func (o LookupTypeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTypeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTypeResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTypeResult) string { return v.Key }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTypeResultOutput{})
}