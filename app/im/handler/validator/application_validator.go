package validator

// 创建申请单验证器
type CreateApplicationValidator struct {
	SenderId   int64 `from:"user_id" json:"user_id" validate:"required,min=1"`
	ReceiverId int64 `validate:"required,min=1"`
}

// 申请单列表验证器
type ApplicationListValidator struct {
	UserId int64 `from:"user_id" json:"user_id" validate:"required,min=1"`
}

// 申请单信息验证器
type ApplicationInfoValidator struct {
	Id int64 `validate:"required,min=1"`
}

// 通过申请验证器
type PassApplicationValidator struct {
	Id int64 `validate:"required,min=1"`
}

// 拒绝申请验证器
type RejectApplicationValidator struct {
	Id int64 `validate:"required,min=1"`
}

// 回复申请验证器
type ReplyApplicationValidator struct {
	Id       int64  `validate:"required,min=1"`
	SenderId int64  `from:"user_id" json:"user_id" validate:"required,min=1"`
	Content  string `validate:"required,min=1"`
}
