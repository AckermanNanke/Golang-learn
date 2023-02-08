package request

import request "gorm-demo/model/request/common"

type RegisterReq struct {
	UserName string `json:"userName" example:"用户登录名"`
	Phone    int    `json:"phone" example:"用户手机号"`
	Password string `json:"password" example:"用户密码"`
}

type FreeUserRsp struct {
	request.ReqCommon
	AcToken string `json:"actoken" gorm:"comment:用户登陆凭证"`
}
