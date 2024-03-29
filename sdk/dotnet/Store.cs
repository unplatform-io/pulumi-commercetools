// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Commercetools
{
    [CommercetoolsResourceType("commercetools:index/store:Store")]
    public partial class Store : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)
        /// </summary>
        [Output("countries")]
        public Output<ImmutableArray<string>> Countries { get; private set; } = null!;

        [Output("custom")]
        public Output<Outputs.StoreCustom?> Custom { get; private set; } = null!;

        /// <summary>
        /// Set of ResourceIdentifier to a Channel with ProductDistribution
        /// </summary>
        [Output("distributionChannels")]
        public Output<ImmutableArray<string>> DistributionChannels { get; private set; } = null!;

        /// <summary>
        /// User-specific unique identifier for the store. The key is mandatory and immutable. It is used to reference the store
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        /// <summary>
        /// [IETF Language Tag](https://en.wikipedia.org/wiki/IETF_language_tag)
        /// </summary>
        [Output("languages")]
        public Output<ImmutableArray<string>> Languages { get; private set; } = null!;

        /// <summary>
        /// [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
        /// </summary>
        [Output("name")]
        public Output<ImmutableDictionary<string, object>?> Name { get; private set; } = null!;

        /// <summary>
        /// Controls availability of Products for this Store via Product Selections
        /// </summary>
        [Output("productSelections")]
        public Output<ImmutableArray<Outputs.StoreProductSelection>> ProductSelections { get; private set; } = null!;

        /// <summary>
        /// Set of ResourceIdentifier of Channels with InventorySupply
        /// </summary>
        [Output("supplyChannels")]
        public Output<ImmutableArray<string>> SupplyChannels { get; private set; } = null!;

        [Output("version")]
        public Output<int> Version { get; private set; } = null!;


        /// <summary>
        /// Create a Store resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Store(string name, StoreArgs args, CustomResourceOptions? options = null)
            : base("commercetools:index/store:Store", name, args ?? new StoreArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Store(string name, Input<string> id, StoreState? state = null, CustomResourceOptions? options = null)
            : base("commercetools:index/store:Store", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Store resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Store Get(string name, Input<string> id, StoreState? state = null, CustomResourceOptions? options = null)
        {
            return new Store(name, id, state, options);
        }
    }

    public sealed class StoreArgs : global::Pulumi.ResourceArgs
    {
        [Input("countries")]
        private InputList<string>? _countries;

        /// <summary>
        /// A two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)
        /// </summary>
        public InputList<string> Countries
        {
            get => _countries ?? (_countries = new InputList<string>());
            set => _countries = value;
        }

        [Input("custom")]
        public Input<Inputs.StoreCustomArgs>? Custom { get; set; }

        [Input("distributionChannels")]
        private InputList<string>? _distributionChannels;

        /// <summary>
        /// Set of ResourceIdentifier to a Channel with ProductDistribution
        /// </summary>
        public InputList<string> DistributionChannels
        {
            get => _distributionChannels ?? (_distributionChannels = new InputList<string>());
            set => _distributionChannels = value;
        }

        /// <summary>
        /// User-specific unique identifier for the store. The key is mandatory and immutable. It is used to reference the store
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        [Input("languages")]
        private InputList<string>? _languages;

        /// <summary>
        /// [IETF Language Tag](https://en.wikipedia.org/wiki/IETF_language_tag)
        /// </summary>
        public InputList<string> Languages
        {
            get => _languages ?? (_languages = new InputList<string>());
            set => _languages = value;
        }

        [Input("name")]
        private InputMap<object>? _name;

        /// <summary>
        /// [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
        /// </summary>
        public InputMap<object> Name
        {
            get => _name ?? (_name = new InputMap<object>());
            set => _name = value;
        }

        [Input("productSelections")]
        private InputList<Inputs.StoreProductSelectionArgs>? _productSelections;

        /// <summary>
        /// Controls availability of Products for this Store via Product Selections
        /// </summary>
        public InputList<Inputs.StoreProductSelectionArgs> ProductSelections
        {
            get => _productSelections ?? (_productSelections = new InputList<Inputs.StoreProductSelectionArgs>());
            set => _productSelections = value;
        }

        [Input("supplyChannels")]
        private InputList<string>? _supplyChannels;

        /// <summary>
        /// Set of ResourceIdentifier of Channels with InventorySupply
        /// </summary>
        public InputList<string> SupplyChannels
        {
            get => _supplyChannels ?? (_supplyChannels = new InputList<string>());
            set => _supplyChannels = value;
        }

        public StoreArgs()
        {
        }
        public static new StoreArgs Empty => new StoreArgs();
    }

    public sealed class StoreState : global::Pulumi.ResourceArgs
    {
        [Input("countries")]
        private InputList<string>? _countries;

        /// <summary>
        /// A two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)
        /// </summary>
        public InputList<string> Countries
        {
            get => _countries ?? (_countries = new InputList<string>());
            set => _countries = value;
        }

        [Input("custom")]
        public Input<Inputs.StoreCustomGetArgs>? Custom { get; set; }

        [Input("distributionChannels")]
        private InputList<string>? _distributionChannels;

        /// <summary>
        /// Set of ResourceIdentifier to a Channel with ProductDistribution
        /// </summary>
        public InputList<string> DistributionChannels
        {
            get => _distributionChannels ?? (_distributionChannels = new InputList<string>());
            set => _distributionChannels = value;
        }

        /// <summary>
        /// User-specific unique identifier for the store. The key is mandatory and immutable. It is used to reference the store
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        [Input("languages")]
        private InputList<string>? _languages;

        /// <summary>
        /// [IETF Language Tag](https://en.wikipedia.org/wiki/IETF_language_tag)
        /// </summary>
        public InputList<string> Languages
        {
            get => _languages ?? (_languages = new InputList<string>());
            set => _languages = value;
        }

        [Input("name")]
        private InputMap<object>? _name;

        /// <summary>
        /// [LocalizedString](https://docs.commercetools.com/api/types#localizedstring)
        /// </summary>
        public InputMap<object> Name
        {
            get => _name ?? (_name = new InputMap<object>());
            set => _name = value;
        }

        [Input("productSelections")]
        private InputList<Inputs.StoreProductSelectionGetArgs>? _productSelections;

        /// <summary>
        /// Controls availability of Products for this Store via Product Selections
        /// </summary>
        public InputList<Inputs.StoreProductSelectionGetArgs> ProductSelections
        {
            get => _productSelections ?? (_productSelections = new InputList<Inputs.StoreProductSelectionGetArgs>());
            set => _productSelections = value;
        }

        [Input("supplyChannels")]
        private InputList<string>? _supplyChannels;

        /// <summary>
        /// Set of ResourceIdentifier of Channels with InventorySupply
        /// </summary>
        public InputList<string> SupplyChannels
        {
            get => _supplyChannels ?? (_supplyChannels = new InputList<string>());
            set => _supplyChannels = value;
        }

        [Input("version")]
        public Input<int>? Version { get; set; }

        public StoreState()
        {
        }
        public static new StoreState Empty => new StoreState();
    }
}
