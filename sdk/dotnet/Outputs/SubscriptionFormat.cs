// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Commercetools.Outputs
{

    [OutputType]
    public sealed class SubscriptionFormat
    {
        public readonly string? CloudEventsVersion;
        public readonly string? Type;

        [OutputConstructor]
        private SubscriptionFormat(
            string? cloudEventsVersion,

            string? type)
        {
            CloudEventsVersion = cloudEventsVersion;
            Type = type;
        }
    }
}
