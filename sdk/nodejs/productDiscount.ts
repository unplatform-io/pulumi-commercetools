// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class ProductDiscount extends pulumi.CustomResource {
    /**
     * Get an existing ProductDiscount resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProductDiscountState, opts?: pulumi.CustomResourceOptions): ProductDiscount {
        return new ProductDiscount(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'commercetools:index/productDiscount:ProductDiscount';

    /**
     * Returns true if the given object is an instance of ProductDiscount.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProductDiscount {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProductDiscount.__pulumiType;
    }

    /**
     * [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
     */
    public readonly description!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * When set the product discount is applied to products matching the predicate
     */
    public readonly isActive!: pulumi.Output<boolean | undefined>;
    /**
     * User-defined unique identifier for the ProductDiscount. Must be unique across a project
     */
    public readonly key!: pulumi.Output<string | undefined>;
    /**
     * [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
     */
    public readonly name!: pulumi.Output<{[key: string]: any}>;
    /**
     * A valid [Product Predicate](https://docs.commercetools.com/api/projects/predicates#product-predicates)
     */
    public readonly predicate!: pulumi.Output<string>;
    /**
     * The string must contain a number between 0 and 1. All matching product discounts are applied to a product in the order
     * defined by this field. A discount with greater sort order is prioritized higher than a discount with lower sort order.
     * The sort order is unambiguous among all product discounts
     */
    public readonly sortOrder!: pulumi.Output<string>;
    public readonly validFrom!: pulumi.Output<string | undefined>;
    public readonly validUntil!: pulumi.Output<string | undefined>;
    /**
     * Defines the effect the discount will have.
     * [ProductDiscountValue](https://docs.commercetools.com/api/projects/productDiscounts#productdiscountvalue)
     */
    public readonly value!: pulumi.Output<outputs.ProductDiscountValue>;
    public /*out*/ readonly version!: pulumi.Output<number>;

    /**
     * Create a ProductDiscount resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProductDiscountArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProductDiscountArgs | ProductDiscountState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProductDiscountState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["isActive"] = state ? state.isActive : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["predicate"] = state ? state.predicate : undefined;
            resourceInputs["sortOrder"] = state ? state.sortOrder : undefined;
            resourceInputs["validFrom"] = state ? state.validFrom : undefined;
            resourceInputs["validUntil"] = state ? state.validUntil : undefined;
            resourceInputs["value"] = state ? state.value : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as ProductDiscountArgs | undefined;
            if ((!args || args.predicate === undefined) && !opts.urn) {
                throw new Error("Missing required property 'predicate'");
            }
            if ((!args || args.sortOrder === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sortOrder'");
            }
            if ((!args || args.value === undefined) && !opts.urn) {
                throw new Error("Missing required property 'value'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["isActive"] = args ? args.isActive : undefined;
            resourceInputs["key"] = args ? args.key : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["predicate"] = args ? args.predicate : undefined;
            resourceInputs["sortOrder"] = args ? args.sortOrder : undefined;
            resourceInputs["validFrom"] = args ? args.validFrom : undefined;
            resourceInputs["validUntil"] = args ? args.validUntil : undefined;
            resourceInputs["value"] = args ? args.value : undefined;
            resourceInputs["version"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ProductDiscount.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProductDiscount resources.
 */
export interface ProductDiscountState {
    /**
     * [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
     */
    description?: pulumi.Input<{[key: string]: any}>;
    /**
     * When set the product discount is applied to products matching the predicate
     */
    isActive?: pulumi.Input<boolean>;
    /**
     * User-defined unique identifier for the ProductDiscount. Must be unique across a project
     */
    key?: pulumi.Input<string>;
    /**
     * [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
     */
    name?: pulumi.Input<{[key: string]: any}>;
    /**
     * A valid [Product Predicate](https://docs.commercetools.com/api/projects/predicates#product-predicates)
     */
    predicate?: pulumi.Input<string>;
    /**
     * The string must contain a number between 0 and 1. All matching product discounts are applied to a product in the order
     * defined by this field. A discount with greater sort order is prioritized higher than a discount with lower sort order.
     * The sort order is unambiguous among all product discounts
     */
    sortOrder?: pulumi.Input<string>;
    validFrom?: pulumi.Input<string>;
    validUntil?: pulumi.Input<string>;
    /**
     * Defines the effect the discount will have.
     * [ProductDiscountValue](https://docs.commercetools.com/api/projects/productDiscounts#productdiscountvalue)
     */
    value?: pulumi.Input<inputs.ProductDiscountValue>;
    version?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a ProductDiscount resource.
 */
export interface ProductDiscountArgs {
    /**
     * [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
     */
    description?: pulumi.Input<{[key: string]: any}>;
    /**
     * When set the product discount is applied to products matching the predicate
     */
    isActive?: pulumi.Input<boolean>;
    /**
     * User-defined unique identifier for the ProductDiscount. Must be unique across a project
     */
    key?: pulumi.Input<string>;
    /**
     * [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
     */
    name?: pulumi.Input<{[key: string]: any}>;
    /**
     * A valid [Product Predicate](https://docs.commercetools.com/api/projects/predicates#product-predicates)
     */
    predicate: pulumi.Input<string>;
    /**
     * The string must contain a number between 0 and 1. All matching product discounts are applied to a product in the order
     * defined by this field. A discount with greater sort order is prioritized higher than a discount with lower sort order.
     * The sort order is unambiguous among all product discounts
     */
    sortOrder: pulumi.Input<string>;
    validFrom?: pulumi.Input<string>;
    validUntil?: pulumi.Input<string>;
    /**
     * Defines the effect the discount will have.
     * [ProductDiscountValue](https://docs.commercetools.com/api/projects/productDiscounts#productdiscountvalue)
     */
    value: pulumi.Input<inputs.ProductDiscountValue>;
}
