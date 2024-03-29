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

type StateTransitions struct {
	pulumi.CustomResourceState

	// ID of the state to transition from
	From pulumi.StringOutput `pulumi:"from"`
	// Transitions are a way to describe possible transformations of the current state to other states of the same type (for
	// example: Initial -> Shipped). When performing a transitionState update action and transitions is set, the currently
	// referenced state must have a transition to the new state. If transitions is an empty list, it means the current state is
	// a final state and no further transitions are allowed. If transitions is not set, the validation is turned off. When
	// performing a transitionState update action, any other state of the same type can be transitioned to
	Tos pulumi.StringArrayOutput `pulumi:"tos"`
}

// NewStateTransitions registers a new resource with the given unique name, arguments, and options.
func NewStateTransitions(ctx *pulumi.Context,
	name string, args *StateTransitionsArgs, opts ...pulumi.ResourceOption) (*StateTransitions, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.From == nil {
		return nil, errors.New("invalid value for required argument 'From'")
	}
	if args.Tos == nil {
		return nil, errors.New("invalid value for required argument 'Tos'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource StateTransitions
	err := ctx.RegisterResource("commercetools:index/stateTransitions:StateTransitions", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStateTransitions gets an existing StateTransitions resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStateTransitions(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StateTransitionsState, opts ...pulumi.ResourceOption) (*StateTransitions, error) {
	var resource StateTransitions
	err := ctx.ReadResource("commercetools:index/stateTransitions:StateTransitions", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StateTransitions resources.
type stateTransitionsState struct {
	// ID of the state to transition from
	From *string `pulumi:"from"`
	// Transitions are a way to describe possible transformations of the current state to other states of the same type (for
	// example: Initial -> Shipped). When performing a transitionState update action and transitions is set, the currently
	// referenced state must have a transition to the new state. If transitions is an empty list, it means the current state is
	// a final state and no further transitions are allowed. If transitions is not set, the validation is turned off. When
	// performing a transitionState update action, any other state of the same type can be transitioned to
	Tos []string `pulumi:"tos"`
}

type StateTransitionsState struct {
	// ID of the state to transition from
	From pulumi.StringPtrInput
	// Transitions are a way to describe possible transformations of the current state to other states of the same type (for
	// example: Initial -> Shipped). When performing a transitionState update action and transitions is set, the currently
	// referenced state must have a transition to the new state. If transitions is an empty list, it means the current state is
	// a final state and no further transitions are allowed. If transitions is not set, the validation is turned off. When
	// performing a transitionState update action, any other state of the same type can be transitioned to
	Tos pulumi.StringArrayInput
}

func (StateTransitionsState) ElementType() reflect.Type {
	return reflect.TypeOf((*stateTransitionsState)(nil)).Elem()
}

type stateTransitionsArgs struct {
	// ID of the state to transition from
	From string `pulumi:"from"`
	// Transitions are a way to describe possible transformations of the current state to other states of the same type (for
	// example: Initial -> Shipped). When performing a transitionState update action and transitions is set, the currently
	// referenced state must have a transition to the new state. If transitions is an empty list, it means the current state is
	// a final state and no further transitions are allowed. If transitions is not set, the validation is turned off. When
	// performing a transitionState update action, any other state of the same type can be transitioned to
	Tos []string `pulumi:"tos"`
}

// The set of arguments for constructing a StateTransitions resource.
type StateTransitionsArgs struct {
	// ID of the state to transition from
	From pulumi.StringInput
	// Transitions are a way to describe possible transformations of the current state to other states of the same type (for
	// example: Initial -> Shipped). When performing a transitionState update action and transitions is set, the currently
	// referenced state must have a transition to the new state. If transitions is an empty list, it means the current state is
	// a final state and no further transitions are allowed. If transitions is not set, the validation is turned off. When
	// performing a transitionState update action, any other state of the same type can be transitioned to
	Tos pulumi.StringArrayInput
}

func (StateTransitionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*stateTransitionsArgs)(nil)).Elem()
}

type StateTransitionsInput interface {
	pulumi.Input

	ToStateTransitionsOutput() StateTransitionsOutput
	ToStateTransitionsOutputWithContext(ctx context.Context) StateTransitionsOutput
}

func (*StateTransitions) ElementType() reflect.Type {
	return reflect.TypeOf((**StateTransitions)(nil)).Elem()
}

func (i *StateTransitions) ToStateTransitionsOutput() StateTransitionsOutput {
	return i.ToStateTransitionsOutputWithContext(context.Background())
}

func (i *StateTransitions) ToStateTransitionsOutputWithContext(ctx context.Context) StateTransitionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StateTransitionsOutput)
}

// StateTransitionsArrayInput is an input type that accepts StateTransitionsArray and StateTransitionsArrayOutput values.
// You can construct a concrete instance of `StateTransitionsArrayInput` via:
//
//	StateTransitionsArray{ StateTransitionsArgs{...} }
type StateTransitionsArrayInput interface {
	pulumi.Input

	ToStateTransitionsArrayOutput() StateTransitionsArrayOutput
	ToStateTransitionsArrayOutputWithContext(context.Context) StateTransitionsArrayOutput
}

type StateTransitionsArray []StateTransitionsInput

func (StateTransitionsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StateTransitions)(nil)).Elem()
}

