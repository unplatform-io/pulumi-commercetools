// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Commercetools
{
    [CommercetoolsResourceType("commercetools:index/category:Category")]
    public partial class Category : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Can be used to store images, icons or movies related to this category
        /// </summary>
        [Output("assets")]
        public Output<ImmutableArray<Outputs.CategoryAsset>> Assets { get; private set; } = null!;

        [Output("custom")]
        public Output<Outputs.CategoryCustom?> Custom { get; private set; } = null!;

        [Output("description")]
        public Output<ImmutableDictionary<string, object>?> Description { get; private set; } = null!;

        [Output("externalId")]
        public Output<string?> ExternalId { get; private set; } = null!;

        /// <summary>
        /// Category-specific unique identifier. Must be unique across a project
        /// </summary>
        [Output("key")]
        public Output<string?> Key { get; private set; } = null!;

        [Output("metaDescription")]
        public Output<ImmutableDictionary<string, object>?> MetaDescription { get; private set; } = null!;

        [Output("metaKeywords")]
        public Output<ImmutableDictionary<string, object>?> MetaKeywords { get; private set; } = null!;

        [Output("metaTitle")]
        public Output<ImmutableDictionary<string, object>?> MetaTitle { get; private set; } = null!;

        [Output("name")]
        public Output<ImmutableDictionary<string, object>> Name { get; private set; } = null!;

        /// <summary>
        /// An attribute as base for a custom category order in one level, filled with random value when left empty
        /// </summary>
        [Output("orderHint")]
        public Output<string?> OrderHint { get; private set; } = null!;

        /// <summary>
        /// A category that is the parent of this category in the category tree
        /// </summary>
        [Output("parent")]
        public Output<string?> Parent { get; private set; } = null!;

        /// <summary>
        /// Human readable identifiers, needs to be unique
        /// </summary>
        [Output("slug")]
        public Output<ImmutableDictionary<string, object>> Slug { get; private set; } = null!;

        [Output("version")]
        public Output<int> Version { get; private set; } = null!;


        /// <summary>
        /// Create a Category resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Category(string name, CategoryArgs args, CustomResourceOptions? options = null)
            : base("commercetools:index/category:Category", name, args ?? new CategoryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Category(string name, Input<string> id, CategoryState? state = null, CustomResourceOptions? options = null)
            : base("commercetools:index/category:Category", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Category resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Category Get(string name, Input<string> id, CategoryState? state = null, CustomResourceOptions? options = null)
        {
            return new Category(name, id, state, options);
        }
    }

    public sealed class CategoryArgs : global::Pulumi.ResourceArgs
    {
        [Input("assets")]
        private InputList<Inputs.CategoryAssetArgs>? _assets;

        /// <summary>
        /// Can be used to store images, icons or movies related to this category
        /// </summary>
        public InputList<Inputs.CategoryAssetArgs> Assets
        {
            get => _assets ?? (_assets = new InputList<Inputs.CategoryAssetArgs>());
            set => _assets = value;
        }

        [Input("custom")]
        public Input<Inputs.CategoryCustomArgs>? Custom { get; set; }

        [Input("description")]
        private InputMap<object>? _description;
        public InputMap<object> Description
        {
            get => _description ?? (_description = new InputMap<object>());
            set => _description = value;
        }

        [Input("externalId")]
        public Input<string>? ExternalId { get; set; }

        /// <summary>
        /// Category-specific unique identifier. Must be unique across a project
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        [Input("metaDescription")]
        private InputMap<object>? _metaDescription;
        public InputMap<object> MetaDescription
        {
            get => _metaDescription ?? (_metaDescription = new InputMap<object>());
            set => _metaDescription = value;
        }

        [Input("metaKeywords")]
        private InputMap<object>? _metaKeywords;
        public InputMap<object> MetaKeywords
        {
            get => _metaKeywords ?? (_metaKeywords = new InputMap<object>());
            set => _metaKeywords = value;
        }

        [Input("metaTitle")]
        private InputMap<object>? _metaTitle;
        public InputMap<object> MetaTitle
        {
            get => _metaTitle ?? (_metaTitle = new InputMap<object>());
            set => _metaTitle = value;
        }

        [Input("name")]
        private InputMap<object>? _name;
        public InputMap<object> Name
        {
            get => _name ?? (_name = new InputMap<object>());
            set => _name = value;
        }

        /// <summary>
        /// An attribute as base for a custom category order in one level, filled with random value when left empty
        /// </summary>
        [Input("orderHint")]
        public Input<string>? OrderHint { get; set; }

        /// <summary>
        /// A category that is the parent of this category in the category tree
        /// </summary>
        [Input("parent")]
        public Input<string>? Parent { get; set; }

        [Input("slug", required: true)]
        private InputMap<object>? _slug;

        /// <summary>
        /// Human readable identifiers, needs to be unique
        /// </summary>
        public InputMap<object> Slug
        {
            get => _slug ?? (_slug = new InputMap<object>());
            set => _slug = value;
        }

        public CategoryArgs()
        {
        }
        public static new CategoryArgs Empty => new CategoryArgs();
    }

    public sealed class CategoryState : global::Pulumi.ResourceArgs
    {
        [Input("assets")]
        private InputList<Inputs.CategoryAssetGetArgs>? _assets;

        /// <summary>
        /// Can be used to store images, icons or movies related to this category
        /// </summary>
        public InputList<Inputs.CategoryAssetGetArgs> Assets
        {
            get => _assets ?? (_assets = new InputList<Inputs.CategoryAssetGetArgs>());
            set => _assets = value;
        }

        [Input("custom")]
        public Input<Inputs.CategoryCustomGetArgs>? Custom { get; set; }

        [Input("description")]
        private InputMap<object>? _description;
        public InputMap<object> Description
        {
            get => _description ?? (_description = new InputMap<object>());
            set => _description = value;
        }

        [Input("externalId")]
        public Input<string>? ExternalId { get; set; }

        /// <summary>
        /// Category-specific unique identifier. Must be unique across a project
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        [Input("metaDescription")]
        private InputMap<object>? _metaDescription;
        public InputMap<object> MetaDescription
        {
            get => _metaDescription ?? (_metaDescription = new InputMap<object>());
            set => _metaDescription = value;
        }

        [Input("metaKeywords")]
        private InputMap<object>? _metaKeywords;
        public InputMap<object> MetaKeywords
        {
            get => _metaKeywords ?? (_metaKeywords = new InputMap<object>());
            set => _metaKeywords = value;
        }

        [Input("metaTitle")]
        private InputMap<object>? _metaTitle;
        public InputMap<object> MetaTitle
        {
            get => _metaTitle ?? (_metaTitle = new InputMap<object>());
            set => _metaTitle = value;
        }

        [Input("name")]
        private InputMap<object>? _name;
        public InputMap<object> Name
        {
            get => _name ?? (_name = new InputMap<object>());
            set => _name = value;
        }

        /// <summary>
        /// An attribute as base for a custom category order in one level, filled with random value when left empty
        /// </summary>
        [Input("orderHint")]
        public Input<string>? OrderHint { get; set; }

        /// <summary>
        /// A category that is the parent of this category in the category tree
        /// </summary>
        [Input("parent")]
        public Input<string>? Parent { get; set; }

        [Input("slug")]
        private InputMap<object>? _slug;

        /// <summary>
        /// Human readable identifiers, needs to be unique
        /// </summary>
        public InputMap<object> Slug
        {
            get => _slug ?? (_slug = new InputMap<object>());
            set => _slug = value;
        }

        [Input("version")]
        public Input<int>? Version { get; set; }

        public CategoryState()
        {
        }
        public static new CategoryState Empty => new CategoryState();
    }
}