package provider

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceItem1() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRead,
		Schema: map[string]*schema.Schema{
			"things": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"version": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceRead(d *schema.ResourceData, m any) error {
	const jsonStream = `
		[
			{"id": 1, "version": "a"},
			{"id": 2, "version": "b"},
			{"id": 3, "version": "c"},
		]`

	things := make([]map[string]any, 0)

	err := json.NewDecoder(strings.NewReader(jsonStream)).Decode(&things)
	if err != nil {
		return err
	}

	if err := d.Set("things", things); err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return nil
}
