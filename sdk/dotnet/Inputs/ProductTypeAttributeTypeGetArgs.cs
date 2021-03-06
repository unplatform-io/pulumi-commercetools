// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Commercetools.Inputs
{

    public sealed class ProductTypeAttributeTypeGetArgs : Pulumi.ResourceArgs
    {
        [Input("ElementType2")]
        public Input<Inputs.ProductTypeAttributeTypeElementType2GetArgs>? ElementType2 { get; set; }

        [Input("localizedValues")]
        private InputList<Inputs.ProductTypeAttributeTypeLocalizedValueGetArgs>? _localizedValues;
        public InputList<Inputs.ProductTypeAttributeTypeLocalizedValueGetArgs> LocalizedValues
        {
            get => _localizedValues ?? (_localizedValues = new InputList<Inputs.ProductTypeAttributeTypeLocalizedValueGetArgs>());
            set => _localizedValues = value;
        }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("referenceTypeId")]
        public Input<string>? ReferenceTypeId { get; set; }

        [Input("typeReference")]
        public Input<string>? TypeReference { get; set; }

        [Input("values")]
        private InputMap<object>? _values;
        public InputMap<object> Values
        {
            get => _values ?? (_values = new InputMap<object>());
            set => _values = value;
        }

        public ProductTypeAttributeTypeGetArgs()
        {
        }
    }
}
