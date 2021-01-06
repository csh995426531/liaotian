package entity

/**
申请单实体
*/

type Application struct {
	Id         int64
	SenderId   int64
	ReceiverId int64
	Status     int
	SayList    []Say
}

var (
	StatusWait   = 1
	StatusPass   = 2
	StatusReject = 3
)

type ApplicationInterface interface {
	// 创建申请单信息
	CreateApplicationInfo(senderId, receiverId int64) (application *Application, err error)
	// 查询申请单信息
	GetApplicationInfo(id int64) (application *Application, err error)
	// 更新申请单状态
	UpdateApplicationInfoStatus(id int64, status int) (ok bool, err error)
	// 查询申请单列表
	GetApplicationList(userId int64) (list []*Application, err error)
}

func checkStatus(status int) bool {
	if status != StatusWait && status != StatusPass && status != StatusReject {
		return false
	}
	return true
}
