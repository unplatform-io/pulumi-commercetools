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
    public sealed class ProductTypeAttributeTypeElementType
    {
        public readonly ImmutableArray<Outputs.ProductTypeAttributeTypeElementTypeLocalizedValue> LocalizedValues;
        public readonly string Name;
        public readonly string? ReferenceTypeId;
        public readonly string? TypeReference;
        public readonly ImmutableDictionary<string, object>? Values;

        [OutputConstructor]
        private ProductTypeAttributeTypeElementType(
            ImmutableArray<Outputs.ProductTypeAttributeTypeElementTypeLocalizedValue> localizedValues,

            string name,

            string? referenceTypeId,

            string? typeReference,

            ImmutableDictionary<string, object>? values)
        {
            LocalizedValues = localizedValues;
            Name = name;
            ReferenceTypeId = referenceTypeId;
            TypeReference = typeReference;
            Values = values;
        }
    }
}