package validator

//好友列表验证器
type FriendListValidator struct {
	UserId int64 `validate:"required,min=1"`
}

//好友信息验证器
type FriendInfoValidator struct {
	Id int64 `validate:"required,min=1"`
}

//删除好友验证器
type DeleteFriendInfoValidator struct {
	Id int64 `validate:"required,min=1"`
}