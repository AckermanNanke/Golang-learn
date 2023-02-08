package freedb

import (
	"gorm-demo/global"

	"github.com/google/uuid"
)

type FreeUsers struct {
	global.GVA_MODEL
	UUID     uuid.UUID `json:"uuid" gorm:"comment:用户UUID"`
	UserName string    `json:"userName" gorm:"index;comment:用户登录名"`
	Phone    int       `json:"phone" gorm:"comment:用户手机号"`
	Password string    `json:"-" gorm:"comment:用户密码"`
}

// 获取表名
func (FreeUsers) TableName() string {
	return "free_users"
}
