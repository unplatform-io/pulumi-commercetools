// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Commercetools.Inputs
{

    public sealed class ProjectSettingsMessagesGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the number of days each Message should be available via the Messages Query API
        /// </summary>
        [Input("deleteDaysAfterCreation")]
        public Input<int>? DeleteDaysAfterCreation { get; set; }

        /// <summary>
        /// When true the creation of messages on the Messages Query HTTP API is enabled
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        public ProjectSettingsMessagesGetArgs()
        {
        }
        public static new ProjectSettingsMessagesGetArgs Empty => new ProjectSettingsMessagesGetArgs();
    }
}
