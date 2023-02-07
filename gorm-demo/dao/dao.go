package dao

import (
	"errors"
	"gorm-demo/global"
	"gorm-demo/model/request"
)

var user request.FreeUserReq

// 用户注册，判断数据库中是否已有数据，有则返回错误信息，没有则创建用户数据
func Register(u *request.FreeUserReq) error {
	err := global.MY_SQL.Where("phone_number = ? ", u.Phone).Find(&user).Error
	if err != nil {
		return errors.New("数据已存在")
	}
	global.MY_SQL.Create(&u)
	return err
}
