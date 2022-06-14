package demo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
)

func resourceJackdemoJack() *schema.Resource {
	return &schema.Resource{
		CreateWithoutTimeout: resourceJackdemoJackCreate,
		ReadWithoutTimeout:   resourceJackdemoJackRead,
		UpdateWithoutTimeout: resourceJackdemoJackUpdate,
		DeleteWithoutTimeout: resourceJackdemoJackDelete,

		Schema: map[string]*schema.Schema{
			"instance_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "实例名称",
			},
			"disk_size": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "磁盘大小",
			},
			"tags": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "实例标签",
			},
		},
	}
}

func resourceJackdemoJackCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	body := &map[string]interface{}{
		"instance_name": data.Get("instance_name"),
		"disk_size":     data.Get("disk_size"),
		"tags":          data.Get("tags"),
	}

	bodyEncode, err := json.Marshal(body)
	if err != nil {
		return diag.Errorf("failed to marshal message body, err: %v", err)
	}

	conf := meta.(*Configuration)
	endpoint := conf.Endpoint

	client := &http.Client{}
	buf := bytes.NewBuffer(bodyEncode)
	req, err := http.NewRequest(http.MethodPost, endpoint+"/create", buf)
	if err != nil {
		return diag.Errorf("failed to create request, err: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

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
	type CreateResp struct {
		Id string `json:"id"`
	}
	var resp CreateResp

	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		return diag.Errorf("failed to unmarshal response body, err: %v", err)
	}

	data.SetId(resp.Id)

	return nil
}

func resourceJackdemoJackRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conf := meta.(*Configuration)
	endpoint := conf.Endpoint

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, endpoint+"/get?id="+data.Id(), nil)
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
		InstanceName string `json:"instance_name"`
		DiskSize     string `json:"disk_size"`
		Tags         string `json:"tags"`
		Id           string `json:"id"`
	}
	var resp ReadResp

	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		return diag.Errorf("failed to unmarshal response body, err: %v", err)
	}

	data.Set("instance_name", resp.InstanceName)
	data.Set("disk_size", resp.DiskSize)
	data.Set("tags", resp.Tags)
	return nil
}

func resourceJackdemoJackUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	body := map[string]interface{}{}

	if data.HasChange("instance_name") {
		body["instance_name"] = data.Get("instance_name")
	}
	if data.HasChange("disk_size") {
		body["disk_size"] = data.Get("disk_size")
	}
	if data.HasChange("tags") {
		body["tags"] = data.Get("tags")
	}

	bodyEncode, err := json.Marshal(body)
	if err != nil {
		return diag.Errorf("failed to marshal message body, err: %v", err)
	}

	conf := meta.(*Configuration)
	endpoint := conf.Endpoint

	client := &http.Client{}
	buf := bytes.NewBuffer(bodyEncode)
	req, err := http.NewRequest(http.MethodPut, endpoint+"/update?id="+data.Id(), buf)
	if err != nil {
		return diag.Errorf("failed to create request, err: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	reqDump, _ := httputil.DumpRequestOut(req, true)
	fmt.Printf("REQUEST:\n%s", string(reqDump))

	res, err := client.Do(req)
	if err != nil {
		return diag.Errorf("failed to send request, err: %v", err)
	}
	respDump, _ := httputil.DumpResponse(res, true)
	fmt.Printf("RESPONSE:\n%s", string(respDump))

	return resourceJackdemoJackRead(ctx, data, meta)
}

func resourceJackdemoJackDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	conf := meta.(*Configuration)
	endpoint := conf.Endpoint

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodDelete, endpoint+"/delete?id="+data.Id(), nil)
	if err != nil {
		return diag.Errorf("failed to create request, err: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	reqDump, _ := httputil.DumpRequestOut(req, true)
	fmt.Printf("REQUEST:\n%s", string(reqDump))

	res, err := client.Do(req)
	if err != nil {
		return diag.Errorf("failed to send request, err: %v", err)
	}
	respDump, _ := httputil.DumpResponse(res, true)
	fmt.Printf("RESPONSE:\n%s", string(respDump))

	return nil
}
