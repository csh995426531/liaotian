package entity

import (
	"errors"
	"liaotian/domain/friend/repository"
	"time"
)

type FriendModel struct {
	Friend `gorm:"embedded"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (FriendModel) TableName() string {
	return "friend"
}

func (f *Friend) GetFriendList(userId int64) (list []*Friend, err error) {
	if userId == 0 {
		err = errors.New("缺少必要参数")
		return
	}

	where1 := new(FriendModel)
	where1.UserIdA = userId
	where2 := new(FriendModel)
	where2.UserIdB = userId
	err = repository.Repo.MysqlDb.Where(where1).Or(where2).Find(&list).Error
	return
}

func (f *Friend) CreateFriendInfo(userIdA, userIdB int64) (friend *Friend, err error) {
	if userIdA == 0 || userIdB == 0 {
		err = errors.New("缺少必要参数")
		return
	}

	model := new(FriendModel)
	model.UserIdA = userIdA
	model.UserIdB = userIdB

	if err = repository.Repo.MysqlDb.Create(model).Error; err != nil {
		return
	}
	friend = &Friend{
		Id: model.Id,
		UserIdA: userIdA,
		UserIdB: userIdB,
	}
	return
}

func (f *Friend) DeleteFriendInfo(id int64) (ok bool, err error) {
	if id == 0 {
		err = errors.New("缺少必要参数")
		return
	}

	model := new(FriendModel)
	model.Id = id
	if err = repository.Repo.MysqlDb.Delete(&model).Error; err != nil {
		return
	}
	ok = true
	return
}
