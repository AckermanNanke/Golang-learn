package request

type ReqCommon struct {
	Actoken string `json:"actoken" example:"用户唯一ID"`
	ApiID   string `json:"apiID" example:"api编号"`
}
