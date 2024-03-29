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
    public sealed class ProjectSettingsExternalOauth
    {
        /// <summary>
        /// Partially hidden on retrieval
        /// </summary>
        public readonly string? AuthorizationHeader;
        public readonly string? Url;

        [OutputConstructor]
        private ProjectSettingsExternalOauth(
            string? authorizationHeader,

            string? url)
        {
            AuthorizationHeader = authorizationHeader;
            Url = url;
        }
    }
}
