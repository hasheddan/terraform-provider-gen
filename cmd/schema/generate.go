package schema

import (
	"github.com/crossplane-contrib/terraform-provider-gen/pkg/generator"
	"github.com/crossplane-contrib/terraform-runtime/pkg/client"
)

func GenerateSchema(onlyGenerateResourceFlag *string, provider *client.Provider) error {
	gen, err := generator.NewSchemaGenerator(provider, generator.WithResourceName(onlyGenerateResourceFlag))
	if err != nil {
		return err
	}
	return gen.Generate()
}
