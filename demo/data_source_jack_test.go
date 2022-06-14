package demo

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"testing"
)

var providerFactureis = map[string]func() (*schema.Provider, error){
	"jackdemo": func() (*schema.Provider, error) {
		return Provider(), nil
	},
}

func Test_dataSourceJackdemoEcs(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() {},
		ProviderFactories: providerFactureis,
		Steps: []resource.TestStep{
			{
				Config: testDataSourceJackdemoEcs,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.jackdemo_ecs.test", "name", "ecs"),
					resource.TestCheckResourceAttr("data.jackdemo_ecs.test", "instance_type", "normal"),
				),
			},
		},
	})
}

const testDataSourceJackdemoEcs = `
data "jackdemo_ecs" "test" {
	name = "ecs"
    instance_type = "normal"
}`
