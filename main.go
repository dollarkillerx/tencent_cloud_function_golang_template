package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/dollarkillerx/urllib"
	"github.com/tencentyun/scf-go-lib/cloudfunction"
)

// Request api 请求默认会解析JSON
type Request struct {
	Url             string                 `json:"url"`              // 地址
	Title           string                 `json:"title"`            // title
	Data            []byte                 `json:"data"`             // 数据源
	ApplicationType string                 `json:"application_type"` // 类型
	DownloadTime    int                    `json:"download_time"`    // 下载日期
	ReleaseTime     int                    `json:"release_time"`     // 发布日期
	Other           map[string]interface{} `json:"other"`
}

//// Response api 默认会解析JSON返回
//type Response struct {
//	TaskID string                            `json:"task_id"`
//	Data   map[string]map[string]interface{} `json:"data"`
//}

type DefineEvent struct {
	// test event define
	Body string `json:"body"` // 所有post 参数 都会被  scf-go-lib json 解析放到Body这里
}

func spider(ctx context.Context, event DefineEvent) (*Request, error) { // API 返回结构体会默认解析为 JSON
	var req Request

	err := json.Unmarshal([]byte(event.Body), &req)
	if err != nil {
		return nil, err
	}

	code, respBytes, err := urllib.Get(req.Url).RandUserAgent().RandDisguisedIP().ByteRetry(3, 200)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if code != 200 {
		return nil, errors.New(string(respBytes))
	}

	req.Data = respBytes

	return &req, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by Cloud Function
	cloudfunction.Start(spider)
}
