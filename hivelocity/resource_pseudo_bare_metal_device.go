package hivelocity

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePseudoBareMetalDevice() *schema.Resource {
	return &schema.Resource{
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
		},
		CreateContext: resourcePseudoBareMetalDeviceCreate,
		ReadContext:   resourcePseudoBareMetalDeviceRead,
		UpdateContext: resourcePseudoBareMetalDeviceUpdate,
		DeleteContext: resourcePseudoBareMetalDeviceDelete,
		CustomizeDiff: resourcePseudoBareMetalDeviceDiff,
		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"device_id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"order_id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"service_id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"product_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"product_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"os_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"location_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"power_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"vlan_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  nil,
			},
			"primary_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"tags": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"script": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"period": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"public_ssh_key_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  nil,
			},
			"rack_group": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  nil,
				ForceNew: true,
			},
		},
	}
}

func resourcePseudoBareMetalDeviceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[ERROR]\n\n >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> resourcePseudoBareMetalDeviceCreate %#v", d)
	id := rand.Intn(100)
	d.SetId(fmt.Sprintf("%d", id))

	//rack_group := d.Get("rack_group")
	//log.Printf("[ERROR]\n\n >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> rack_group %#v", rack_group)

	return resourcePseudoBareMetalDeviceRead(ctx, d, m)
}

func resourcePseudoBareMetalDeviceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}

func resourcePseudoBareMetalDeviceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourcePseudoBareMetalDeviceRead(ctx, d, m)
}

func resourcePseudoBareMetalDeviceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	d.SetId("")

	return diags
}

func resourcePseudoBareMetalDeviceDiff(ctx context.Context, d *schema.ResourceDiff, m interface{}) error {
	//log.Printf("[ERROR] >>>>>>>>>>>>>>>>>>>>>>>>> resourcePseudoBareMetalDeviceDiff %#v", d)
	return nil
}
