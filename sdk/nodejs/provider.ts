// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The provider type for the commercetools package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 */
export class Provider extends pulumi.ProviderResource {
    /** @internal */
    public static readonly __pulumiType = 'commercetools';

    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Provider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === "pulumi:providers:" + Provider.__pulumiType;
    }

    /**
     * The API URL of the commercetools platform. https://docs.commercetools.com/http-api
     */
    public readonly apiUrl!: pulumi.Output<string | undefined>;
    /**
     * The OAuth Client ID for a commercetools platform project. https://docs.commercetools.com/http-api-authorization
     */
    public readonly clientId!: pulumi.Output<string | undefined>;
    /**
     * The OAuth Client Secret for a commercetools platform project. https://docs.commercetools.com/http-api-authorization
     */
    public readonly clientSecret!: pulumi.Output<string | undefined>;
    /**
     * The project key of commercetools platform project. https://docs.commercetools.com/getting-started
     */
    public readonly projectKey!: pulumi.Output<string | undefined>;
    /**
     * A list as string of OAuth scopes assigned to a project key, to access resources in a commercetools platform project.
     * https://docs.commercetools.com/http-api-authorization
     */
    public readonly scopes!: pulumi.Output<string | undefined>;
    /**
     * The authentication URL of the commercetools platform. https://docs.commercetools.com/http-api-authorization
     */
    public readonly tokenUrl!: pulumi.Output<string | undefined>;

    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        {
            resourceInputs["apiUrl"] = args ? args.apiUrl : undefined;
            resourceInputs["clientId"] = args?.clientId ? pulumi.secret(args.clientId) : undefined;
            resourceInputs["clientSecret"] = args?.clientSecret ? pulumi.secret(args.clientSecret) : undefined;
            resourceInputs["projectKey"] = args?.projectKey ? pulumi.secret(args.projectKey) : undefined;
            resourceInputs["scopes"] = args ? args.scopes : undefined;
            resourceInputs["tokenUrl"] = args ? args.tokenUrl : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["clientId", "clientSecret", "projectKey"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Provider.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    /**
     * The API URL of the commercetools platform. https://docs.commercetools.com/http-api
     */
    apiUrl?: pulumi.Input<string>;
    /**
     * The OAuth Client ID for a commercetools platform project. https://docs.commercetools.com/http-api-authorization
     */
    clientId?: pulumi.Input<string>;
    /**
     * The OAuth Client Secret for a commercetools platform project. https://docs.commercetools.com/http-api-authorization
     */
    clientSecret?: pulumi.Input<string>;
    /**
     * The project key of commercetools platform project. https://docs.commercetools.com/getting-started
     */
    projectKey?: pulumi.Input<string>;
    /**
     * A list as string of OAuth scopes assigned to a project key, to access resources in a commercetools platform project.
     * https://docs.commercetools.com/http-api-authorization
     */
    scopes?: pulumi.Input<string>;
    /**
     * The authentication URL of the commercetools platform. https://docs.commercetools.com/http-api-authorization
     */
    tokenUrl?: pulumi.Input<string>;
}
