package rabbitmq

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccVhost_importBasic(t *testing.T) {
	resourceName := "rabbitmq_vhost.test"
	var vhost string

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccVhostCheckDestroy(vhost),
		Steps: []resource.TestStep{
			{
				Config: testAccVhostConfig_basic,
				Check: testAccVhostCheck(
					resourceName, &vhost,
				),
			},

			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
