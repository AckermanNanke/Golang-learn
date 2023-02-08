package controller

import (
	"encoding/json"
	"gorm-demo/dao"
	"gorm-demo/global"
	"gorm-demo/model/freedb"
	"gorm-demo/model/request"
	reqCommon "gorm-demo/model/request/common"
	"io"
	"net/http"

	"go.uber.org/zap"
)

type UserApi struct{}

func (u UserApi) Register(w http.ResponseWriter, r *http.Request) {
	var reqData request.RegisterReq

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		w.Write([]byte("ReadAllerror"))
		global.GVA_LOG.Error("没有body体", zap.Error(err))
		return
	}
	err = json.Unmarshal(body, &reqData)
	if err != nil {
		w.Write([]byte("Unmarshalerror"))
		global.GVA_LOG.Error("JSOn转换失败", zap.Error(err))
		return
	}

	global.GVA_LOG.Info("请求数据", zap.Any("Header", r.Header))
	global.GVA_LOG.Info("请求数据", zap.Any("Method", r.Method))
	global.GVA_LOG.Info("请求数据", zap.Any("PostForm", reqData))

	user := &freedb.FreeUsers{UserName: reqData.UserName, Phone: reqData.Phone, Password: reqData.Password}
	err = dao.UserDaoApi.Register(*user)
	if err != nil {
		w.Write([]byte("error"))
		return
	}

	rspBody := &request.FreeUserRsp{
		reqCommon.ReqCommon{
			Status:  200,
			Message: "注册成功！",
		},
		"actoken",
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Authorization", "4654654sdf54sf6ew4fef")
	rsp, err := json.Marshal(rspBody)
	if err != nil {
		w.Write([]byte("error2222"))
	}
	w.Write([]byte(string(rsp)))
}
