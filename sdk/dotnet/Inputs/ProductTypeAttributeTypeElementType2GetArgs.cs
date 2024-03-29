// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Commercetools.Inputs
{

    public sealed class ProductTypeAttributeTypeElementType2GetArgs : global::Pulumi.ResourceArgs
    {
        [Input("localizedValues")]
        private InputList<Inputs.ProductTypeAttributeTypeElementType2LocalizedValueGetArgs>? _localizedValues;

        /// <summary>
        /// Localized values for the `lenum` type.
        /// </summary>
        public InputList<Inputs.ProductTypeAttributeTypeElementType2LocalizedValueGetArgs> LocalizedValues
        {
            get => _localizedValues ?? (_localizedValues = new InputList<Inputs.ProductTypeAttributeTypeElementType2LocalizedValueGetArgs>());
            set => _localizedValues = value;
        }

        /// <summary>
        /// Name of the field type. Some types require extra fields to be set. Note that changing the type after creating is not supported. You need to delete the attribute and re-add it
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Resource type the Custom Field can reference. Required when type is `reference`
        /// </summary>
        [Input("referenceTypeId")]
        public Input<string>? ReferenceTypeId { get; set; }

        /// <summary>
        /// Reference to another product type. Required when type is `nested`.
        /// </summary>
        [Input("typeReference")]
        public Input<string>? TypeReference { get; set; }

        [Input("values")]
        private InputList<Inputs.ProductTypeAttributeTypeElementType2ValueGetArgs>? _values;

        /// <summary>
        /// Values for the `enum` type.
        /// </summary>
        public InputList<Inputs.ProductTypeAttributeTypeElementType2ValueGetArgs> Values
        {
            get => _values ?? (_values = new InputList<Inputs.ProductTypeAttributeTypeElementType2ValueGetArgs>());
            set => _values = value;
        }

        public ProductTypeAttributeTypeElementType2GetArgs()
        {
        }
        public static new ProductTypeAttributeTypeElementType2GetArgs Empty => new ProductTypeAttributeTypeElementType2GetArgs();
    }
}
