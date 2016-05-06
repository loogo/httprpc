package main

import (
	"github.com/loogo/httprpc/jsonrpc"
	"fmt"
)

func main() {
	params := map[string]interface{}{
		"service": "common",
		"method": "version",

	}
	url := "http://demo.odoo.com/start"
	result, err := jsonrpc.Call(url, params)
	if err == nil {
		fmt.Println(result)
	}
}