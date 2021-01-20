package entity

/**
用户实体
*/

type User struct {
	Id       int64
	Name     string
	Account  string
	Password string
	Avatar   string
}

type UserInterface interface {
	// 创建用户信息
	CreateUserInfo(name, account, password, avatar string) (user *User, err error)
	// 获取用户信息
	GetUserInfo(id int64, name, account string) (user *User, err error)
	// 更新用户信息
	UpdateUserInfo(id int64, name, password, avatar string) (user *User, err error)
	// 批量获取用户信息
	BatchGetUserInfo(ids []int64) (list []*User, err error)
}
