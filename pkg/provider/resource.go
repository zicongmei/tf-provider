package provider

import (
	"time"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceItem1() *schema.Resource {
	return &schema.Resource{
		Create: resourceCreate,
		Read:   resourceRead,
		Update: resourceUpdate,
		Delete: resourceDelete,

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
			"foo": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bar": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"number": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"uuid": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceCreate(d *schema.ResourceData, m any) error {
	d.SetId("123")
	d.Set("last_updated", time.Now().Format(time.RFC850))
	d.Set("created", time.Now().Format(time.RFC850))

	return nil
}

func resourceRead(d *schema.ResourceData, m any) error {
	foo := d.Get("foo")

	for _, f := range foo.([]any) {
		f := f.(map[string]any)

		for _, b := range f["bar"].([]any) {
			b := b.(map[string]any)
			b["uuid"] = uuid.New().String()
			b["number"] = 123
		}
	}

	if err := d.Set("foo", foo); err != nil {
		return err
	}

	d.Set("last_updated", time.Now().Format(time.RFC850))

	return nil
}

func resourceUpdate(d *schema.ResourceData, m any) error {

	if d.HasChange("foo") {
		d.Set("last_updated", time.Now().Format(time.RFC850))
	}

	return nil
}

func resourceDelete(d *schema.ResourceData, m any) error {
	d.SetId("")
	return nil
}
