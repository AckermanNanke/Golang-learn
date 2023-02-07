package initialize

import (
	"gorm-demo/controller"
	"net/http"
)

var mux = http.NewServeMux()

func initRouter() *http.ServeMux {
	userApi := &controller.UserApi{}
	mux.HandleFunc("/register", userApi.Register)
	return mux
}
