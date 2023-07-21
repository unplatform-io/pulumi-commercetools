// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Commercetools.Inputs
{

    public sealed class ProductDiscountValueArgs : global::Pulumi.ResourceArgs
    {
        [Input("monies")]
        private InputList<Inputs.ProductDiscountValueMoneyArgs>? _monies;
        public InputList<Inputs.ProductDiscountValueMoneyArgs> Monies
        {
            get => _monies ?? (_monies = new InputList<Inputs.ProductDiscountValueMoneyArgs>());
            set => _monies = value;
        }

        [Input("permyriad")]
        public Input<int>? Permyriad { get; set; }

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ProductDiscountValueArgs()
        {
        }
        public static new ProductDiscountValueArgs Empty => new ProductDiscountValueArgs();
    }
}