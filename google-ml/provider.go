package ml_datasource

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var project string
var credentials string

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	project = d.Get("project").(string)
	credentials = d.Get("credentials").(string)

	return nil, nil
}

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"project": &schema.Schema{
				Type: schema.TypeString,
				Required: true,
			},
			"credentials": &schema.Schema{
				Type: schema.TypeString,
				Required: true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"ml_config": dataSourceMlConfig(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}
