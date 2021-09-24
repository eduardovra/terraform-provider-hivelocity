---
layout: ""
page_title: "Provider: Hivelocity"
description: |-
  The Hivelocity provider provides resources to interact with the Hivelocity API.
---

# Hivelocity Provider

The Hivelocity provider provides resources to interact with the Hivelocity API.

## Example Usage

```terraform
terraform {
  required_providers {
    hivelocity = {
      versions = ["0.1.0"]
      source   = "hivelocity/hivelocity"
    }
  }
}

// Find a plan with 16GB of memory in Tampa.
data "hivelocity_product" "tampa_product" {
  first = true

  filter {
    name   = "product_memory"
    values = ["16GB"]
  }

  filter {
    name   = "data_center"
    values = ["TPA1"]
  }

  filter {
    name   = "stock"
    values = ["limited", "available"]
  }
}

data "hivelocity_ssh_key" "ssh_keys" {
  first = true

  filter {
    name   = "name"
    values = ["This is my Terraform SSH Key"]
  }
}

// Provision your device with CentOS 7.
resource "hivelocity_bare_metal_device" "tampa_server" {
  product_id        = data.hivelocity_product.tampa_product.product_id
  os_name           = "CentOS 7.x"
  location_name     = data.hivelocity_product.tampa_product.data_center
  hostname          = "hivelocity.terraform.test"
  tags              = ["hello", "world"]
  script            = file("${path.module}/cloud_init_example.yaml")
  public_ssh_key_id = data.hivelocity_ssh_key.ssh_keys.ssh_key_id
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **api_key** (String, Sensitive) Your API Key from the https://my.hivelocity.net portal.

### Optional

- **api_url** (String) The API instance to communicate with defaults to https://core.hivelocity.net/api/v2