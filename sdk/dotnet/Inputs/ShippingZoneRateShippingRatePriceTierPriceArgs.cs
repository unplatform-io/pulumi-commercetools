// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Commercetools.Inputs
{

    public sealed class ShippingZoneRateShippingRatePriceTierPriceArgs : Pulumi.ResourceArgs
    {
        [Input("centAmount", required: true)]
        public Input<int> CentAmount { get; set; } = null!;

        [Input("currencyCode", required: true)]
        public Input<string> CurrencyCode { get; set; } = null!;

        public ShippingZoneRateShippingRatePriceTierPriceArgs()
        {
        }
    }
}
