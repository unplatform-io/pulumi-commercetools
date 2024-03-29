// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Commercetools.Inputs
{

    public sealed class ApiExtensionTriggerArgs : global::Pulumi.ResourceArgs
    {
        [Input("actions", required: true)]
        private InputList<string>? _actions;

        /// <summary>
        /// Currently, Create and Update are supported
        /// </summary>
        public InputList<string> Actions
        {
            get => _actions ?? (_actions = new InputList<string>());
            set => _actions = value;
        }

        /// <summary>
        /// Valid predicate that controls the conditions under which the API Extension is called.
        /// </summary>
        [Input("condition")]
        public Input<string>? Condition { get; set; }

        /// <summary>
        /// Currently, cart, order, payment, and customer are supported
        /// </summary>
        [Input("resourceTypeId", required: true)]
        public Input<string> ResourceTypeId { get; set; } = null!;

        public ApiExtensionTriggerArgs()
        {
        }
        public static new ApiExtensionTriggerArgs Empty => new ApiExtensionTriggerArgs();
    }
}
