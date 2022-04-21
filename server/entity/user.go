package entity

import "time"

type User struct {
	Id          int       `gorm:"primaryKey; column:Id"`
	UserName    string    `gorm:"default:'test'; column:UserName; type:nvarchar(100);not null;comment:用户姓名"`
	PassWord    string    `gorm:"column:PassWord; type:nvarchar(100);not null;comment:用户密码"`
	Age         int       `gorm:"column:Age"`
	PhoneNumber string    `gorm:"column:PhoneNumber"`
	Address     string    `gorm:"column:Address"`
	CreateTime  time.Time `gorm:"column:CreateTime"`
}

// TableName 重命名表名User否则为小写的user+复数s也就是users
func (User) TableName() string {
	return "User"
}
