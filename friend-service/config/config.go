package config

import (
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/debug/log"
	"github.com/micro/go-plugins/config/source/configmap"
)

var (
	Config      *config.Config
	MysqlConfig struct {
		Url string
	}
	SkywalkingConfig struct {
		Url string
	}
)

func Init() {
	configMapSource := configmap.NewSource(
		configmap.WithNamespace("liaotian"),
		configmap.WithName("friend-service-cm"),
	)

	Config := config.NewConfig()

	err := Config.Load(configMapSource)
	if err != nil {
		log.Errorf("【config】初始化Load失败，错误：%s", err)
	}
	MysqlConfig.Url = Config.Get("mysql_url").StringMap(map[string]string{"url": "localhost"})["url"]
	SkywalkingConfig.Url = Config.Get("skywalking_url").StringMap(map[string]string{"url": "localhost"})["url"]
}
