package dao

import (
	"errors"
	"gorm-demo/global"
	"gorm-demo/model/freedb"

	"github.com/google/uuid"
)

type UserDao struct{}

// 用户注册，判断数据库中是否已有数据，有则返回错误信息，没有则创建用户数据
func (ud *UserDao) Register(u freedb.FreeUsers) error {
	var user freedb.FreeUsers
	err := global.MY_SQL.Where("phone = ? ", u.Phone).Find(&user).Error
	if err != nil {
		return errors.New("数据已存在")
	}
	u.UUID = uuid.New()
	err = global.MY_SQL.Create(&u).Error
	if err != nil {
		return errors.New("创建数据失败")
	}
	return err
}
