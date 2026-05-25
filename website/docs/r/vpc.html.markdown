---
subcategory: "VPC"
layout: "alicloud"
page_title: "Alicloud: alicloud_vpc"
description: |-
  Provides a Alicloud VPC VPC resource.
---

# alicloud_vpc

Provides a VPC VPC resource.

A VPC instance creates a VPC. You can fully control your own VPC, such as selecting IP address ranges, configuring routing tables, and gateways. You can use Alibaba cloud resources such as cloud servers, apsaradb for RDS, and load balancer in your own VPC. 

-> **NOTE:** This resource will auto build a router and a route table while it uses `alicloud_vpc` to build a vpc resource. 

-> **NOTE:** Currently, the IPv4 / IPv6 dual-stack VPC function is under public testing. Only the following regions support IPv4 / IPv6 dual-stack VPC: `cn-hangzhou`, `cn-shanghai`, `cn-shenzhen`, `cn-beijing`, `cn-huhehaote`, `cn-hongkong` and `ap-southeast-1`, and need to apply for public beta qualification. To use, please [submit an application](https://help.aliyun.com/document_detail/100334.html).

## Module Support

You can use the existing [vpc module](https://registry.terraform.io/modules/alibaba/vpc/alicloud) 
to create a VPC and several VSwitches one-click.

For information about VPC VPC and how to use it, see [What is VPC](https://www.alibabacloud.com/help/en/virtual-private-cloud/latest/what-is-a-vpc).

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}


resource "alicloud_vpc" "default" {
  ipv6_isp    = "BGP"
  description = "test"
  cidr_block  = "10.0.0.0/8"
  vpc_name    = var.name
  enable_ipv6 = true
}
```

## Argument Reference

The following arguments are supported:
* `cidr_block` - (Optional, Computed) The CIDR block of the VPC.

  - You can specify one of the following CIDR blocks or their subsets as the primary IPv4 CIDR block of the VPC: 192.168.0.0/16, 172.16.0.0/12, and 10.0.0.0/8. These CIDR blocks are standard private CIDR blocks as defined by Request for Comments (RFC) documents. The subnet mask must be 8 to 28 bits in length.
  - You can also use a custom CIDR block other than 100.64.0.0/10, 224.0.0.0/4, 127.0.0.0/8, 169.254.0.0/16, and their subnets as the primary IPv4 CIDR block of the VPC.
* `classic_link_enabled` - (Optional) The status of ClassicLink function.
* `description` - (Optional) The new description of the VPC.
The description must be 1 to 256 characters in length, and cannot start with `http://` or `https://`.
* `dns_hostname_status` - (Optional, Computed) The status of VPC DNS Hostname
* `dry_run` - (Optional, Available since v1.119.0) Whether to PreCheck only this request. Value:
  - `true`: The check request is sent without creating a VPC. Check items include whether required parameters, request format, and business restrictions are filled in. If the check does not pass, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.
  - `false` (default): Sends a normal request, returns an HTTP 2xx status code and directly creates a VPC.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `enable_ipv6` - (Optional, Computed, Available since v1.119.0) Whether to enable the IPv6 network segment. Value:
  - `false` (default): Not enabled.
  - `true`: enabled.
* `force_delete` - (Optional) Force delete vpc or not.

-> **NOTE:** This parameter configures deletion behavior and is only evaluated when Terraform attempts to destroy the resource. Changes to this parameter during updates are stored but have no immediate effect.

* `ipv4_cidr_mask` - (Optional, Int) Allocate VPC from The IPAM address pool by entering a mask.

-> **NOTE:**  when you specify the IPAM address pool to create a VPC, enter at least one of the CidrBlock or Ipv4CidrMask parameters.


-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `ipv4_ipam_pool_id` - (Optional) The ID of the IP Address Manager (IPAM) pool that contains IPv4 addresses.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `ipv6_cidr_block` - (Optional, Computed) The IPv6 CIDR block of the default VPC.

-> **NOTE:**  When `EnableIpv6` is set to `true`, this parameter is required.

* `ipv6_cidr_mask` - (Optional, Int, Available since v1.280.0) Add an IPv6 CIDR block from the IPAM pool to the VPC by entering a mask.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `ipv6_ipam_pool_id` - (Optional, Available since v1.280.0) The ID of the IP Address Manager (IPAM) pool that contains IPv6 addresses.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `ipv6_isp` - (Optional) The IPv6 address segment type of the VPC. Value:
  - `BGP` (default): Alibaba Cloud BGP IPv6.
  - `ChinaMobile`: China Mobile (single line).
  - `ChinaUnicom`: China Unicom (single line).
  - `ChinaTelecom`: China Telecom (single line).

-> **NOTE:**  If a single-line bandwidth whitelist is enabled, this field can be set to `ChinaTelecom` (China Telecom), `ChinaUnicom` (China Unicom), or `ChinaMobile` (China Mobile).


-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `resource_group_id` - (Optional, Computed, Available since v1.115) The ID of the resource group to which you want to move the resource.

-> **NOTE:**   You can use resource groups to facilitate resource grouping and permission management for an Alibaba Cloud. For more information, see [What is resource management?](https://www.alibabacloud.com/help/en/doc-detail/94475.html)

* `route_table_id` - (Optional, ForceNew, Computed) The ID of the system route table.
* `secondary_cidr_blocks` - (Optional, Computed, List, Deprecated since v1.185.0) Field 'secondary_cidr_blocks' has been deprecated from provider version 1.185.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_ipv4_cidr_block'. `secondary_cidr_blocks` attributes and `alicloud_vpc_ipv4_cidr_block` resource cannot be used at the same time.
* `system_route_table_description` - (Optional) The description of the route table.
The description must be 1 to 256 characters in length, and cannot start with `http://` or `https://`.
* `system_route_table_name` - (Optional) The name of the route table.
The name must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
* `system_route_table_route_propagation_enable` - (Optional, Computed) Whether the system route table receives propagation routes.
* `tags` - (Optional, Map, Available since v1.55.3) The tags of Vpc.
* `user_cidrs` - (Optional, ForceNew, Computed, List, Available since v1.119.0) A list of user CIDRs.
* `vpc_name` - (Optional, Computed, Available since v1.119.0) The new name of the VPC.
The name must be 1 to 128 characters in length and cannot start with `http://` or `https://`.

The following arguments will be discarded. Please use new fields as soon as possible:
* `name` - (Deprecated since v1.119.0). Field 'name' has been deprecated from provider version 1.119.0. New field 'vpc_name' instead.
* `router_table_id` - (Deprecated since v1.280.0). Field 'router_table_id' has been deprecated from provider version 1.280.0. New field 'route_table_id' instead.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - The creation time of the VPC.
* `dhcp_options_set_id` - The ID of the DHCP options set.
* `ipv6_cidr_blocks` - The IPv6 CIDR block information of the VPC.
  * `ipv6_cidr_block` - The IPv6 CIDR block of the VPC.
  * `ipv6_isp` - Valid values: **BGP** (default): Alibaba Cloud BGP IPv6.
* `is_default` - Specifies whether to create the default VPC in the specified region.
* `region_id` - The ID of the region where the VPC is located.
* `router_id` - The region ID of the VPC to which the route table belongs.
* `status` - The status of the VPC.
* `vswitch_ids` - A list of VSwitches in the VPC.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 10 mins) Used when create the VPC.
* `delete` - (Defaults to 5 mins) Used when delete the VPC.
* `update` - (Defaults to 5 mins) Used when update the VPC.

## Import

VPC VPC can be imported using the id, e.g.

```shell
$ terraform import alicloud_vpc.example <vpc_id>
```