package entity

/**
朋友实体
*/
type Friend struct {
	Id      int64
	UserIdA int64
	UserIdB int64
}

type FriendInterface interface {
	// 查询好友列表
	GetFriendList(userId int64) (list []*Friend, err error)
	// 创建好友信息
	CreateFriendInfo(userIdA, userIdB int64) (friend *Friend, err error)
	// 删除好友信息
	DeleteFriendInfo(id int64) (ok bool, err error)
}
