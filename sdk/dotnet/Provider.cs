// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Commercetools
{
    /// <summary>
    /// The provider type for the commercetools package. By default, resources use package-wide configuration
    /// settings, however an explicit `Provider` instance may be created and passed during resource
    /// construction to achieve fine-grained programmatic control over provider settings. See the
    /// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
    /// </summary>
    [CommercetoolsResourceType("pulumi:providers:commercetools")]
    public partial class Provider : global::Pulumi.ProviderResource
    {
        /// <summary>
        /// The API URL of the commercetools platform. https://docs.commercetools.com/http-api
        /// </summary>
        [Output("apiUrl")]
        public Output<string?> ApiUrl { get; private set; } = null!;

        /// <summary>
        /// The OAuth Client ID for a commercetools platform project. https://docs.commercetools.com/http-api-authorization
        /// </summary>
        [Output("clientId")]
        public Output<string?> ClientId { get; private set; } = null!;

        /// <summary>
        /// The OAuth Client Secret for a commercetools platform project. https://docs.commercetools.com/http-api-authorization
        /// </summary>
        [Output("clientSecret")]
        public Output<string?> ClientSecret { get; private set; } = null!;

        /// <summary>
        /// The project key of commercetools platform project. https://docs.commercetools.com/getting-started
        /// </summary>
        [Output("projectKey")]
        public Output<string?> ProjectKey { get; private set; } = null!;

        /// <summary>
        /// A list as string of OAuth scopes assigned to a project key, to access resources in a commercetools platform project.
        /// https://docs.commercetools.com/http-api-authorization
        /// </summary>
        [Output("scopes")]
        public Output<string?> Scopes { get; private set; } = null!;

        /// <summary>
        /// The authentication URL of the commercetools platform. https://docs.commercetools.com/http-api-authorization
        /// </summary>
        [Output("tokenUrl")]
        public Output<string?> TokenUrl { get; private set; } = null!;


        /// <summary>
        /// Create a Provider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Provider(string name, ProviderArgs? args = null, CustomResourceOptions? options = null)
            : base("commercetools", name, args ?? new ProviderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "clientId",
                    "clientSecret",
                    "projectKey",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
    }

    public sealed class ProviderArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The API URL of the commercetools platform. https://docs.commercetools.com/http-api
        /// </summary>
        [Input("apiUrl")]
        public Input<string>? ApiUrl { get; set; }

        [Input("clientId")]
        private Input<string>? _clientId;

        /// <summary>
        /// The OAuth Client ID for a commercetools platform project. https://docs.commercetools.com/http-api-authorization
        /// </summary>
        public Input<string>? ClientId
        {
            get => _clientId;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _clientId = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("clientSecret")]
        private Input<string>? _clientSecret;

        /// <summary>
        /// The OAuth Client Secret for a commercetools platform project. https://docs.commercetools.com/http-api-authorization
        /// </summary>
        public Input<string>? ClientSecret
        {
            get => _clientSecret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _clientSecret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("projectKey")]
        private Input<string>? _projectKey;

        /// <summary>
        /// The project key of commercetools platform project. https://docs.commercetools.com/getting-started
        /// </summary>
        public Input<string>? ProjectKey
        {
            get => _projectKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _projectKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// A list as string of OAuth scopes assigned to a project key, to access resources in a commercetools platform project.
        /// https://docs.commercetools.com/http-api-authorization
        /// </summary>
        [Input("scopes")]
        public Input<string>? Scopes { get; set; }

        /// <summary>
        /// The authentication URL of the commercetools platform. https://docs.commercetools.com/http-api-authorization
        /// </summary>
        [Input("tokenUrl")]
        public Input<string>? TokenUrl { get; set; }

        public ProviderArgs()
        {
        }
        public static new ProviderArgs Empty => new ProviderArgs();
    }
}
