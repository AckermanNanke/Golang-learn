package controller

import (
	"encoding/json"
	"gorm-demo/dao"
	"gorm-demo/global"
	"gorm-demo/model/freedb"
	"gorm-demo/model/request"
	"io"
	"net/http"

	"go.uber.org/zap"
)

type UserApi struct{}

func (u UserApi) Register(w http.ResponseWriter, r *http.Request) {
	var req request.RegisterReq
	var res *request.RegisterRsp

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		res = &request.RegisterRsp{
			request.RspCommon{
				Status:  201,
				Message: "缺少请求数据！",
			},
			"",
		}
		global.GVA_LOG.Error("获取Body体失败", zap.Error(err))
		return
	}
	err = json.Unmarshal(body, &req)
	if err != nil {
		res = &request.RegisterRsp{
			request.RspCommon{
				Status:  202,
				Message: "请求数据格式错误！",
			},
			"",
		}
		global.GVA_LOG.Error("JSOn转换失败", zap.Error(err))
		return
	}

	global.GVA_LOG.Info("请求数据", zap.Any("Header", r.Header))
	global.GVA_LOG.Info("请求数据", zap.Any("Method", r.Method))
	global.GVA_LOG.Info("请求数据", zap.Any("PostForm", req))

	user := &freedb.FreeUsers{UserName: req.UserName, Phone: req.Phone, Password: req.Password}
	err = dao.UserDaoApi.Register(*user)
	if err != nil {
		res = &request.RegisterRsp{
			request.RspCommon{
				Status:  203,
				Message: "数据写入失败！",
			},
			"",
		}
		return
	}

	res = &request.RegisterRsp{
		request.RspCommon{
			Status:  200,
			Message: "注册成功！",
		},
		"actoken",
	}
	rsp, err := json.Marshal(res)
	if err != nil {
		w.Write([]byte("error2222"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Authorization", "4654654sdf54sf6ew4fef")
	w.Write([]byte(string(rsp)))
}
