package router

import (
	"gorm-demo/dao"
	"net/http"
	"time"
)

func Register(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	isSuccess := dao.Register(&dao.User{UserID: 1, Name: "Tom", PhoneNumber: 15500000001, Password: "qweqweqwe"})
	if isSuccess == nil {
		w.Write([]byte("The time is: " + tm))
	} else {
		w.Write([]byte("Error" + tm))
	}
}