func (i StateTransitionsArray) ToStateTransitionsArrayOutput() StateTransitionsArrayOutput {
	return i.ToStateTransitionsArrayOutputWithContext(context.Background())
}

func (i StateTransitionsArray) ToStateTransitionsArrayOutputWithContext(ctx context.Context) StateTransitionsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StateTransitionsArrayOutput)
}

// StateTransitionsMapInput is an input type that accepts StateTransitionsMap and StateTransitionsMapOutput values.
// You can construct a concrete instance of `StateTransitionsMapInput` via:
//
//	StateTransitionsMap{ "key": StateTransitionsArgs{...} }
type StateTransitionsMapInput interface {
	pulumi.Input

	ToStateTransitionsMapOutput() StateTransitionsMapOutput
	ToStateTransitionsMapOutputWithContext(context.Context) StateTransitionsMapOutput
}

type StateTransitionsMap map[string]StateTransitionsInput

func (StateTransitionsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StateTransitions)(nil)).Elem()
}

func (i StateTransitionsMap) ToStateTransitionsMapOutput() StateTransitionsMapOutput {
	return i.ToStateTransitionsMapOutputWithContext(context.Background())
}

func (i StateTransitionsMap) ToStateTransitionsMapOutputWithContext(ctx context.Context) StateTransitionsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StateTransitionsMapOutput)
}

type StateTransitionsOutput struct{ *pulumi.OutputState }

func (StateTransitionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StateTransitions)(nil)).Elem()
}

func (o StateTransitionsOutput) ToStateTransitionsOutput() StateTransitionsOutput {
	return o
}

func (o StateTransitionsOutput) ToStateTransitionsOutputWithContext(ctx context.Context) StateTransitionsOutput {
	return o
}

// ID of the state to transition from
func (o StateTransitionsOutput) From() pulumi.StringOutput {
	return o.ApplyT(func(v *StateTransitions) pulumi.StringOutput { return v.From }).(pulumi.StringOutput)
}

// Transitions are a way to describe possible transformations of the current state to other states of the same type (for
// example: Initial -> Shipped). When performing a transitionState update action and transitions is set, the currently
// referenced state must have a transition to the new state. If transitions is an empty list, it means the current state is
// a final state and no further transitions are allowed. If transitions is not set, the validation is turned off. When
// performing a transitionState update action, any other state of the same type can be transitioned to
func (o StateTransitionsOutput) Tos() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *StateTransitions) pulumi.StringArrayOutput { return v.Tos }).(pulumi.StringArrayOutput)
}

type StateTransitionsArrayOutput struct{ *pulumi.OutputState }

func (StateTransitionsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StateTransitions)(nil)).Elem()
}

func (o StateTransitionsArrayOutput) ToStateTransitionsArrayOutput() StateTransitionsArrayOutput {
	return o
}

func (o StateTransitionsArrayOutput) ToStateTransitionsArrayOutputWithContext(ctx context.Context) StateTransitionsArrayOutput {
	return o
}

func (o StateTransitionsArrayOutput) Index(i pulumi.IntInput) StateTransitionsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *StateTransitions {
		return vs[0].([]*StateTransitions)[vs[1].(int)]
	}).(StateTransitionsOutput)
}

type StateTransitionsMapOutput struct{ *pulumi.OutputState }

func (StateTransitionsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StateTransitions)(nil)).Elem()
}

func (o StateTransitionsMapOutput) ToStateTransitionsMapOutput() StateTransitionsMapOutput {
	return o
}

func (o StateTransitionsMapOutput) ToStateTransitionsMapOutputWithContext(ctx context.Context) StateTransitionsMapOutput {
	return o
}

func (o StateTransitionsMapOutput) MapIndex(k pulumi.StringInput) StateTransitionsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *StateTransitions {
		return vs[0].(map[string]*StateTransitions)[vs[1].(string)]
	}).(StateTransitionsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StateTransitionsInput)(nil)).Elem(), &StateTransitions{})
	pulumi.RegisterInputType(reflect.TypeOf((*StateTransitionsArrayInput)(nil)).Elem(), StateTransitionsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StateTransitionsMapInput)(nil)).Elem(), StateTransitionsMap{})
	pulumi.RegisterOutputType(StateTransitionsOutput{})
	pulumi.RegisterOutputType(StateTransitionsArrayOutput{})
	pulumi.RegisterOutputType(StateTransitionsMapOutput{})
}
