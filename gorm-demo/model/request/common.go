package request

type ReqCommon struct {
	ApiID   string `json:"apiID" example:"api编号"`
	Message string `json:"message" example:"返回信息"`
	Status  int    `json:"status" example:"接口状态"`
}

type RspCommon struct {
	ApiID   string `json:"apiID" example:"api编号"`
	Message string `json:"message" example:"返回信息"`
	Status  int    `json:"status" example:"接口状态"`
}

// var ResFailInfo []byte = ""
