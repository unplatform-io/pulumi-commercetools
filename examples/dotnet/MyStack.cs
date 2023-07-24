using Pulumi.Commercetools;
using Pulumi.Commercetools.Inputs;

namespace dotnet;

class MyStack : Stack
{
    public MyStack()
    {
        ProvisionProjectSettings();
        ProvisionDefaultProductType();
    }

    private void ProvisionProjectSettings()
    {
        var _ = new ProjectSettings("Unplatform storefront", new ProjectSettingsArgs
        {
            EnableSearchIndexProducts = true,
            Currencies = new List<string>
            {
                "EUR",
                "USD"
            },
            Languages = new List<string>
            {
                "de-DE",
                "en-US"
            },
            Countries = new List<string>
            {
                "DE",
                "US",
                "NL",
            }
        });
    }

    private static void ProvisionDefaultProductType()
    {
        var _ = new ProductType("Product", new ProductTypeArgs
        {
            Name = "Product",
            Description = "Product",
            Key = "product",
            Attributes = new List<ProductTypeAttributeArgs>
            {
                new()
                {
                    Name = "brand",
                    Constraint = "SameForAll",
                    Type = new ProductTypeAttributeTypeArgs
                    {
                        Name = "text"
                    },
                    Searchable = true,
                    Label = new Dictionary<string, object>
                    {
                        { "en-US", "Brand" }
                    }
                },
                new()
                {
                    Name = "color",
                    Constraint = "Unique",
                    Type = new ProductTypeAttributeTypeArgs
                    {
                        Name = "text"
                    },
                    Searchable = true,
                    Label = new Dictionary<string, object>
                    {
                        { "en-US", "Color" }
                    }
                },
                new()
                {
                    Name = "from-price",
                    Constraint = "None",
                    Type = new ProductTypeAttributeTypeArgs
                    {
                        Name = "money"
                    },
                    Searchable = false,
                    Label = new Dictionary<string, object>
                    {
                        { "en-US", "From price" }
                    }
                },
                new()
                {
                    Name = "featured-product-tag",
                    Constraint = "SameForAll",
                    Type = new ProductTypeAttributeTypeArgs
                    {
                        Name = "text"
                    },
                    Searchable = true,
                    Label = new Dictionary<string, object>
                    {
                        { "en-US", "Featured Product Tag" }
                    }
                },
                new()
                {
                    Name = "similar-products",
                    Constraint = "SameForAll",
                    Type = new ProductTypeAttributeTypeArgs
                    {
                        Name = "set",
                        ElementType2 = new ProductTypeAttributeTypeElementType2Args
                        {
                            Name = "reference",
                            ReferenceTypeId = "product",
                        }
                    },
                    Searchable = false,
                    Label = new Dictionary<string, object>
                    {
                        { "en-US", "Similar Products" }
                    },
                },
            }
        });
    }
}