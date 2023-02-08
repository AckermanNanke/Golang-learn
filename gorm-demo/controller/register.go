package controller

import (
	"encoding/json"
	"gorm-demo/global"
	"gorm-demo/model/request"
	"net/http"

	"go.uber.org/zap"
)

type UserApi struct{}

func (u UserApi) Register(w http.ResponseWriter, r *http.Request) {
	var reqPamra request.RegisterReq
	// reqBody, err := io.ReadAll(r.Body)
	// defer r.Body.Close()
	// if err != nil {
	// 	w.Write([]byte("ReadAllerror"))
	// 	global.GVA_LOG.Error("没有body体", zap.Error(err))
	// 	return
	// }
	// err = json.Unmarshal(reqBody, postPrarm)
	// if err != nil {
	// 	w.Write([]byte("Unmarshalerror"))
	// 	global.GVA_LOG.Error("JSOn转换失败", zap.Error(err))
	// 	return
	// }
	r.ParseForm()
	for i := range reqPamra {
		global.GVA_LOG.Info("请求数据", zap.Any("PostForm", i))

	}
	if r.Method == "POST" {
	}
	if r.Method == "GET" {
	}
	global.GVA_LOG.Info("请求数据", zap.Any("PostForm", r.Header))
	global.GVA_LOG.Info("请求数据", zap.Any("PostForm", r.Method))
	global.GVA_LOG.Info("请求数据", zap.Any("PostForm", reqPamra))

	// userDao := &dao.UserDao{}
	// user := &freedb.FreeUsers{UUID: 2323, Username: r.Form.Get("username"), Phone: r.Form.Get("phone"), Password: r.Form.Get("password")}
	// err := userDao.Register(*user)
	// if err != nil {
	// 	w.Write([]byte("error"))
	// 	return
	// }

	rspBody := &request.FreeUserRsp{
		Status: 32,
	}
	w.Header().Set("Content-Type", "application/json")
	rspData, err := json.Marshal(rspBody)
	if err != nil {
		w.Write([]byte("error2222"))
	}
	w.Write([]byte(string(rspData)))
}
