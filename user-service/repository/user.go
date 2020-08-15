package repository

//Model model default
//type Model struct {

	//CreateAt *time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP"`
	//UpdateAt *time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP"`
//}

//ModelUser .
type ModelUser struct {
	//gorm.Model
	ID       int64  `gorm:"column:id;primary_key;auto_increment;"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
}

// UserInfo 用户信息
type UserInfo struct {
	id uint
	username string
	password string
}



func (ModelUser) TableName() string {
	return "users"
}

func (r *Repository) Create (name, password string) (user *ModelUser, err error) {
	user = &ModelUser{Username: name, Password: password}
	r.mysqlDB.Create(user)
	return
}

func (r *Repository) Get (name, password string) (user *ModelUser, err error) {

	user = new(ModelUser)
	r.mysqlDB.Where("username = ? and password = ?", name, password).Last(user)
	return
}
