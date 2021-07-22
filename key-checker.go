/*
	kong-go-key校验插件
*/

package main

import (
	"github.com/Kong/go-pdk"
)

// Config 插件配置文件
type Config struct {
	Apikey string
}

// New一个插件
func New() interface{} {
	return &Config{}
}

// Access kong 入口
func (conf Config) Access(kong *pdk.PDK) {
	// 路由key校验
	key, err := kong.Request.GetQueryArg("key")
	apiKey := conf.Apikey
	if err != nil {
		kong.Log.Err(err.Error())
	}
	x := make(map[string][]string)
	x["Content-Type"] = append(x["Content-Type"], "application/json")
	if apiKey != key {
		kong.Response.Exit(403, "Youu have no correct key", x)
	}
}
