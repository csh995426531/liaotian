package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2/logger"
	"liaotian/friend-service/config"
	"sync"
)

var (
	m 	sync.Mutex
	repo	*Repository
)

type Interface interface {
	Add (operatorId, buddyId int64) (friend *ModelFriend, err error)
	Del (friendId int64) (err error)
	List (operatorId, offset, limit int64) (friends []*ModelFriend, err error)
	Get (friendId int64) (friend *ModelFriend, err error)
}

type Repository struct {
	mysqlDB *gorm.DB
}

func Init() Interface {
	m.Lock()
	defer m.Unlock()

	if repo != nil {
		logger.Fatal("repo 已经初始化过了")
		return repo
	}

	repo = &Repository{
		mysqlDB: newDb(),
	}
	return repo
}

func newDb() *gorm.DB {

	logger.Infof("db链接地址：%#+v", config.MysqlConfig.Url)

	mysqlDb, err := gorm.Open("mysql", config.MysqlConfig.Url + "?charset=utf8&parseTime=true")
	if err != nil {
		logger.Error(err)
		panic(err)
	}

	return mysqlDb
}