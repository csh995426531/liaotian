package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2/logger"
	"liaotian/user-service/config"
	"sync"

)

var (
	m sync.Mutex
	repo *Repository
)

type Interface interface {
	Create(name, password string) (user *ModelUser, err error)
	Get(name, password string) (user *ModelUser, err error)
}

type Repository struct {
	mysqlDB *gorm.DB
}

func Init () *Repository {
	m.Lock()
	defer m.Unlock()

	if repo != nil {
		logger.Fatal("repo 已经初始化过了")
	}

	repo = &Repository{
		mysqlDB: newDb(),
	}

	return repo
}

/**
	创建新db链接实例
 */
func newDb () *gorm.DB {


	logger.Errorf("%#+v", config.MysqlConfig.Url)
	mysqlDb, err := gorm.Open("mysql", config.MysqlConfig.Url + "?charset=utf8&parseTime=true")
	if err != nil {
		logger.Error(err)
		panic(err)
	}
	mysqlDb.LogMode(true)

	return mysqlDb
}

