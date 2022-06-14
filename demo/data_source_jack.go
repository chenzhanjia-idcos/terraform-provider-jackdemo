package demo

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
)

func dataSourceJackdemoEcs() *schema.Resource {
	return &schema.Resource{
		ReadWithoutTimeout: dataSourceJackdemoEcsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"instance_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tags": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceJackdemoEcsRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conf := meta.(*Configuration)
	endpoint := conf.Endpoint

	name := data.Get("name").(string)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, endpoint+"/data_source?name="+name, nil)
	if err != nil {
		return diag.Errorf("failed to create request, err: %v", err)
	}

	reqDump, _ := httputil.DumpRequestOut(req, true)
	fmt.Printf("REQUEST:\n%s", string(reqDump))

	res, err := client.Do(req)
	if err != nil {
		return diag.Errorf("failed to send request, err: %v", err)
	}
	respDump, _ := httputil.DumpResponse(res, true)
	fmt.Printf("RESPONSE:\n%s", string(respDump))

	resBody, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return diag.Errorf("failed to read response body, err: %v", readErr)
	}

	defer res.Body.Close()

	// 创建成功后，设置新ID
	type ReadResp struct {
		Name         string `json:"name"`
		InstanceType string `json:"instance_type"`
		Tags         string `json:"tags"`
		Id           string `json:"id"`
	}
	var resp ReadResp

	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		return diag.Errorf("failed to unmarshal response body, err: %v", err)
	}

	data.SetId(resp.Id)
	data.Set("name", resp.Name)
	data.Set("instance_type", resp.InstanceType)
	data.Set("tags", resp.Tags)
	return nil

}
