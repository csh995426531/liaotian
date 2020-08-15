package config

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-plugins/config/source/configmap/v2"
)

var (
	Config 			*config.Config
	MysqlConfig		struct{
		Url			string
	}
)

func Init() {
	configMapSource := configmap.NewSource(
		// optionally specify a namespace; default to default
		configmap.WithNamespace("liaotian"),
		// optionally specify name for ConfigMap; defaults micro
		configmap.WithName("user-service-cm"),
		// optionally strip the provided path to a kube config file mostly used outside of a cluster, defaults to "" for in cluster support.
		//configmap.WithConfigPath(os.Getenv("CONFIG_PATH")),
	)
	// Create new config
	Config, err := config.NewConfig()
	if err != nil {
		logger.Errorf("【config】初始化NewConfig失败，错误：%s", err)
	}

	// Load file source
	err = Config.Load(configMapSource)
	if err != nil {
		logger.Errorf("【config】初始化Load失败，错误：%s", err)
	}

	MysqlConfig.Url = Config.Get("mysql_url").StringMap(map[string]string{"url":"localhost"})["url"]
}