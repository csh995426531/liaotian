package entity

/**
好友实体
*/
type Friend struct {
	Id      int64 `json:"id"`
	UserIdA int64 `json:"user_id_a"`
	UserIdB int64 `json:"user_id_b"`
}

type FriendInterface interface {
	// 查询好友列表
	GetFriendList(userId int64) (list []*Friend, err error)
	// 创建好友信息
	CreateFriendInfo(userIdA, userIdB int64) (friend *Friend, err error)
	// 删除好友信息
	DeleteFriendInfo(id int64) (ok bool, err error)
}
