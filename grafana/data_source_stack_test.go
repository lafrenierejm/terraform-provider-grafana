package grafana

import (
	"fmt"
	"testing"

	gapi "github.com/grafana/grafana-api-golang-client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceStack_Basic(t *testing.T) {
	CheckCloudTestsEnabled(t)

	resourceName := GetRandomStackName()
	var stack gapi.Stack
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheckCloudStack(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccStackCheckDestroy(&stack),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceStackConfig(resourceName),
				Check: resource.ComposeTestCheckFunc(
					testAccStackCheckExists("grafana_cloud_stack.test", &stack),
					resource.TestCheckResourceAttrSet("data.grafana_cloud_stack.test", "id"),
					resource.TestCheckResourceAttr("data.grafana_cloud_stack.test", "name", resourceName),
					resource.TestCheckResourceAttr("data.grafana_cloud_stack.test", "slug", resourceName),
					resource.TestCheckResourceAttrSet("data.grafana_cloud_stack.test", "prometheus_url"),
					resource.TestCheckResourceAttrSet("data.grafana_cloud_stack.test", "prometheus_user_id"),
					resource.TestCheckResourceAttrSet("data.grafana_cloud_stack.test", "alertmanager_user_id"),
				),
			},
		},
	})
}

func testAccDataSourceStackConfig(resourceName string) string {
	return fmt.Sprintf(`
resource "grafana_cloud_stack" "test" {
  name = "%s"
  slug = "%s"
  region_slug = "eu"
}
data "grafana_cloud_stack" "test" {
  slug = grafana_cloud_stack.test.slug
  depends_on = [grafana_cloud_stack.test]
}
`, resourceName, resourceName)
}