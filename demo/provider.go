package demo

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider 初始化
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{},
		ResourcesMap: map[string]*schema.Resource{
			"jackdemo_jack": resourceJackdemoJack(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"jackdemo_ecs": dataSourceJackdemoEcs(),
		},
		ConfigureContextFunc: providerJackConfigFunc,
	}
}

type Configuration struct {
	Endpoint string `json:"endpoint"`
}

func providerJackConfigFunc(ctx context.Context, data *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return &Configuration{
		Endpoint: "http://127.0.0.1:8888",
	}, nil
}
