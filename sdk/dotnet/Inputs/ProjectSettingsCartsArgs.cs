// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Commercetools.Inputs
{

    public sealed class ProjectSettingsCartsArgs : Pulumi.ResourceArgs
    {
        [Input("countryTaxRateFallbackEnabled", required: true)]
        public Input<bool> CountryTaxRateFallbackEnabled { get; set; } = null!;

        [Input("deleteDaysAfterLastModification")]
        public Input<int>? DeleteDaysAfterLastModification { get; set; }

        public ProjectSettingsCartsArgs()
        {
        }
    }
}
