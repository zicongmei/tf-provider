package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	item1 = "item1"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"base_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ZICONG_BASE_URL", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			item1: resourceItem1(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			item1: dataSourceExample(),
		},
	}
}
