package repository

type ModelFriend struct {
	ID       int64  `gorm:"column:id;primary_key;auto_increment;"`
	OperatorId int64	`gorm:"column:operator_id"`
	BuddyId		int64	`gorm:"column:buddy_id"`
}

func (ModelFriend) TableName() string {
	return "friends"
}

func (r *Repository) Add (operatorId, buddyId int64) (friend *ModelFriend, err error) {

	friend = new(ModelFriend)
	friend.OperatorId = operatorId
	friend.BuddyId = buddyId
	res := r.mysqlDB.Create(friend)
	err = res.Error
	if res.RowsAffected == 1 {
		friend = new(ModelFriend)
		friend.OperatorId = buddyId
		friend.BuddyId = operatorId
		res := r.mysqlDB.Create(friend)
		err = res.Error
	}
	return
}

func (r *Repository) Del (friendId int64) (err error) {

	err = r.mysqlDB.Where("id = ?", friendId).Delete(&ModelFriend{}).Error
	return
}

func (r *Repository) List (operatorId, offset, limit int64) (friends []*ModelFriend, err error) {

	err = r.mysqlDB.Where("operator_id = ?", operatorId).Offset(offset).Limit(limit).Find(&friends).Error
	return
}

func (r *Repository) Get (friendId int64) (friend *ModelFriend, err error) {

	friend = new(ModelFriend)
	err = r.mysqlDB.Where("id = ?", friendId).First(&friend).Error
	return
}
