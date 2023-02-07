package controller

import (
	"gorm-demo/dao"
	"net/http"
	"time"
)

type UserApi struct{}

func (UserApi) Register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	reqBody := r.PostForm
	tm := time.Now().Format(time.RFC1123)
	isSuccess := dao.Register(reqBody)
	if isSuccess == nil {
		w.Write([]byte("The time is: " + tm))
	} else {
		w.Write([]byte("Error" + tm))
	}
}
