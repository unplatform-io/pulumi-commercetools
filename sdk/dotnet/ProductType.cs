// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Commercetools
{
    [CommercetoolsResourceType("commercetools:index/productType:ProductType")]
    public partial class ProductType : Pulumi.CustomResource
    {
        /// <summary>
        /// [Product attribute fefinition](https://docs.commercetools.com/api/projects/productTypes#attributedefinition)
        /// </summary>
        [Output("attributes")]
        public Output<ImmutableArray<Outputs.ProductTypeAttribute>> Attributes { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// User-specific unique identifier for the product type (max. 256 characters)
        /// </summary>
        [Output("key")]
        public Output<string?> Key { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("version")]
        public Output<int> Version { get; private set; } = null!;


        /// <summary>
        /// Create a ProductType resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProductType(string name, ProductTypeArgs? args = null, CustomResourceOptions? options = null)
            : base("commercetools:index/productType:ProductType", name, args ?? new ProductTypeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProductType(string name, Input<string> id, ProductTypeState? state = null, CustomResourceOptions? options = null)
            : base("commercetools:index/productType:ProductType", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ProductType resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProductType Get(string name, Input<string> id, ProductTypeState? state = null, CustomResourceOptions? options = null)
        {
            return new ProductType(name, id, state, options);
        }
    }

    public sealed class ProductTypeArgs : Pulumi.ResourceArgs
    {
        [Input("attributes")]
        private InputList<Inputs.ProductTypeAttributeArgs>? _attributes;

        /// <summary>
        /// [Product attribute fefinition](https://docs.commercetools.com/api/projects/productTypes#attributedefinition)
        /// </summary>
        public InputList<Inputs.ProductTypeAttributeArgs> Attributes
        {
            get => _attributes ?? (_attributes = new InputList<Inputs.ProductTypeAttributeArgs>());
            set => _attributes = value;
        }

        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// User-specific unique identifier for the product type (max. 256 characters)
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        public ProductTypeArgs()
        {
        }
    }

    public sealed class ProductTypeState : Pulumi.ResourceArgs
    {
        [Input("attributes")]
        private InputList<Inputs.ProductTypeAttributeGetArgs>? _attributes;

        /// <summary>
        /// [Product attribute fefinition](https://docs.commercetools.com/api/projects/productTypes#attributedefinition)
        /// </summary>
        public InputList<Inputs.ProductTypeAttributeGetArgs> Attributes
        {
            get => _attributes ?? (_attributes = new InputList<Inputs.ProductTypeAttributeGetArgs>());
            set => _attributes = value;
        }

        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// User-specific unique identifier for the product type (max. 256 characters)
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("version")]
        public Input<int>? Version { get; set; }

        public ProductTypeState()
        {
        }
    }
}
