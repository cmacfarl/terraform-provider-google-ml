package ml_datasource

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"google.golang.org/api/ml/v1"
	"google.golang.org/api/option"
)

func dataSourceMlConfig() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceMlConfigRead,
		Schema: map[string]*schema.Schema {
			"service_account": &schema.Schema {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_account_project": &schema.Schema {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"tpu_service_account": &schema.Schema {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceMlConfigRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	if v, ok := d.GetOk("credentials"); ok {
		credentials = v.(string)

	} else if filename := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"); filename != "" {
		credentials = filename
	}

	mlService, err := ml.NewService(ctx, option.WithCredentialsFile(credentials))
	if err != nil {
		return diag.FromErr(err)
	}

	mlProjectsService := mlService.Projects
	configCall := mlProjectsService.GetConfig("projects/" + project)
	res, err := configCall.Do()
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("service_account", res.ServiceAccount); err != nil {
		diags = append(diags, diag.Diagnostic{ 
			Severity: diag.Warning, Summary: err.Error()})
		return diags
	}
	if err := d.Set("service_account_project", res.ServiceAccountProject); err != nil {
		diags = append(diags, diag.Diagnostic{ 
			Severity: diag.Warning, Summary: err.Error()})
		return diags
	}
	if err := d.Set("tpu_service_account", res.Config.TpuServiceAccount); err != nil {
		diags = append(diags, diag.Diagnostic{ 
			Severity: diag.Warning, Summary: err.Error()})
		return diags
	}
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
