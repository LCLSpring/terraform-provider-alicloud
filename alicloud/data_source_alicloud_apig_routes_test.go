// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"strings"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
)

func TestAccAlicloudApigRouteDataSource(t *testing.T) {
	t.Skip("requires alicloud_apig_service; will be re-enabled after apig_service resource is released")
	testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-hangzhou"})
	rand := acctest.RandIntRange(1000000, 9999999)

	idsConf := dataSourceTestAccConfig{
		existConfig: testAccCheckAlicloudApigRouteSourceConfig(rand, map[string]string{
			"ids":         `["${alicloud_apig_route.default.id}"]`,
			"http_api_id": `"${alicloud_apig_http_api.route_httpapi_arr.id}"`,
		}),
		fakeConfig: testAccCheckAlicloudApigRouteSourceConfig(rand, map[string]string{
			"ids":         `["${alicloud_apig_route.default.id}_fake"]`,
			"http_api_id": `"${alicloud_apig_http_api.route_httpapi_arr.id}"`,
		}),
	}

	RouteNameConf := dataSourceTestAccConfig{
		existConfig: testAccCheckAlicloudApigRouteSourceConfig(rand, map[string]string{
			"ids":         `["${alicloud_apig_route.default.id}"]`,
			"route_name":  `"${var.name}"`,
			"http_api_id": `"${alicloud_apig_http_api.route_httpapi_arr.id}"`,
		}),
		fakeConfig: testAccCheckAlicloudApigRouteSourceConfig(rand, map[string]string{
			"ids":         `["${alicloud_apig_route.default.id}_fake"]`,
			"route_name":  `"${var.name}_fake"`,
			"http_api_id": `"${alicloud_apig_http_api.route_httpapi_arr.id}"`,
		}),
	}
	HttpApiIdConf := dataSourceTestAccConfig{
		existConfig: testAccCheckAlicloudApigRouteSourceConfig(rand, map[string]string{
			"ids":         `["${alicloud_apig_route.default.id}"]`,
			"http_api_id": `"${alicloud_apig_http_api.route_httpapi_arr.id}"`,
		}),
		fakeConfig: testAccCheckAlicloudApigRouteSourceConfig(rand, map[string]string{
			"ids":         `["${alicloud_apig_route.default.id}_fake"]`,
			"http_api_id": `"${alicloud_apig_http_api.route_httpapi_arr.id}"`,
		}),
	}

	allConf := dataSourceTestAccConfig{
		existConfig: testAccCheckAlicloudApigRouteSourceConfig(rand, map[string]string{
			"ids":        `["${alicloud_apig_route.default.id}"]`,
			"route_name": `"${var.name}"`,

			"http_api_id": `"${alicloud_apig_http_api.route_httpapi_arr.id}"`,
		}),
		fakeConfig: testAccCheckAlicloudApigRouteSourceConfig(rand, map[string]string{
			"ids":        `["${alicloud_apig_route.default.id}_fake"]`,
			"route_name": `"${var.name}_fake"`,

			"http_api_id": `"${alicloud_apig_http_api.route_httpapi_arr.id}"`,
		}),
	}

	ApigRouteCheckInfo.dataSourceTestCheck(t, rand, idsConf, RouteNameConf, HttpApiIdConf, allConf)
}

var existApigRouteMapFunc = func(rand int) map[string]string {
	return map[string]string{
		"routes.#":                    "1",
		"routes.0.status":             CHECKSET,
		"routes.0.description":        CHECKSET,
		"routes.0.create_time":        CHECKSET,
		"routes.0.domain_infos.#":     CHECKSET,
		"routes.0.route_id":           CHECKSET,
		"routes.0.match.#":            CHECKSET,
		"routes.0.gateway_status.%":   CHECKSET,
		"routes.0.backend.#":          CHECKSET,
		"routes.0.environment_info.#": CHECKSET,
		"routes.0.update_time":        CHECKSET,
		"routes.0.builtin":            CHECKSET,
	}
}

var fakeApigRouteMapFunc = func(rand int) map[string]string {
	return map[string]string{
		"routes.#": "0",
	}
}

var ApigRouteCheckInfo = dataSourceAttr{
	resourceId:   "data.alicloud_apig_routes.default",
	existMapFunc: existApigRouteMapFunc,
	fakeMapFunc:  fakeApigRouteMapFunc,
}

func testAccCheckAlicloudApigRouteSourceConfig(rand int, attrMap map[string]string) string {
	var pairs []string
	for k, v := range attrMap {
		pairs = append(pairs, k+" = "+v)
	}
	config := fmt.Sprintf(`
variable "name" {
	default = "tf-testAccApigRoute%d"
}
resource "alicloud_apig_http_api" "route_httpapi_arr" {
  http_api_name = "cspec-route-arr-httpapi"
  protocols     = ["HTTP"]
  type          = "Rest"
  description   = "array test httpapi"
  base_path     = "/cspec-route-arr"
}



resource "alicloud_apig_route" "default" {
  backend {
    services {
      version    = "v1.0"
      port       = "8080"
      protocol   = "HTTP"
      weight     = "70"
      service_id = "svc-a"
    }
    services {
      version    = "v1.0"
      port       = "8081"
      protocol   = "HTTP"
      weight     = "30"
      service_id = "svc-b"
    }
    scene = "SingleService"
  }
  description = "array route description"
  route_name  = "cspec-arr-route"
  http_api_id = alicloud_apig_http_api.route_httpapi_arr.id
  match {
    path {
      type  = "Prefix"
      value = "/arr-path"
    }
    headers {
      type  = "Exact"
      value = "v1"
      name  = "h1"
    }
    headers {
      type  = "Prefix"
      value = "v2"
      name  = "h2"
    }
    query_params {
      type  = "Exact"
      value = "v1"
      name  = "q1"
    }
    query_params {
      type  = "Prefix"
      value = "v2"
      name  = "q2"
    }
    methods         = ["GET", "POST", "PUT"]
    ignore_uri_case = false
  }
}

data "alicloud_apig_routes" "default" {
%s
}
`, rand, strings.Join(pairs, "\n   "))
	return config
}
