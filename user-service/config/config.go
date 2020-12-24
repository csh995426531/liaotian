package config

import (
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/util/log"
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
		// optionally specify a namespace; default to default
		configmap.WithNamespace("liaotian"),
		// optionally specify name for ConfigMap; defaults micro
		configmap.WithName("user-handler-cm"),
		// optionally strip the provided path to a kube config file mostly used outside of a cluster, defaults to "" for in cluster support.
		//configmap.WithConfigPath(os.Getenv("CONFIG_PATH")),
	)
	// Create new config
	Config := config.NewConfig()

	// Load file source
	err := Config.Load(configMapSource)
	if err != nil {
		log.Errorf("【config】初始化Load失败，错误：%s", err)
	}

	MysqlConfig.Url = Config.Get("mysql_url").StringMap(map[string]string{"url": "localhost"})["url"]
	SkywalkingConfig.Url = Config.Get("skywalking_url").StringMap(map[string]string{"url": "localhost"})["url"]
}
