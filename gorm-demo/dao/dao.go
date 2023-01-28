package dao

import (
	"errors"
	"gorm-demo/global"
)

var user User

// 用户注册，判断数据库中是否已有数据，有则返回错误信息，没有则创建用户数据
func Register(data *User) error {
	err := global.MY_SQL.Where("phone_number = ? ", data.PhoneNumber).Find(&user).Error
	if err != nil {
		return errors.New("数据已存在")
	}
	global.MY_SQL.Create(&data)
	return err
}
