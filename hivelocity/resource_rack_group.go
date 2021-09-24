package hivelocity

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRackGroup() *schema.Resource {
	resourceBareMetalDevice := resourcePseudoBareMetalDevice()

	return &schema.Resource{
		Description: "Rack group resource",
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
		},
		CreateContext: resourceRackGroupCreate,
		ReadContext:   resourceRackGroupRead,
		UpdateContext: resourceRackGroupUpdate,
		DeleteContext: resourceRackGroupDelete,
		CustomizeDiff: resourceRackGroupDiff,
		Schema: map[string]*schema.Schema{
			"bare_metal_devices": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Name identifier for this Bare Metal Device",
							Required:    true,
						},
						"product_id": {
							Type:        resourceBareMetalDevice.Schema["product_id"].Type,
							Description: resourceBareMetalDevice.Schema["product_id"].Description,
							Required:    true,
						},
						"location_name": {
							Type:        resourceBareMetalDevice.Schema["location_name"].Type,
							Description: resourceBareMetalDevice.Schema["location_name"].Description,
							Required:    true,
						},
					},
				},
				Required:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				//Set: schema.SchemaSetFunc
				//DiffSuppressFunc: func(k, old, new string, d *ResourceData) bool
			},
		},
	}
}

func resourceRackGroupCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[ERROR]\n\n >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> resourceRackGroupCreate %#v", d)

	bare_metal_devices := d.Get("bare_metal_devices").(*schema.Set)
	if bare_metal_devices.Len() > 2 {
		return diag.Errorf("There are no racks in stock with all the devices available to fulfill your order.")
	}

	id := rand.Intn(100)
	d.SetId(fmt.Sprintf("%d", id))

	return resourceRackGroupRead(ctx, d, m)
}

func resourceRackGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	log.Printf("[ERROR]\n\n >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> resourceRackGroupRead %#v", d)

	return diags
}

func resourceRackGroupUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	bare_metal_devices := d.Get("bare_metal_devices").(*schema.Set)
	if bare_metal_devices.Len() > 2 {
		old, _ := d.GetChange("bare_metal_devices")

		ctx.Err()

		return diag.Errorf("resourceRackGroupUpdate old %#v", old)
		//d.Set("bare_metal_devices", old_set)

		return diag.Errorf("There are not enough devices available in the current rack to fulfill your order. If you want to deploy new devices in another rack, use `terraform taint` on the resource to force the creation of a new rack group.")
	}

	return resourceRackGroupRead(ctx, d, m)
}

func resourceRackGroupDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	d.SetId("")

	return diags
}

func resourceRackGroupDiff(ctx context.Context, d *schema.ResourceDiff, m interface{}) error {
	//log.Printf("[ERROR] >>>>>>>>>>>>>>>>>>>>>>>>> resourceRackGroupDiff %#v", d)
	//old, new := d.GetChange("bare_metal_devices")
	//devicesToAdd := new.(*schema.Set).Difference(old.(*schema.Set))
	//return fmt.Errorf("%#v", devicesToAdd.List())

	bare_metal_devices := d.Get("bare_metal_devices").(*schema.Set)
	if bare_metal_devices.Len() > 2 {
		// TODO check if there is another rack capable of fulfilling the order
		return errors.New("There are not enough devices available in the current rack to fulfill your order. If you want to deploy new devices in another rack, use `terraform taint` on the resource to force the creation of a new rack group.")
	}
	return errors.New("Please dont change my state")
}
