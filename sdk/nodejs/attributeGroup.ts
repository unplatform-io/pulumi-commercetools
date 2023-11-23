// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class AttributeGroup extends pulumi.CustomResource {
    /**
     * Get an existing AttributeGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AttributeGroupState, opts?: pulumi.CustomResourceOptions): AttributeGroup {
        return new AttributeGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'commercetools:index/attributeGroup:AttributeGroup';

    /**
     * Returns true if the given object is an instance of AttributeGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AttributeGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AttributeGroup.__pulumiType;
    }

    /**
     * Attributes with unique values.
     */
    public readonly attributes!: pulumi.Output<outputs.AttributeGroupAttribute[] | undefined>;
    /**
     * Description of the State as localized string.
     */
    public readonly description!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * User-defined unique identifier of the AttributeGroup.
     */
    public readonly key!: pulumi.Output<string | undefined>;
    /**
     * Name of the State as localized string.
     */
    public readonly name!: pulumi.Output<{[key: string]: string}>;
    /**
     * Current version of the AttributeGroup.
     */
    public /*out*/ readonly version!: pulumi.Output<number>;

    /**
     * Create a AttributeGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AttributeGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AttributeGroupArgs | AttributeGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AttributeGroupState | undefined;
            resourceInputs["attributes"] = state ? state.attributes : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as AttributeGroupArgs | undefined;
            resourceInputs["attributes"] = args ? args.attributes : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["key"] = args ? args.key : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["version"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AttributeGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AttributeGroup resources.
 */
export interface AttributeGroupState {
    /**
     * Attributes with unique values.
     */
    attributes?: pulumi.Input<pulumi.Input<inputs.AttributeGroupAttribute>[]>;
    /**
     * Description of the State as localized string.
     */
    description?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * User-defined unique identifier of the AttributeGroup.
     */
    key?: pulumi.Input<string>;
    /**
     * Name of the State as localized string.
     */
    name?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Current version of the AttributeGroup.
     */
    version?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a AttributeGroup resource.
 */
export interface AttributeGroupArgs {
    /**
     * Attributes with unique values.
     */
    attributes?: pulumi.Input<pulumi.Input<inputs.AttributeGroupAttribute>[]>;
    /**
     * Description of the State as localized string.
     */
    description?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * User-defined unique identifier of the AttributeGroup.
     */
    key?: pulumi.Input<string>;
    /**
     * Name of the State as localized string.
     */
    name?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}