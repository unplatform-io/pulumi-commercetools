// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Commercetools.Inputs
{

    public sealed class ProductTypeAttributeTypeElementType2LocalizedValueGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        [Input("label", required: true)]
        private InputMap<object>? _label;
        public InputMap<object> Label
        {
            get => _label ?? (_label = new InputMap<object>());
            set => _label = value;
        }

        public ProductTypeAttributeTypeElementType2LocalizedValueGetArgs()
        {
        }
        public static new ProductTypeAttributeTypeElementType2LocalizedValueGetArgs Empty => new ProductTypeAttributeTypeElementType2LocalizedValueGetArgs();
    }
}
