package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Apig HttpApi. >>> Resource test cases, automatically generated.
// Case http_api_crud_test 12898
func TestAccAliCloudApigHttpApi_basic12898(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_apig_http_api.default"
	ra := resourceAttrInit(resourceId, AlicloudApigHttpApiMap12898)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &ApigServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeApigHttpApi")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccapig%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudApigHttpApiBasicDependence12898)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"http_api_name": name,
					"protocols": []string{
						"HTTP"},
					"type":              "Rest",
					"description":       "test description for cspec",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"ai_protocols": []string{
						"OpenAI"},
					"base_path":   "/cspec-test",
					"enable_auth": "false",
					"deploy_configs": []string{
						"{\\\"autoDeploy\\\":true,\\\"mock\\\":{\\\"enable\\\":true,\\\"responseContent\\\":\\\"mock response\\\",\\\"statusCode\\\":200},\\\"backendScene\\\":\\\"Mock\\\"}"},
					"model_category": "LLM",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"http_api_name":     name,
						"protocols.#":       "1",
						"type":              "Rest",
						"description":       "test description for cspec",
						"resource_group_id": CHECKSET,
						"ai_protocols.#":    "1",
						"base_path":         "/cspec-test",
						"enable_auth":       "false",
						"deploy_configs.#":  "1",
						"model_category":    "LLM",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "updated description for cspec",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": "updated description for cspec",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"protocols": []string{
						"HTTPS"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"protocols.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"protocols": []string{
						"HTTP", "HTTPS"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"protocols.#": "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"protocols": []string{
						"HTTP"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"protocols.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"deploy_configs": []string{
						"{\\\"autoDeploy\\\":false,\\\"mock\\\":{\\\"enable\\\":false,\\\"responseContent\\\":\\\"updated mock response\\\",\\\"statusCode\\\":404},\\\"backendScene\\\":\\\"Mock\\\"}"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"deploy_configs.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"deploy_configs": []string{
						"{\\\"autoDeploy\\\":true,\\\"mock\\\":{\\\"enable\\\":true,\\\"responseContent\\\":\\\"second mock update\\\",\\\"statusCode\\\":500},\\\"backendScene\\\":\\\"Mock\\\"}"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"deploy_configs.#": "1",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"ai_protocols", "base_path", "deploy_configs", "enable_auth", "model_category", "resource_group_id"},
			},
		},
	})
}

var AlicloudApigHttpApiMap12898 = map[string]string{}

func AlicloudApigHttpApiBasicDependence12898(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

data "alicloud_resource_manager_resource_groups" "default" {}

`, name)
}

// Test Apig HttpApi. <<< Resource test cases, automatically generated.
