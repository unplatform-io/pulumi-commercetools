// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Commercetools
{
    public partial class Subscription : Pulumi.CustomResource
    {
        [Output("changes")]
        public Output<ImmutableArray<Outputs.SubscriptionChange>> Changes { get; private set; } = null!;

        [Output("destination")]
        public Output<Outputs.SubscriptionDestination?> Destination { get; private set; } = null!;

        [Output("format")]
        public Output<Outputs.SubscriptionFormat?> Format { get; private set; } = null!;

        [Output("key")]
        public Output<string?> Key { get; private set; } = null!;

        [Output("messages")]
        public Output<ImmutableArray<Outputs.SubscriptionMessage>> Messages { get; private set; } = null!;

        [Output("version")]
        public Output<int> Version { get; private set; } = null!;


        /// <summary>
        /// Create a Subscription resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Subscription(string name, SubscriptionArgs? args = null, CustomResourceOptions? options = null)
            : base("commercetools:index/subscription:Subscription", name, args ?? new SubscriptionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Subscription(string name, Input<string> id, SubscriptionState? state = null, CustomResourceOptions? options = null)
            : base("commercetools:index/subscription:Subscription", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Subscription resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Subscription Get(string name, Input<string> id, SubscriptionState? state = null, CustomResourceOptions? options = null)
        {
            return new Subscription(name, id, state, options);
        }
    }

    public sealed class SubscriptionArgs : Pulumi.ResourceArgs
    {
        [Input("changes")]
        private InputList<Inputs.SubscriptionChangeArgs>? _changes;
        public InputList<Inputs.SubscriptionChangeArgs> Changes
        {
            get => _changes ?? (_changes = new InputList<Inputs.SubscriptionChangeArgs>());
            set => _changes = value;
        }

        [Input("destination")]
        public Input<Inputs.SubscriptionDestinationArgs>? Destination { get; set; }

        [Input("format")]
        public Input<Inputs.SubscriptionFormatArgs>? Format { get; set; }

        [Input("key")]
        public Input<string>? Key { get; set; }

        [Input("messages")]
        private InputList<Inputs.SubscriptionMessageArgs>? _messages;
        public InputList<Inputs.SubscriptionMessageArgs> Messages
        {
            get => _messages ?? (_messages = new InputList<Inputs.SubscriptionMessageArgs>());
            set => _messages = value;
        }

        public SubscriptionArgs()
        {
        }
    }

    public sealed class SubscriptionState : Pulumi.ResourceArgs
    {
        [Input("changes")]
        private InputList<Inputs.SubscriptionChangeGetArgs>? _changes;
        public InputList<Inputs.SubscriptionChangeGetArgs> Changes
        {
            get => _changes ?? (_changes = new InputList<Inputs.SubscriptionChangeGetArgs>());
            set => _changes = value;
        }

        [Input("destination")]
        public Input<Inputs.SubscriptionDestinationGetArgs>? Destination { get; set; }

        [Input("format")]
        public Input<Inputs.SubscriptionFormatGetArgs>? Format { get; set; }

        [Input("key")]
        public Input<string>? Key { get; set; }

        [Input("messages")]
        private InputList<Inputs.SubscriptionMessageGetArgs>? _messages;
        public InputList<Inputs.SubscriptionMessageGetArgs> Messages
        {
            get => _messages ?? (_messages = new InputList<Inputs.SubscriptionMessageGetArgs>());
            set => _messages = value;
        }

        [Input("version")]
        public Input<int>? Version { get; set; }

        public SubscriptionState()
        {
        }
    }
}