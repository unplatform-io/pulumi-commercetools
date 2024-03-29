// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class State extends pulumi.CustomResource {
    /**
     * Get an existing State resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StateState, opts?: pulumi.CustomResourceOptions): State {
        return new State(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'commercetools:index/state:State';

    /**
     * Returns true if the given object is an instance of State.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is State {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === State.__pulumiType;
    }

    /**
     * Description of the State as localized string.
     */
    public readonly description!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A state can be declared as an initial state for any state machine. When a workflow starts, this first state must be an
     * initial state
     */
    public readonly initial!: pulumi.Output<boolean>;
    /**
     * Timestamp of the last Terraform update of the order.
     */
    public readonly key!: pulumi.Output<string | undefined>;
    /**
     * Name of the State as localized string.
     */
    public readonly name!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * [State Role](https://docs.commercetools.com/api/projects/states#staterole)
     */
    public readonly roles!: pulumi.Output<string[] | undefined>;
    /**
     * [StateType](https://docs.commercetools.com/api/projects/states#statetype)
     */
    public readonly type!: pulumi.Output<string>;
    public /*out*/ readonly version!: pulumi.Output<number>;

    /**
     * Create a State resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StateArgs | StateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as StateState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["initial"] = state ? state.initial : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["roles"] = state ? state.roles : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as StateArgs | undefined;
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["initial"] = args ? args.initial : undefined;
            resourceInputs["key"] = args ? args.key : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["roles"] = args ? args.roles : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["version"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(State.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering State resources.
 */
export interface StateState {
    /**
     * Description of the State as localized string.
     */
    description?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A state can be declared as an initial state for any state machine. When a workflow starts, this first state must be an
     * initial state
     */
    initial?: pulumi.Input<boolean>;
    /**
     * Timestamp of the last Terraform update of the order.
     */
    key?: pulumi.Input<string>;
    /**
     * Name of the State as localized string.
     */
    name?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * [State Role](https://docs.commercetools.com/api/projects/states#staterole)
     */
    roles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * [StateType](https://docs.commercetools.com/api/projects/states#statetype)
     */
    type?: pulumi.Input<string>;
    version?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a State resource.
 */
export interface StateArgs {
    /**
     * Description of the State as localized string.
     */
    description?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A state can be declared as an initial state for any state machine. When a workflow starts, this first state must be an
     * initial state
     */
    initial?: pulumi.Input<boolean>;
    /**
     * Timestamp of the last Terraform update of the order.
     */
    key?: pulumi.Input<string>;
    /**
     * Name of the State as localized string.
     */
    name?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * [State Role](https://docs.commercetools.com/api/projects/states#staterole)
     */
    roles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * [StateType](https://docs.commercetools.com/api/projects/states#statetype)
     */
    type: pulumi.Input<string>;
}
