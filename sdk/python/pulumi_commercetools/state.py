# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['StateArgs', 'State']

@pulumi.input_type
class StateArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 type: pulumi.Input[str],
                 description: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 initial: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 transitions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a State resource.
        :param pulumi.Input[str] key: A unique identifier for the state
        :param pulumi.Input[str] type: [StateType](https://docs.commercetools.com/api/projects/states#statetype)
        :param pulumi.Input[Mapping[str, Any]] description: [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
        :param pulumi.Input[bool] initial: A state can be declared as an initial state for any state machine. When a workflow starts, this first state must be an
               initial state
        :param pulumi.Input[Mapping[str, Any]] name: [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
        :param pulumi.Input[Sequence[pulumi.Input[str]]] roles: Array of [State Role](https://docs.commercetools.com/api/projects/states#staterole)
        :param pulumi.Input[Sequence[pulumi.Input[str]]] transitions: Transitions are a way to describe possible transformations of the current state to other states of the same type (for
               example: Initial -> Shipped). When performing a transitionState update action and transitions is set, the currently
               referenced state must have a transition to the new state. If transitions is an empty list, it means the current state is
               a final state and no further transitions are allowed. If transitions is not set, the validation is turned off. When
               performing a transitionState update action, any other state of the same type can be transitioned to
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "type", type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if initial is not None:
            pulumi.set(__self__, "initial", initial)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if roles is not None:
            pulumi.set(__self__, "roles", roles)
        if transitions is not None:
            pulumi.set(__self__, "transitions", transitions)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        A unique identifier for the state
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        [StateType](https://docs.commercetools.com/api/projects/states#statetype)
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def initial(self) -> Optional[pulumi.Input[bool]]:
        """
        A state can be declared as an initial state for any state machine. When a workflow starts, this first state must be an
        initial state
        """
        return pulumi.get(self, "initial")

    @initial.setter
    def initial(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "initial", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def roles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Array of [State Role](https://docs.commercetools.com/api/projects/states#staterole)
        """
        return pulumi.get(self, "roles")

    @roles.setter
    def roles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "roles", value)

    @property
    @pulumi.getter
    def transitions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Transitions are a way to describe possible transformations of the current state to other states of the same type (for
        example: Initial -> Shipped). When performing a transitionState update action and transitions is set, the currently
        referenced state must have a transition to the new state. If transitions is an empty list, it means the current state is
        a final state and no further transitions are allowed. If transitions is not set, the validation is turned off. When
        performing a transitionState update action, any other state of the same type can be transitioned to
        """
        return pulumi.get(self, "transitions")

    @transitions.setter
    def transitions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "transitions", value)


@pulumi.input_type
class _StateState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 initial: Optional[pulumi.Input[bool]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 transitions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering State resources.
        :param pulumi.Input[Mapping[str, Any]] description: [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
        :param pulumi.Input[bool] initial: A state can be declared as an initial state for any state machine. When a workflow starts, this first state must be an
               initial state
        :param pulumi.Input[str] key: A unique identifier for the state
        :param pulumi.Input[Mapping[str, Any]] name: [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
        :param pulumi.Input[Sequence[pulumi.Input[str]]] roles: Array of [State Role](https://docs.commercetools.com/api/projects/states#staterole)
        :param pulumi.Input[Sequence[pulumi.Input[str]]] transitions: Transitions are a way to describe possible transformations of the current state to other states of the same type (for
               example: Initial -> Shipped). When performing a transitionState update action and transitions is set, the currently
               referenced state must have a transition to the new state. If transitions is an empty list, it means the current state is
               a final state and no further transitions are allowed. If transitions is not set, the validation is turned off. When
               performing a transitionState update action, any other state of the same type can be transitioned to
        :param pulumi.Input[str] type: [StateType](https://docs.commercetools.com/api/projects/states#statetype)
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if initial is not None:
            pulumi.set(__self__, "initial", initial)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if roles is not None:
            pulumi.set(__self__, "roles", roles)
        if transitions is not None:
            pulumi.set(__self__, "transitions", transitions)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def initial(self) -> Optional[pulumi.Input[bool]]:
        """
        A state can be declared as an initial state for any state machine. When a workflow starts, this first state must be an
        initial state
        """
        return pulumi.get(self, "initial")

    @initial.setter
    def initial(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "initial", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        """
        A unique identifier for the state
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def roles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Array of [State Role](https://docs.commercetools.com/api/projects/states#staterole)
        """
        return pulumi.get(self, "roles")

    @roles.setter
    def roles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "roles", value)

    @property
    @pulumi.getter
    def transitions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Transitions are a way to describe possible transformations of the current state to other states of the same type (for
        example: Initial -> Shipped). When performing a transitionState update action and transitions is set, the currently
        referenced state must have a transition to the new state. If transitions is an empty list, it means the current state is
        a final state and no further transitions are allowed. If transitions is not set, the validation is turned off. When
        performing a transitionState update action, any other state of the same type can be transitioned to
        """
        return pulumi.get(self, "transitions")

    @transitions.setter
    def transitions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "transitions", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        [StateType](https://docs.commercetools.com/api/projects/states#statetype)
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "version", value)


class State(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 initial: Optional[pulumi.Input[bool]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 transitions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a State resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, Any]] description: [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
        :param pulumi.Input[bool] initial: A state can be declared as an initial state for any state machine. When a workflow starts, this first state must be an
               initial state
        :param pulumi.Input[str] key: A unique identifier for the state
        :param pulumi.Input[Mapping[str, Any]] name: [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
        :param pulumi.Input[Sequence[pulumi.Input[str]]] roles: Array of [State Role](https://docs.commercetools.com/api/projects/states#staterole)
        :param pulumi.Input[Sequence[pulumi.Input[str]]] transitions: Transitions are a way to describe possible transformations of the current state to other states of the same type (for
               example: Initial -> Shipped). When performing a transitionState update action and transitions is set, the currently
               referenced state must have a transition to the new state. If transitions is an empty list, it means the current state is
               a final state and no further transitions are allowed. If transitions is not set, the validation is turned off. When
               performing a transitionState update action, any other state of the same type can be transitioned to
        :param pulumi.Input[str] type: [StateType](https://docs.commercetools.com/api/projects/states#statetype)
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: StateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a State resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param StateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 initial: Optional[pulumi.Input[bool]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 transitions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = StateArgs.__new__(StateArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["initial"] = initial
            if key is None and not opts.urn:
                raise TypeError("Missing required property 'key'")
            __props__.__dict__["key"] = key
            __props__.__dict__["name"] = name
            __props__.__dict__["roles"] = roles
            __props__.__dict__["transitions"] = transitions
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["version"] = None
        super(State, __self__).__init__(
            'commercetools:index/state:State',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            initial: Optional[pulumi.Input[bool]] = None,
            key: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            transitions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            type: Optional[pulumi.Input[str]] = None,
            version: Optional[pulumi.Input[int]] = None) -> 'State':
        """
        Get an existing State resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, Any]] description: [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
        :param pulumi.Input[bool] initial: A state can be declared as an initial state for any state machine. When a workflow starts, this first state must be an
               initial state
        :param pulumi.Input[str] key: A unique identifier for the state
        :param pulumi.Input[Mapping[str, Any]] name: [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
        :param pulumi.Input[Sequence[pulumi.Input[str]]] roles: Array of [State Role](https://docs.commercetools.com/api/projects/states#staterole)
        :param pulumi.Input[Sequence[pulumi.Input[str]]] transitions: Transitions are a way to describe possible transformations of the current state to other states of the same type (for
               example: Initial -> Shipped). When performing a transitionState update action and transitions is set, the currently
               referenced state must have a transition to the new state. If transitions is an empty list, it means the current state is
               a final state and no further transitions are allowed. If transitions is not set, the validation is turned off. When
               performing a transitionState update action, any other state of the same type can be transitioned to
        :param pulumi.Input[str] type: [StateType](https://docs.commercetools.com/api/projects/states#statetype)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _StateState.__new__(_StateState)

        __props__.__dict__["description"] = description
        __props__.__dict__["initial"] = initial
        __props__.__dict__["key"] = key
        __props__.__dict__["name"] = name
        __props__.__dict__["roles"] = roles
        __props__.__dict__["transitions"] = transitions
        __props__.__dict__["type"] = type
        __props__.__dict__["version"] = version
        return State(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def initial(self) -> pulumi.Output[Optional[bool]]:
        """
        A state can be declared as an initial state for any state machine. When a workflow starts, this first state must be an
        initial state
        """
        return pulumi.get(self, "initial")

    @property
    @pulumi.getter
    def key(self) -> pulumi.Output[str]:
        """
        A unique identifier for the state
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Mapping[str, Any]]:
        """
        [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def roles(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Array of [State Role](https://docs.commercetools.com/api/projects/states#staterole)
        """
        return pulumi.get(self, "roles")

    @property
    @pulumi.getter
    def transitions(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Transitions are a way to describe possible transformations of the current state to other states of the same type (for
        example: Initial -> Shipped). When performing a transitionState update action and transitions is set, the currently
        referenced state must have a transition to the new state. If transitions is an empty list, it means the current state is
        a final state and no further transitions are allowed. If transitions is not set, the validation is turned off. When
        performing a transitionState update action, any other state of the same type can be transitioned to
        """
        return pulumi.get(self, "transitions")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        [StateType](https://docs.commercetools.com/api/projects/states#statetype)
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[int]:
        return pulumi.get(self, "version")

