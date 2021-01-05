package entity


type Say struct {
	Id            int64
	ApplicationId int64
	SenderId      int64
	Content       string
}

type SayInterface interface {
	// 创建申请单内容
	CreateApplicationSay(applicationId, senderId int64, content string) (say *Say, err error)
}