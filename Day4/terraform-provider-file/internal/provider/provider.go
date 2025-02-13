package provider

import (
	"context"
	//"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	schema.DescriptionKind = schema.StringMarkdown
}

func New(version string) func() *schema.Provider {
	//fmt.Println("New func invoked ...")
	return func() *schema.Provider {
		p := &schema.Provider{
			DataSourcesMap: map[string]*schema.Resource{
				//"localfile": dataSourceLocalFile(),
			},
			ResourcesMap: map[string]*schema.Resource{
				"localfile": resourceLocalFile(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

type FileConfig struct {
	name string
	content string
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (any, diag.Diagnostics) {
	return func(context.Context, *schema.ResourceData) (any, diag.Diagnostics) {
		fileConfig := FileConfig {}
		fileConfig.name = "./test.txt"
		fileConfig.content = "This is a test file created by terraform"

		return &fileConfig, nil
	}
}
