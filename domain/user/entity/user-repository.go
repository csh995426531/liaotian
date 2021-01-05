package entity

import (
	"errors"
	"liaotian/domain/user/repository"
	"time"
)

/**
用户实体仓库实现
*/

type UserModel struct {
	User      `gorm:"embedded"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (UserModel) TableName() string {
	return "users"
}

func (e *User) CreateUserInfo(name, account, password, avatar string) (user *User, err error) {

	if name == "" || account == "" || password == "" {
		err = errors.New("缺少必要参数")
		return
	}

	user = &User{
		Name:     name,
		Account:  account,
		Password: password,
		Avatar:   avatar,
	}

	model := new(UserModel)
	model.Name = name
	model.Account = account
	model.Password = password
	model.Avatar = avatar

	err = repository.Repo.MysqlDb.Create(model).Error
	user.Id = model.Id

	return
}

func (e *User) GetUserInfo(id int64, name, account string) (user *User, err error) {

	if id == 0 && name == "" && account == "" {
		err = errors.New("缺少必要参数")
		return
	}

	model := new(UserModel)
	if id > 0 {
		model.Id = id
	}
	if name != "" {
		model.Name = name
	}
	if account != "" {
		model.Account = account
	}

	user = &User{}
	err = repository.Repo.MysqlDb.Where(model).Limit(1).Find(&user).Error
	return
}

func (e *User) UpdateUserInfo(id int64, name, password, avatar string) (user *User, err error) {

	if id == 0 || name == "" || password == "" {
		err = errors.New("缺少必要参数")
		return
	}

	model := new(UserModel)
	model.Id = id

	data := map[string]interface{}{
		"name":     name,
		"password": password,
		"avatar":   avatar,
	}
	column := []string{
		"name",
		"password",
		"avatar",
	}

	err = repository.Repo.MysqlDb.Model(model).Select(column).Updates(data).Error

	user = &User{
		Id:       model.Id,
		Name:     model.Name,
		Account:  model.Account,
		Password: model.Password,
		Avatar:   model.Avatar,
	}
	return
}
