package config

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-plugins/config/source/configmap/v2"
)

var (
	Config *config.Config
	MysqlConfig struct {
		Url string
	}
)

func Init() {
	configMapSource := configmap.NewSource(
		configmap.WithNamespace("liaotian"),
		configmap.WithName("friend-service-cm"),
	)

	Config, _ := config.NewConfig()

	err := Config.Load(configMapSource)
	if err != nil {
		logger.Errorf("【config】初始化Load失败，错误：%s", err)
	}
	MysqlConfig.Url = Config.Get("mysql_url").StringMap(map[string]string{"url":"localhost"})["url"]
}