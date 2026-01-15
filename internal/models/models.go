package models

// User 测试用的用户模型
type User struct {
	ID    int64  `gorm:"primaryKey" xorm:"pk autoincr 'id'" json:"id" zorm:"id,auto_incr" borm:"id"`
	Name  string `gorm:"column:name" xorm:"varchar(100) 'name'" json:"name" zorm:"name" borm:"name"`
	Email string `gorm:"column:email" xorm:"varchar(100) 'email'" json:"email" zorm:"email" borm:"email"`
	Age   int    `gorm:"column:age" xorm:"int 'age'" json:"age" zorm:"age" borm:"age"`
}

// TableName 表名
func (User) TableName() string {
	return "users"
}
