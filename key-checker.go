/*
	kong-go-key校验插件
*/

package main

import (
	"fmt"
	"github.com/Kong/go-pdk"
	"log"
	"time"
)

// Config 插件配置文件
type Config struct {
	Apikey string
}

// New一个插件
func New() interface{} {
	return &Config{}
}

const version = "v0.1.1"

// Access kong 入口
func (conf Config) Access(kong *pdk.PDK) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("kong plugin panic at: %v, err: %v", time.Now(), err)
			if kong == nil {
				log.Printf("kong fatal err ===> kong is nil at: %v", time.Now())
			} else {
				_ = kong.Log.Err(fmt.Sprint(err))
			}
		}
	}()

	//检查配置
	if err := conf.checkConfig(); err != nil {
		_ = kong.Log.Err("[checkConfig] ", err.Error())
		return
	}

	_ = kong.Response.SetHeader("X-Key-Checker-Plugin-Version", version)

	// 路由key校验
	key, err := kong.Request.GetQueryArg("key")
	apiKey := conf.Apikey
	//apiKey := "apikey"
	if err != nil {
		kong.Log.Err(err.Error())
	}
	x := make(map[string][]string)
	x["Content-Type"] = append(x["Content-Type"], "application/json")
	if apiKey != key {
		kong.Response.Exit(403, "Youu have no correct key", x)
	}
}

//进入此插件，说明kong中已经启用插件
func (conf Config) checkConfig() error {

	return nil
}
