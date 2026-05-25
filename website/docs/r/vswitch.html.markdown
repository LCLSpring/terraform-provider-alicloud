---
subcategory: "VPC"
layout: "alicloud"
page_title: "Alicloud: alicloud_vswitch"
description: |-
  Provides a Alicloud VPC Vswitch resource.
---

# alicloud_vswitch

Provides a VPC Vswitch resource.

## Module Support

You can use to the existing [vpc module](https://registry.terraform.io/modules/alibaba/vpc/alicloud)  to create a VPC and several VSwitches one-click.

For information about VPC Vswitch and how to use it, see [What is Vswitch](https://www.alibabacloud.com/help/en/virtual-private-cloud/latest/work-with-vswitches).

## Example Usage

Basic Usage

```terraform
data "alicloud_zones" "foo" {
  available_resource_creation = "VSwitch"
}

resource "alicloud_vpc" "foo" {
  vpc_name   = "terraform-example"
  cidr_block = "172.16.0.0/12"
}

resource "alicloud_vswitch" "foo" {
  vswitch_name = "terraform-example"
  cidr_block   = "172.16.0.0/21"
  vpc_id       = alicloud_vpc.foo.id
  zone_id      = data.alicloud_zones.foo.zones.0.id
}
```

## Argument Reference

The following arguments are supported:
* `cidr_block` - (Required, ForceNew) The region ID of the vSwitch.
You can call the [DescribeRegions](https://www.alibabacloud.com/help/en/doc-detail/36063.html) operation to query the most recent region list.
* `description` - (Optional) The new description for the vSwitch.
The description must be 1 to 256 characters in length and cannot start with `http://` or `https://`.
* `enable_ipv6` - (Optional, Available since v1.201.1) Specifies whether to query vSwitches with IPv6 enabled in the region. Valid values:

  - `true`
  - `false`

If you do not set this parameter, the system queries all vSwitches in the specified region by default.

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `ipv6_cidr_block` - (Optional, ForceNew, Computed) The last eight bits of the IPv6 CIDR block of the vSwitch. Valid values: `0` to `255`.
* `ipv6_cidr_block_mask` - (Optional, Computed, Int, Available since v1.115) The CIDR block of the vSwitch. Take note of the following limits:

  - The subnet mask of the CIDR block must be 16 to 29 bits in length.
  - The CIDR block of the vSwitch must fall within the CIDR block of the VPC to which the vSwitch belongs.
  - The CIDR block of a vSwitch cannot be the same as the destination CIDR block in a route entry of the VPC. However, it can be a subset of the destination CIDR block.

-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `tags` - (Optional, Map, Available since v1.55.3) The tags of VSwitch.
* `vswitch_name` - (Optional, Computed, Available since v1.119.0) The new name for the vSwitch.
The name must be 1 to 128 characters in length, and cannot start with `http://` or `https://`.
* `vpc_id` - (Required, ForceNew) The ID of the virtual private cloud (VPC) to which the vSwitches belong.

-> **NOTE:**   You must set at least one of `RegionId` and `VpcId`.

* `vpc_ipv6_cidr_block` - (Optional, Available since v1.280.0) The IPv6 CIDR block of the VPC. If the VPC to which the vSwitch belongs has multiple IPv6 CIDR blocks, you can enter this parameter to specify the IPv6 CIDR block range to which the vSwitch belongs. If this parameter is not specified, the IPv6 CIDR block allocated when the VPC enables IPv6 is selected.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `zone_id` - (Optional, ForceNew) The ID of the zone to which the vSwitches belong. You can call the [DescribeZones](https://www.alibabacloud.com/help/en/doc-detail/36064.html) operation to query the most recent zone list.

The following arguments will be discarded. Please use new fields as soon as possible:
* `name` - (Deprecated since v1.119.0). Field 'name' has been deprecated from provider version 1.119.0. New field 'vswitch_name' instead.
* `availability_zone` - (Deprecated since v1.119.0). Field 'availability_zone' has been deprecated from provider version 1.119.0. New field 'zone_id' instead.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `available_ip_address_count` - The number of available IP addresses.
* `create_time` - The creation time of the VSwitch.
* `is_default` - Specifies whether to query the default vSwitches in the specified region.
* `network_acl_id` - The ID of the network ACL.
* `route_table_id` - The ID of the route table.
* `status` - The status of the resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Vswitch.
* `delete` - (Defaults to 5 mins) Used when delete the Vswitch.
* `update` - (Defaults to 5 mins) Used when update the Vswitch.

## Import

VPC Vswitch can be imported using the id, e.g.

```shell
$ terraform import alicloud_vswitch.example <vswitch_id>
```