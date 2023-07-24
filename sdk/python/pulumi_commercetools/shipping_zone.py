# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['ShippingZoneArgs', 'ShippingZone']

@pulumi.input_type
class ShippingZoneArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 locations: Optional[pulumi.Input[Sequence[pulumi.Input['ShippingZoneLocationArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ShippingZone resource.
        :param pulumi.Input[str] key: User-specific unique identifier for a zone. Must be unique across a project
        :param pulumi.Input[Sequence[pulumi.Input['ShippingZoneLocationArgs']]] locations: [Location](https://docs.commercetoolstools.pi/projects/zones#location)
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if locations is not None:
            pulumi.set(__self__, "locations", locations)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        """
        User-specific unique identifier for a zone. Must be unique across a project
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def locations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ShippingZoneLocationArgs']]]]:
        """
        [Location](https://docs.commercetoolstools.pi/projects/zones#location)
        """
        return pulumi.get(self, "locations")

    @locations.setter
    def locations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ShippingZoneLocationArgs']]]]):
        pulumi.set(self, "locations", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _ShippingZoneState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 locations: Optional[pulumi.Input[Sequence[pulumi.Input['ShippingZoneLocationArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering ShippingZone resources.
        :param pulumi.Input[str] key: User-specific unique identifier for a zone. Must be unique across a project
        :param pulumi.Input[Sequence[pulumi.Input['ShippingZoneLocationArgs']]] locations: [Location](https://docs.commercetoolstools.pi/projects/zones#location)
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if locations is not None:
            pulumi.set(__self__, "locations", locations)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        """
        User-specific unique identifier for a zone. Must be unique across a project
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def locations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ShippingZoneLocationArgs']]]]:
        """
        [Location](https://docs.commercetoolstools.pi/projects/zones#location)
        """
        return pulumi.get(self, "locations")

    @locations.setter
    def locations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ShippingZoneLocationArgs']]]]):
        pulumi.set(self, "locations", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "version", value)


class ShippingZone(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 locations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ShippingZoneLocationArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a ShippingZone resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] key: User-specific unique identifier for a zone. Must be unique across a project
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ShippingZoneLocationArgs']]]] locations: [Location](https://docs.commercetoolstools.pi/projects/zones#location)
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ShippingZoneArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a ShippingZone resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ShippingZoneArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ShippingZoneArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 locations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ShippingZoneLocationArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ShippingZoneArgs.__new__(ShippingZoneArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["key"] = key
            __props__.__dict__["locations"] = locations
            __props__.__dict__["name"] = name
            __props__.__dict__["version"] = None
        super(ShippingZone, __self__).__init__(
            'commercetools:index/shippingZone:ShippingZone',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            key: Optional[pulumi.Input[str]] = None,
            locations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ShippingZoneLocationArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            version: Optional[pulumi.Input[int]] = None) -> 'ShippingZone':
        """
        Get an existing ShippingZone resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] key: User-specific unique identifier for a zone. Must be unique across a project
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ShippingZoneLocationArgs']]]] locations: [Location](https://docs.commercetoolstools.pi/projects/zones#location)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ShippingZoneState.__new__(_ShippingZoneState)

        __props__.__dict__["description"] = description
        __props__.__dict__["key"] = key
        __props__.__dict__["locations"] = locations
        __props__.__dict__["name"] = name
        __props__.__dict__["version"] = version
        return ShippingZone(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def key(self) -> pulumi.Output[Optional[str]]:
        """
        User-specific unique identifier for a zone. Must be unique across a project
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def locations(self) -> pulumi.Output[Optional[Sequence['outputs.ShippingZoneLocation']]]:
        """
        [Location](https://docs.commercetoolstools.pi/projects/zones#location)
        """
        return pulumi.get(self, "locations")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[int]:
        return pulumi.get(self, "version")

