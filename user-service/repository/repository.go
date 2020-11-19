package repository

import (
	"liaotian/user-service/config"
	"log"
	"os"
	"sync"
	"time"

	microlog "github.com/micro/go-micro/util/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	m    sync.Mutex
	repo *Repository
)

type Interface interface {
	Create(name, password string) (user *ModelUser, err error)
	Get(name, password string, id int64) (user *ModelUser, err error)
}

type Repository struct {
	mysqlDB *gorm.DB
}

func Init() *Repository {
	m.Lock()
	defer m.Unlock()

	if repo != nil {
		microlog.Fatal("repo 已经初始化过了")
		return repo
	}

	repo = &Repository{
		mysqlDB: newDb(),
	}

	return repo
}

/**
创建新db链接实例
*/
func newDb() *gorm.DB {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Silent, // Log level
			Colorful:      false,         // 禁用彩色打印
		},
	)

	mysqlDb, err := gorm.Open(mysql.New(mysql.Config{
		DSN: config.MysqlConfig.Url + "?charset=utf8&parseTime=true&loc=Local", // Data Source Name，参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name
	}), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		microlog.Error(err)
		panic(err)
	}
	sqlDB, err := mysqlDb.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	return mysqlDb
}
