package controller

import (
	"encoding/json"
	"gorm-demo/model/request"
	"net/http"
)

type UserApi struct{}

func (u UserApi) Register(w http.ResponseWriter, r *http.Request) {
	// var postPrarm *request.FreeUserReq
	// reqBody, err := io.ReadAll(r.Body)
	// defer r.Body.Close()
	// if err != nil {
	// 	w.Write([]byte("error"))
	// 	return
	// }
	// err = json.Unmarshal(reqBody, postPrarm)

	// if err != nil {
	// 	w.Write([]byte("error"))
	// }
	// fmt.Fprintf(w, "%#v\n", postPrarm)
	// isSuccess := dao.Register(postPrarm)
	// if isSuccess != nil {
	// 	w.Write([]byte("error"))
	// }

	rspBody := &request.FreeUserRsp{
		Status: 32,
	}
	w.Header().Set("Content-Type", "application/json")
	rspData, err := json.Marshal(rspBody)
	// json.NewEncoder(w).Encode(&rspBody)
	if err != nil {
		w.Write([]byte("error2222"))
	}
	w.Write([]byte(string(rspData)))
}
