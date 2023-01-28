package initialize

import (
	"gorm-demo/router"
	"net/http"
)

var mux = http.NewServeMux()

func initRouter() *http.ServeMux {
	mux.HandleFunc("/register", router.Register)
	return mux
}
