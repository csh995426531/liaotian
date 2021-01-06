package entity

import (
	"errors"
	"liaotian/domain/friend/repository"
	"time"
)

/**
申请单实体-仓库实现
*/
type ApplicationModel struct {
	Application `gorm:"embedded"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (ApplicationModel) TableName() string {
	return "users"
}

func (a *Application) CreateApplicationInfo(senderId, receiverId int64) (application *Application, err error) {

	if senderId == 0 || receiverId == 0 {
		err = errors.New("缺少必要参数")
		return
	}

	model := new(ApplicationModel)
	model.SenderId = senderId
	model.ReceiverId = receiverId
	model.Status = StatusWait

	if err = repository.Repo.MysqlDb.Create(model).Error; err != nil {
		return
	}

	application = &Application{
		Id:         model.Id,
		SenderId:   model.SenderId,
		ReceiverId: model.ReceiverId,
		Status:     model.Status,
	}

	return
}

func (a *Application) GetApplicationInfo(id int64) (application *Application, err error) {
	if id == 0 {
		err = errors.New("缺少必要参数")
		return
	}

	model := new(ApplicationModel)
	model.Id = id

	err = repository.Repo.MysqlDb.Where(model).Limit(1).Find(&application).Error
	if err == nil {
		model := new(SayModel)
		model.SenderId = id
		err = repository.Repo.MysqlDb.Where(model).Find(&application.SayList).Error
	}

	return
}

func (a *Application) UpdateApplicationInfoStatus(id int64, status int) (ok bool, err error) {
	if id == 0 {
		err = errors.New("缺少必要参数")
		return
	}
	if !checkStatus(status) {
		err = errors.New("非法参数")
		return
	}

	model := new(ApplicationModel)
	model.Id = id

	updated := repository.Repo.MysqlDb.Model(model).Update("status", status)
	if err = updated.Error; err != nil {
		return
	}

	ok = true
	return
}

func (a *Application) GetApplicationList(userId int64) (list []*Application, err error) {
	if userId == 0 {
		err = errors.New("缺少必要参数")
		return
	}

	where1 := new(ApplicationModel)
	where1.SenderId = userId
	where2 := new(ApplicationModel)
	where2.ReceiverId = userId
	err = repository.Repo.MysqlDb.Where(where1).Or(where2).Find(&list).Error

	return
}
