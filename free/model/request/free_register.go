package request

type RegisterReq struct {
	UserName string `json:"userName" example:"用户登录名"`
	Phone    int    `json:"phone" example:"用户手机号"`
	Password string `json:"password" example:"用户密码"`
}

type RegisterRsp struct {
	RspCommon
	AcToken string `json:"actoken" gorm:"comment:用户登陆凭证"`
}
