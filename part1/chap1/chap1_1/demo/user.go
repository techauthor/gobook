package demo

import "time"

// 用户
type User struct {
	Id int64 //`xorm:"pk"`
	//用户名
	UserName string
	//密码
	Pwd string
	//邮箱
	Email string
	//手机号
	Tel string
	//状态
	Status   string
	CreateAt time.Time `json:"omitempty" xorm:"created"`
	UpdateAt time.Time `xorm:"updated"`
	//删除标志
	DeletedAt time.Time `xorm:"deleted"`
}
