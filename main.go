package main

import (
	"context"
	"fmt"

	"github.com/tencentyun/scf-go-lib/cloudfunction"
)

// DefineEvent api 请求默认会解析JSON
type DefineEvent struct {
	// test event define
	Key1 string `json:"key1"`
	Key2 string `json:"key2"`
}

func hello(ctx context.Context, event DefineEvent) (*DefineEvent, error) { // API 返回结构体会默认解析为 JSON
	fmt.Println("key1:", event.Key1)
	fmt.Println("key2:", event.Key2)

	return &event, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by Cloud Function
	cloudfunction.Start(hello)
}
