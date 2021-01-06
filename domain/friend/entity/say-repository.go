package entity

import (
	"errors"
	"liaotian/domain/friend/repository"
	"time"
)

type SayModel struct {
	Say       `gorm:"embedded"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (SayModel) TableName() string {
	return "application_say"
}

func (s *Say) CreateApplicationSay(applicationId, senderId int64, content string) (say *Say, err error) {

	if applicationId == 0 || senderId == 0 || content == "" {
		err = errors.New("缺少必要参数")
		return
	}

	model := new(SayModel)
	model.ApplicationId = applicationId
	model.SenderId = senderId
	model.Content = content

	if err = repository.Repo.MysqlDb.Create(model).Error; err != nil {
		return
	}
	say = &Say{
		Id:            model.Id,
		ApplicationId: applicationId,
		SenderId:      senderId,
		Content:       content,
	}

	return
}
