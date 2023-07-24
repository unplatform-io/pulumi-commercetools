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
    public sealed class ProductTypeAttributeType
    {
        public readonly Outputs.ProductTypeAttributeTypeElementType2? ElementType2;
        public readonly ImmutableArray<Outputs.ProductTypeAttributeTypeLocalizedValue> LocalizedValues;
        public readonly string Name;
        public readonly string? ReferenceTypeId;
        public readonly string? TypeReference;
        public readonly ImmutableArray<Outputs.ProductTypeAttributeTypeValue> Values;

        [OutputConstructor]
        private ProductTypeAttributeType(
            Outputs.ProductTypeAttributeTypeElementType2? ElementType2,

            ImmutableArray<Outputs.ProductTypeAttributeTypeLocalizedValue> localizedValues,

            string name,

            string? referenceTypeId,

            string? typeReference,

            ImmutableArray<Outputs.ProductTypeAttributeTypeValue> values)
        {
            this.ElementType2 = ElementType2;
            LocalizedValues = localizedValues;
            Name = name;
            ReferenceTypeId = referenceTypeId;
            TypeReference = typeReference;
            Values = values;
        }
    }
}
