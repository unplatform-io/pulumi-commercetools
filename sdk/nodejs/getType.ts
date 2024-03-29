// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getType(args: GetTypeArgs, opts?: pulumi.InvokeOptions): Promise<GetTypeResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("commercetools:index/getType:getType", {
        "key": args.key,
    }, opts);
}

/**
 * A collection of arguments for invoking getType.
 */
export interface GetTypeArgs {
    key: string;
}

/**
 * A collection of values returned by getType.
 */
export interface GetTypeResult {
    readonly id: string;
    readonly key: string;
}
export function getTypeOutput(args: GetTypeOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTypeResult> {
    return pulumi.output(args).apply((a: any) => getType(a, opts))
}

/**
 * A collection of arguments for invoking getType.
 */
export interface GetTypeOutputArgs {
    key: pulumi.Input<string>;
}
