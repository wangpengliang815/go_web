package common

import "time"

type User struct {
	// gorm.Model        // gorm里自定义的结构包含ID/CreatedAt/UpdatedAt/DeletedAt四个字段,其中DeletedAt标记当前数据是否被删除(gorm中默认使用逻辑删除)如果使用的话对应的Sql语句会自动添加相关筛选语句
	Id          int       `gorm:"primaryKey; column:Id"`
	Name        string    `gorm:"default:'test'; column:Name; type:nvarchar(64);not null;comment:用户姓名"`
	Age         int       `gorm:"column:Age"`
	PhoneNumber string    `gorm:"column:PhoneNumber"`
	Address     string    `gorm:"column:Address"`
	CreateTime  time.Time `gorm:"column:CreateTime"`
}

// TableName 重命名表名User否则为小写的user+复数s也就是users
func (User) TableName() string {
	return "User"
}
