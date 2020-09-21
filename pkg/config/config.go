package config

import (
	"filebak/pkg/data"
	"github.com/jinzhu/configor"
	"time"
)

var Config = struct {
	//AppName string `json:"app_name"`
	BaseDir string      `json:"base_dir"`
	Tasks   []data.Task `json:"tasks"`
}{}

func init() {
	err := configor.New(&configor.Config{
		AutoReload:         true,
		AutoReloadInterval: time.Second * 5,
	}).Load(&Config, "config.yaml")
	if err != nil {
		panic(err)
	}
}
