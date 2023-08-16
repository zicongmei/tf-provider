package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	resourceKey   = "resource_key"
	resourceValue = "resource_value"
)

func resourceItem1() *schema.Resource {
	return &schema.Resource{
		Create: resource1Create,
		Read:   resource1Read,
		Update: resource1Update,
		Delete: resource1Delete,

		Schema: map[string]*schema.Schema{
			"last_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"not_computed_optional": {
				Type:     schema.TypeString,
				Optional: true,
			},
			resourceKey: {
				Type:     schema.TypeString,
				Required: true,
			},
			resourceValue: {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceItem2() *schema.Resource {
	return &schema.Resource{
		Create: resource2Create,
		Read:   resource2Read,
		Update: resource2Update,
		Delete: resource2Delete,

		Schema: map[string]*schema.Schema{
			"last_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"not_computed_optional": {
				Type:     schema.TypeString,
				Optional: true,
			},
			resourceKey: {
				Type:     schema.TypeString,
				Required: true,
			},
			resourceValue: {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
