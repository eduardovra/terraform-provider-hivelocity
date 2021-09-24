terraform {
  required_providers {
    hivelocity = {
      source   = "hivelocity/hivelocity"
      version  = "0.1.0"
    }
  }
}

variable "devices_rack_group" {
  type = any
  default = {
    database = {
      product_id    = 525
      os_name       = "CentOS 7.x"
      hostname      = "database.terraform.test"
      location_name = "DAL1"
    },
    storage = {
      product_id    = 580
      os_name       = "CentOS 7.x"
      hostname      = "storage.terraform.test"
      location_name = "DAL1"
    },
    webserver = {
      product_id    = 525
      os_name       = "CentOS 7.x"
      hostname      = "webserver.terraform.test"
      location_name = "DAL1"
    },
  }
}

resource "hivelocity_rack_group" "web_rack" {
  bare_metal_devices = [for name, device in var.devices_rack_group: {
    name          = name
    product_id    = device.product_id
    location_name = device.location_name
  }]
}

resource "hivelocity_pseudo_bare_metal_device" "devices" {
  for_each = var.devices_rack_group

  product_id    = each.value.product_id
  os_name       = each.value.os_name
  hostname      = each.value.hostname
  location_name = each.value.location_name
  rack_group    = hivelocity_rack_group.web_rack.id
}

output "output_rack_group" {
  value = hivelocity_rack_group.web_rack
}

output "output_devices" {
  value = hivelocity_pseudo_bare_metal_device.devices
}
