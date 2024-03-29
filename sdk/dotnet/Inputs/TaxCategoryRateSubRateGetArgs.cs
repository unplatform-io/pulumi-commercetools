// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Commercetools.Inputs
{

    public sealed class TaxCategoryRateSubRateGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Number Percentage in the range of [0..1]
        /// </summary>
        [Input("amount", required: true)]
        public Input<double> Amount { get; set; } = null!;

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public TaxCategoryRateSubRateGetArgs()
        {
        }
        public static new TaxCategoryRateSubRateGetArgs Empty => new TaxCategoryRateSubRateGetArgs();
    }
}
