package hivelocity

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	swagger "github.com/hivelocity/terraform-provider-hivelocity/hivelocity-client-go"
)

func resourceBareMetalDevice() *schema.Resource {
	return &schema.Resource{
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
		},
		CreateContext: resourceBareMetalDeviceCreate,
		ReadContext:   resourceBareMetalDeviceRead,
		UpdateContext: resourceBareMetalDeviceUpdate,
		DeleteContext: resourceBareMetalDeviceDelete,
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
				Type: schema.TypeString,
				//Computed: true,
				Optional: true,
				Default:  nil,
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
		},
	}
}

// TODO: Test what happens when you change hostname, tags, etc anything that is required.

func resourceBareMetalDeviceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	hv, _ := m.(*Client)

	tags, err := getTags(d)
	if err != nil {
		d.SetId("")
		return diag.FromErr(err)
	}

	payload := swagger.BareMetalDeviceCreate{
		ProductId:      int32(d.Get("product_id").(int)),
		Hostname:       d.Get("hostname").(string),
		OsName:         d.Get("os_name").(string),
		VlanId:         int32(d.Get("vlan_id").(int)),
		LocationName:   d.Get("location_name").(string),
		Script:         d.Get("script").(string),
		Period:         d.Get("period").(string),
		PublicSshKeyId: int32(d.Get("public_ssh_key_id").(int)),
	}

	bareMetalDevice, _, err := hv.client.BareMetalDevicesApi.PostBareMetalDeviceResource(hv.auth, payload, nil)
	if err != nil {
		d.SetId("")
		myErr := err.(swagger.GenericSwaggerError)
		return diag.Errorf("POST /bare-metal-device failed! (%s)\n\n %s", err, myErr.Body())
	}

	_, err = waitForOrder(d, hv, bareMetalDevice.OrderId)
	if err != nil {
		d.SetId("")
		if strings.Contains(fmt.Sprint(err), "'cancelled'") {
			return diag.Errorf("Your deployment (order %s) has been 'cancelled'. Please contact Hivelocity support if you believe this is a mistake.", fmt.Sprint(bareMetalDevice.OrderId))
		}
		return diag.Errorf("error provisioning order %s. The Hivelocity team will investigate:\n\n%s", fmt.Sprint(bareMetalDevice.OrderId), err)
	}

	device, err := waitForDevice(d, hv, bareMetalDevice.OrderId)
	if err != nil {
		d.SetId("")
		return diag.Errorf("error finding devices for order %s. The Hivelocity team will investigate:\n\n%s", fmt.Sprint(bareMetalDevice.OrderId), err)
	}

	newDeviceId := device.(swagger.BareMetalDevice).DeviceId
	_, err = updateTagsForCreate(hv, newDeviceId, tags)
	if err != nil {
		// TODO: The deployment was successful, so we should throw a warning here that tags failed for some reason.
	}
	d.SetId(fmt.Sprint(newDeviceId))

	return resourceBareMetalDeviceRead(ctx, d, m)
}

func resourceBareMetalDeviceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	hv, _ := m.(*Client)

	deviceId, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	deviceResponse, _, err := hv.client.BareMetalDevicesApi.GetBareMetalDeviceIdResource(hv.auth, int32(deviceId), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("device_id", deviceResponse.DeviceId)
	d.Set("hostname", deviceResponse.Hostname)
	d.Set("location_name", deviceResponse.LocationName)
	d.Set("order_id", deviceResponse.OrderId)
	d.Set("os_name", deviceResponse.OsName)
	d.Set("power_status", deviceResponse.PowerStatus)
	d.Set("primary_ip", deviceResponse.PrimaryIp)
	d.Set("product_id", deviceResponse.ProductId)
	d.Set("product_name", deviceResponse.ProductName)
	d.Set("service_id", deviceResponse.ServiceId)
	d.Set("tags", deviceResponse.Tags)
	d.Set("vlan_d", deviceResponse.VlanId)
	d.Set("public_ssh_key_id", deviceResponse.PublicSshKeyId)
	d.Set("script", deviceResponse.Script)

	if deviceResponse.LocationName == "DEV-TPA1" {
		d.Set("location_name", "TPA1")
	}

	return diags
}

func resourceBareMetalDeviceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	hv, _ := m.(*Client)

	deviceId, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	payload := swagger.BareMetalDeviceUpdate{}
	reload_required := false

	tags, err := getTags(d)
	if err != nil {
		return diag.FromErr(err)
	}
	payload.Tags = tags

	hostname := d.Get("hostname").(string)
	payload.Hostname = hostname
	if d.HasChange("hostname") {
		reload_required = true
	}

	osName := d.Get("os_name").(string)
	payload.OsName = osName
	if d.HasChange("os_name") {
		reload_required = true
	}

	publicSshKeyId := d.Get("public_ssh_key_id").(int)
	payload.PublicSshKeyId = int32(publicSshKeyId)
	if d.HasChange("public_ssh_key_id") {
		reload_required = true
	}

	script := d.Get("script").(string)
	payload.Script = script
	if d.HasChange("script") {
		reload_required = true
	}

	if d.HasChange("vlan_id") {
		// TODO: Currently no-op until VLAN IDS deployed
	}

	// If a reload is required, it's necessary to turn the device off first
	if reload_required {
		devicePower, _, err := hv.client.DeviceApi.GetPowerResource(hv.auth, int32(deviceId), nil)
		if err != nil {
			myErr := err.(swagger.GenericSwaggerError)
			return diag.Errorf("GET /device/%s/power failed! (%s)\n\n %s", fmt.Sprint(deviceId), err, myErr.Body())
		}

		if devicePower.PowerStatus == "ON" {
			_, _, err = hv.client.DeviceApi.PostPowerResource(hv.auth, int32(deviceId), "shutdown", nil)
			if err != nil {
				myErr := err.(swagger.GenericSwaggerError)
				return diag.Errorf("POST /device/%s/power failed! (%s)\n\n %s", fmt.Sprint(deviceId), err, myErr.Body())
			}

			// Power status will transition to PENDING, then OFF
			_, err := waitForDevicePowerOff(d, hv, int32(deviceId))
			if err != nil {
				return diag.Errorf("error powering off device %s. The Hivelocity team will investigate:\n\n%s", fmt.Sprint(deviceId), err)
			}
		}
	}

	_, _, err = hv.client.BareMetalDevicesApi.PutBareMetalDeviceIdResource(hv.auth, int32(deviceId), payload, nil)
	if err != nil {
		myErr := err.(swagger.GenericSwaggerError)
		return diag.Errorf("PUT /bare-metal-device/%s failed! (%s)\n\n %s", fmt.Sprint(deviceId), err, myErr.Body())
	}

	if reload_required {
		_, err := waitForDeviceReload(d, hv, int32(deviceId))
		if err != nil {
			return diag.Errorf("error reloading device %s. The Hivelocity team will investigate:\n\n%s", fmt.Sprint(deviceId), err)
		}
	}

	d.Set("last_updated", time.Now().Format(time.RFC850))
	return resourceBareMetalDeviceRead(ctx, d, m)
}

func resourceBareMetalDeviceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	hv, _ := m.(*Client)

	deviceId, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	// Check device exists still, if not mark as already destroyed.
	_, _, err = hv.client.BareMetalDevicesApi.GetBareMetalDeviceIdResource(hv.auth, int32(deviceId), nil)
	if err != nil {
		d.SetId("")
		return diags
	}

	_, err = hv.client.BareMetalDevicesApi.DeleteBareMetalDeviceIdResource(hv.auth, int32(deviceId), nil)
	if err != nil {
		myErr := err.(swagger.GenericSwaggerError)
		return diag.Errorf("DELETE /bare-metal-device/%s failed! (%s)\n\n %s", fmt.Sprint(deviceId), err, myErr.Body())
	}

	d.SetId("")

	return diags
}
