package model

import (
	"errors"
	"gorm-demo/initSQL"
)

type User struct {
	ID          int
	Name        string
	PhoneNumber int
	Password    string
}

var user User

func Register(data *User) error {
	err := initSQL.MY_SQL_DB.Where("Phone = ? ", data.PhoneNumber).Find(&user).Error
	if err == nil {
		return errors.New("数据已存在")
	}
	initSQL.MY_SQL_DB.Create(&data)
	return err
}
