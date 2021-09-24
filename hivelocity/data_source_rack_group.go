package hivelocity

import (
	"context"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceRackGroup() *schema.Resource {
	return &schema.Resource{
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
		},
		ReadContext: dataSourceRackGroupRead,
		Schema: map[string]*schema.Schema{
			"location_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"product_quantities": &schema.Schema{
				Type: schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Required: true,
			},
		},
	}
}

func dataSourceRackGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	quantities := d.Get("product_quantities")

	log.Printf("[ERROR] >>>>>>>>>>>>>>>>>>>>>>>>>> product_quantities %#v", quantities)

	d.SetId("2")

	return diags
}
