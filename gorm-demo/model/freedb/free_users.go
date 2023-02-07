package freedb

import "gorm-demo/global"

type FreeUsers struct {
	global.GVA_MODEL
	UUID     int    `json:"uuid" gorm:"index;comment:用户UUID"`
	Username string `json:"userName" gorm:"index;comment:用户登录名"`
	Phone    int    `json:"phone" gorm:"index;comment:用户手机号"`
	Password string `json:"-" gorm:"comment:用户密码"`
}

//获取表名
func (FreeUsers) TableName() string {
	return "free_users"
}
