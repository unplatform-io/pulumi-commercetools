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
    public sealed class TypeFieldTypeElementType2
    {
        public readonly ImmutableArray<Outputs.TypeFieldTypeElementType2LocalizedValue> LocalizedValues;
        public readonly string Name;
        public readonly string? ReferenceTypeId;
        public readonly ImmutableArray<Outputs.TypeFieldTypeElementType2Value> Values;

        [OutputConstructor]
        private TypeFieldTypeElementType2(
            ImmutableArray<Outputs.TypeFieldTypeElementType2LocalizedValue> localizedValues,

            string name,

            string? referenceTypeId,

            ImmutableArray<Outputs.TypeFieldTypeElementType2Value> values)
        {
            LocalizedValues = localizedValues;
            Name = name;
            ReferenceTypeId = referenceTypeId;
            Values = values;
        }
    }
}
