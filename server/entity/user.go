package entity

import "time"

type User struct {
	Id          int       `gorm:"primaryKey; column:Id"`
	UserName    string    `gorm:"column:UserName; type:nvarchar(100);not null;comment:用户姓名"`
	PassWord    string    `gorm:"column:PassWord; type:nvarchar(100);not null;comment:用户密码"`
	Age         int       `gorm:"column:Age"`
	PhoneNumber string    `gorm:"column:PhoneNumber"`
	Address     string    `gorm:"column:Address"`
	Avatar      string    `gorm:"column:Avatar; type:nvarchar(max);"`
	CreateTime  time.Time `gorm:"column:CreateTime"`
}

func (User) TableName() string {
	return "User"
}
