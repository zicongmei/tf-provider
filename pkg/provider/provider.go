package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	resource1 = "resource1"
	resource2 = "resource2"
)

func Provider() *schema.Provider {
	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"db_file1": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/tmp/tf_demo_db." + resource1,
			},
			"db_file2": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/tmp/tf_demo_db." + resource2,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			resource1: resourceItem1(),
			resource2: resourceItem2(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			resource1: dataSourceItem1(),
			resource2: dataSourceItem1(),
		},
	}
	provider.ConfigureContextFunc = func(_ context.Context, d *schema.ResourceData) (any, diag.Diagnostics) {
		config := Config{
			DBFile1: d.Get("db_file1").(string),
			DBFile2: d.Get("db_file2").(string),
		}
		return config.Client()
	}
	return provider
}
