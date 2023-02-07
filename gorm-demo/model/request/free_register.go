package request

import "gorm-demo/global"

type FreeUserReq struct {
	global.GVA_MODEL
	Username string `json:"userName" gorm:"index;comment:用户登录名"`
	Phone    int    `json:"phone" gorm:"index;comment:用户手机号"`
	Password string `json:"-" gorm:"comment:用户密码"`
}

type FreeUserRsp struct {
	Status int `json:"status"`
}
