package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"liaotian/middlewares/logger/zap"
	"sync"
)

/**
仓库
*/

var (
	m    sync.Mutex
	Repo *Repository
)

type Repository struct {
	MysqlDb *gorm.DB
	MockDb  sqlmock.Sqlmock
}

func Init(db *gorm.DB, mock sqlmock.Sqlmock) {
	m.Lock()
	defer m.Unlock()

	if Repo != nil {
		zap.ZapLogger.Error("仓库已经初始化过了")
		return
	}

	Repo = &Repository{
		MysqlDb: db,
		MockDb:  mock,
	}
}

func NewDb() *gorm.DB {
	dsn := "debian-sys-maint:F0sm3f7WrNJox1oV@(129.211.55.205:3306)/liaotian"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.SugarLogger.Panicf("仓库实例化DB失败，error: %v", err)
	}
	return db
}

func NewMockDb() (db *gorm.DB, mock sqlmock.Sqlmock) {
	mockDb, mock, err := sqlmock.New()
	if err != nil {
		zap.SugarLogger.Panicf("仓库实例化mock-DB失败，error: %v", err)
	}

	db, err = gorm.Open(mysql.New(mysql.Config{
		Conn:                      mockDb,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	if err != nil {
		zap.SugarLogger.Panicf("仓库实例化DB失败，error: %v", err)
	}
	return
}
