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

__all__ = ['ProductDiscountArgs', 'ProductDiscount']

@pulumi.input_type
class ProductDiscountArgs:
    def __init__(__self__, *,
                 predicate: pulumi.Input[str],
                 sort_order: pulumi.Input[str],
                 value: pulumi.Input['ProductDiscountValueArgs'],
                 description: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 is_active: Optional[pulumi.Input[bool]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 valid_from: Optional[pulumi.Input[str]] = None,
                 valid_until: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ProductDiscount resource.
        :param pulumi.Input[str] predicate: A valid [Product Predicate](https://docs.commercetools.com/api/projects/predicates#product-predicates)
        :param pulumi.Input[str] sort_order: The string must contain a number between 0 and 1. All matching product discounts are applied to a product in the order
               defined by this field. A discount with greater sort order is prioritized higher than a discount with lower sort order.
               The sort order is unambiguous among all product discounts
        :param pulumi.Input['ProductDiscountValueArgs'] value: Defines the effect the discount will have.
               [ProductDiscountValue](https://docs.commercetools.com/api/projects/productDiscounts#productdiscountvalue)
        :param pulumi.Input[Mapping[str, Any]] description: [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
        :param pulumi.Input[bool] is_active: When set the product discount is applied to products matching the predicate
        :param pulumi.Input[str] key: User-defined unique identifier for the ProductDiscount. Must be unique across a project
        :param pulumi.Input[Mapping[str, Any]] name: [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
        """
        pulumi.set(__self__, "predicate", predicate)
        pulumi.set(__self__, "sort_order", sort_order)
        pulumi.set(__self__, "value", value)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if is_active is not None:
            pulumi.set(__self__, "is_active", is_active)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if valid_from is not None:
            pulumi.set(__self__, "valid_from", valid_from)
        if valid_until is not None:
            pulumi.set(__self__, "valid_until", valid_until)

    @property
    @pulumi.getter
    def predicate(self) -> pulumi.Input[str]:
        """
        A valid [Product Predicate](https://docs.commercetools.com/api/projects/predicates#product-predicates)
        """
        return pulumi.get(self, "predicate")

    @predicate.setter
    def predicate(self, value: pulumi.Input[str]):
        pulumi.set(self, "predicate", value)

    @property
    @pulumi.getter(name="sortOrder")
    def sort_order(self) -> pulumi.Input[str]:
        """
        The string must contain a number between 0 and 1. All matching product discounts are applied to a product in the order
        defined by this field. A discount with greater sort order is prioritized higher than a discount with lower sort order.
        The sort order is unambiguous among all product discounts
        """
        return pulumi.get(self, "sort_order")

    @sort_order.setter
    def sort_order(self, value: pulumi.Input[str]):
        pulumi.set(self, "sort_order", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input['ProductDiscountValueArgs']:
        """
        Defines the effect the discount will have.
        [ProductDiscountValue](https://docs.commercetools.com/api/projects/productDiscounts#productdiscountvalue)
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input['ProductDiscountValueArgs']):
        pulumi.set(self, "value", value)

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
    @pulumi.getter(name="isActive")
    def is_active(self) -> Optional[pulumi.Input[bool]]:
        """
        When set the product discount is applied to products matching the predicate
        """
        return pulumi.get(self, "is_active")

    @is_active.setter
    def is_active(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_active", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        """
        User-defined unique identifier for the ProductDiscount. Must be unique across a project
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
    @pulumi.getter(name="validFrom")
    def valid_from(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "valid_from")

    @valid_from.setter
    def valid_from(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "valid_from", value)

    @property
    @pulumi.getter(name="validUntil")
    def valid_until(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "valid_until")

    @valid_until.setter
    def valid_until(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "valid_until", value)


@pulumi.input_type
class _ProductDiscountState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 is_active: Optional[pulumi.Input[bool]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 predicate: Optional[pulumi.Input[str]] = None,
                 sort_order: Optional[pulumi.Input[str]] = None,
                 valid_from: Optional[pulumi.Input[str]] = None,
                 valid_until: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input['ProductDiscountValueArgs']] = None,
                 version: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering ProductDiscount resources.
        :param pulumi.Input[Mapping[str, Any]] description: [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
        :param pulumi.Input[bool] is_active: When set the product discount is applied to products matching the predicate
        :param pulumi.Input[str] key: User-defined unique identifier for the ProductDiscount. Must be unique across a project
        :param pulumi.Input[Mapping[str, Any]] name: [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
        :param pulumi.Input[str] predicate: A valid [Product Predicate](https://docs.commercetools.com/api/projects/predicates#product-predicates)
        :param pulumi.Input[str] sort_order: The string must contain a number between 0 and 1. All matching product discounts are applied to a product in the order
               defined by this field. A discount with greater sort order is prioritized higher than a discount with lower sort order.
               The sort order is unambiguous among all product discounts
        :param pulumi.Input['ProductDiscountValueArgs'] value: Defines the effect the discount will have.
               [ProductDiscountValue](https://docs.commercetools.com/api/projects/productDiscounts#productdiscountvalue)
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if is_active is not None:
            pulumi.set(__self__, "is_active", is_active)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if predicate is not None:
            pulumi.set(__self__, "predicate", predicate)
        if sort_order is not None:
            pulumi.set(__self__, "sort_order", sort_order)
        if valid_from is not None:
            pulumi.set(__self__, "valid_from", valid_from)
        if valid_until is not None:
            pulumi.set(__self__, "valid_until", valid_until)
        if value is not None:
            pulumi.set(__self__, "value", value)
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
    @pulumi.getter(name="isActive")
    def is_active(self) -> Optional[pulumi.Input[bool]]:
        """
        When set the product discount is applied to products matching the predicate
        """
        return pulumi.get(self, "is_active")

    @is_active.setter
    def is_active(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_active", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        """
        User-defined unique identifier for the ProductDiscount. Must be unique across a project
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
    def predicate(self) -> Optional[pulumi.Input[str]]:
        """
        A valid [Product Predicate](https://docs.commercetools.com/api/projects/predicates#product-predicates)
        """
        return pulumi.get(self, "predicate")

    @predicate.setter
    def predicate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "predicate", value)

    @property
    @pulumi.getter(name="sortOrder")
    def sort_order(self) -> Optional[pulumi.Input[str]]:
        """
        The string must contain a number between 0 and 1. All matching product discounts are applied to a product in the order
        defined by this field. A discount with greater sort order is prioritized higher than a discount with lower sort order.
        The sort order is unambiguous among all product discounts
        """
        return pulumi.get(self, "sort_order")

    @sort_order.setter
    def sort_order(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sort_order", value)

    @property
    @pulumi.getter(name="validFrom")
    def valid_from(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "valid_from")

    @valid_from.setter
    def valid_from(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "valid_from", value)

    @property
    @pulumi.getter(name="validUntil")
    def valid_until(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "valid_until")

    @valid_until.setter
    def valid_until(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "valid_until", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input['ProductDiscountValueArgs']]:
        """
        Defines the effect the discount will have.
        [ProductDiscountValue](https://docs.commercetools.com/api/projects/productDiscounts#productdiscountvalue)
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input['ProductDiscountValueArgs']]):
        pulumi.set(self, "value", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "version", value)


class ProductDiscount(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 is_active: Optional[pulumi.Input[bool]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 predicate: Optional[pulumi.Input[str]] = None,
                 sort_order: Optional[pulumi.Input[str]] = None,
                 valid_from: Optional[pulumi.Input[str]] = None,
                 valid_until: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[pulumi.InputType['ProductDiscountValueArgs']]] = None,
                 __props__=None):
        """
        Create a ProductDiscount resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, Any]] description: [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
        :param pulumi.Input[bool] is_active: When set the product discount is applied to products matching the predicate
        :param pulumi.Input[str] key: User-defined unique identifier for the ProductDiscount. Must be unique across a project
        :param pulumi.Input[Mapping[str, Any]] name: [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
        :param pulumi.Input[str] predicate: A valid [Product Predicate](https://docs.commercetools.com/api/projects/predicates#product-predicates)
        :param pulumi.Input[str] sort_order: The string must contain a number between 0 and 1. All matching product discounts are applied to a product in the order
               defined by this field. A discount with greater sort order is prioritized higher than a discount with lower sort order.
               The sort order is unambiguous among all product discounts
        :param pulumi.Input[pulumi.InputType['ProductDiscountValueArgs']] value: Defines the effect the discount will have.
               [ProductDiscountValue](https://docs.commercetools.com/api/projects/productDiscounts#productdiscountvalue)
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProductDiscountArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a ProductDiscount resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ProductDiscountArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProductDiscountArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 is_active: Optional[pulumi.Input[bool]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 predicate: Optional[pulumi.Input[str]] = None,
                 sort_order: Optional[pulumi.Input[str]] = None,
                 valid_from: Optional[pulumi.Input[str]] = None,
                 valid_until: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[pulumi.InputType['ProductDiscountValueArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProductDiscountArgs.__new__(ProductDiscountArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["is_active"] = is_active
            __props__.__dict__["key"] = key
            __props__.__dict__["name"] = name
            if predicate is None and not opts.urn:
                raise TypeError("Missing required property 'predicate'")
            __props__.__dict__["predicate"] = predicate
            if sort_order is None and not opts.urn:
                raise TypeError("Missing required property 'sort_order'")
            __props__.__dict__["sort_order"] = sort_order
            __props__.__dict__["valid_from"] = valid_from
            __props__.__dict__["valid_until"] = valid_until
            if value is None and not opts.urn:
                raise TypeError("Missing required property 'value'")
            __props__.__dict__["value"] = value
            __props__.__dict__["version"] = None
        super(ProductDiscount, __self__).__init__(
            'commercetools:index/productDiscount:ProductDiscount',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            is_active: Optional[pulumi.Input[bool]] = None,
            key: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            predicate: Optional[pulumi.Input[str]] = None,
            sort_order: Optional[pulumi.Input[str]] = None,
            valid_from: Optional[pulumi.Input[str]] = None,
            valid_until: Optional[pulumi.Input[str]] = None,
            value: Optional[pulumi.Input[pulumi.InputType['ProductDiscountValueArgs']]] = None,
            version: Optional[pulumi.Input[int]] = None) -> 'ProductDiscount':
        """
        Get an existing ProductDiscount resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, Any]] description: [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
        :param pulumi.Input[bool] is_active: When set the product discount is applied to products matching the predicate
        :param pulumi.Input[str] key: User-defined unique identifier for the ProductDiscount. Must be unique across a project
        :param pulumi.Input[Mapping[str, Any]] name: [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
        :param pulumi.Input[str] predicate: A valid [Product Predicate](https://docs.commercetools.com/api/projects/predicates#product-predicates)
        :param pulumi.Input[str] sort_order: The string must contain a number between 0 and 1. All matching product discounts are applied to a product in the order
               defined by this field. A discount with greater sort order is prioritized higher than a discount with lower sort order.
               The sort order is unambiguous among all product discounts
        :param pulumi.Input[pulumi.InputType['ProductDiscountValueArgs']] value: Defines the effect the discount will have.
               [ProductDiscountValue](https://docs.commercetools.com/api/projects/productDiscounts#productdiscountvalue)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ProductDiscountState.__new__(_ProductDiscountState)

        __props__.__dict__["description"] = description
        __props__.__dict__["is_active"] = is_active
        __props__.__dict__["key"] = key
        __props__.__dict__["name"] = name
        __props__.__dict__["predicate"] = predicate
        __props__.__dict__["sort_order"] = sort_order
        __props__.__dict__["valid_from"] = valid_from
        __props__.__dict__["valid_until"] = valid_until
        __props__.__dict__["value"] = value
        __props__.__dict__["version"] = version
        return ProductDiscount(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="isActive")
    def is_active(self) -> pulumi.Output[Optional[bool]]:
        """
        When set the product discount is applied to products matching the predicate
        """
        return pulumi.get(self, "is_active")

    @property
    @pulumi.getter
    def key(self) -> pulumi.Output[Optional[str]]:
        """
        User-defined unique identifier for the ProductDiscount. Must be unique across a project
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
    def predicate(self) -> pulumi.Output[str]:
        """
        A valid [Product Predicate](https://docs.commercetools.com/api/projects/predicates#product-predicates)
        """
        return pulumi.get(self, "predicate")

    @property
    @pulumi.getter(name="sortOrder")
    def sort_order(self) -> pulumi.Output[str]:
        """
        The string must contain a number between 0 and 1. All matching product discounts are applied to a product in the order
        defined by this field. A discount with greater sort order is prioritized higher than a discount with lower sort order.
        The sort order is unambiguous among all product discounts
        """
        return pulumi.get(self, "sort_order")

    @property
    @pulumi.getter(name="validFrom")
    def valid_from(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "valid_from")

    @property
    @pulumi.getter(name="validUntil")
    def valid_until(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "valid_until")

    @property
    @pulumi.getter
    def value(self) -> pulumi.Output['outputs.ProductDiscountValue']:
        """
        Defines the effect the discount will have.
        [ProductDiscountValue](https://docs.commercetools.com/api/projects/productDiscounts#productdiscountvalue)
        """
        return pulumi.get(self, "value")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[int]:
        return pulumi.get(self, "version")
