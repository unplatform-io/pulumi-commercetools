[![Actions Status](https://github.com/unplatform-io/pulumi-commercetools/workflows/master/badge.svg)](https://github.com/unplatform-io/pulumi-commercetools/actions)
[![NPM version](https://badge.fury.io/js/%40unplatform%2Fcommercetools.svg)](https://www.npmjs.com/package/@unplatform/commercetools)
[![Python version](https://badge.fury.io/py/pulumi-commercetools.svg)](https://pypi.org/project/pulumi-commercetools)
[![NuGet version](https://badge.fury.io/nu/pulumi.commercetools.svg)](https://badge.fury.io/nu/pulumi.commercetools)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/unplatform-io/pulumi-commercetools/sdk/v2/go)](https://pkg.go.dev/github.com/unplatform-io/pulumi-commercetools/sdk/v2/go)

# CommerceTools provider
The CommerceTools resource provider for Pulumi lets you provision [CommerceTools](https://commercetools.com/) resources.

This provider is a [bridge] (https://github.com/pulumi/pulumi-terraform-bridge) to the excellent [CommerceTools Terraform Provider](https://github.com/labd/terraform-provider-commercetools) from the Labd folks.

To use this package, please [install the Pulumi CLI first](https://pulumi.io/).] 

## Installing

The plugin itself needs to be installed by running:

    `pulumi plugin install --server https://github.com/unplatform-io/pulumi-commercetools/releases/download/v0.0.0 resource commercetools v0.0.0`

Where v0.0.0 is the desired version.

This SDK is available in many languages in the standard packaging formats.

### Node.js (Java/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @unplatform/commercetools

or `yarn`:

    $ yarn add @unplatform/commercetools

### Python

To use from Python, install using `pip`:

    $ pip install pulumi_commercetools

### Go

To use from Go, use `go get` to grab the latest version of the library

    $ go get github.com/unplatform-io/pulumi-commercetools/sdk/v2/go/...

### .NET

To use from .NET, install using `dotnet add package`:

    $ dotnet add package Pulumi.CommerceTools

## Configuration

The following configuration points are available for the `commercetools` provider:

- `commercetools:apiUrl` - the url of the commercetools api (e.g. `https://api.europe-west1.gcp.commercetools.com`)
- `commercetools:tokenUrl` - the url used to authenticate (e.g. `https://auth.europe-west1.gcp.commercetools.com`)
- `commercetools:scopes` - the authentication scopes needed
- `commercetools:clientId` - the client id used to authenticate
- `commercetools:clientSecret` - the client secret used to authenticate
- `commercetools:projectKey` - the key of the commerce tools project

## Reference

For detailed reference documentation, please visit [the docs of the commerce tools terraform provider][1].

[1]: https://commercetools-terraform-provider.readthedocs.io/en/latest/
