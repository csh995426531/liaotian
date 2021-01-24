package validator

type SendRequest struct {
	FriendId   int64  `json:"id" validate:"required,min=1"`
	SenderId   int64  `json:"user_id" validate:"required,min=1"`
	ReceiverId int64  `json:"receiver_id" validate:"required,min=1"`
	Content    string `json:"content" validate:"required,min=1"`
}

type ConnRequest struct {
	UserId int64 `form:"user_id" validate:"required,min=1"`
}
