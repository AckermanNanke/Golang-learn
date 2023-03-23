package initialize

import (
	"gorm-demo/core"
	"net/http"
	"time"
)

func InitServer(address string, handler *core.Fhttp) *http.Server {
	return &http.Server{
		Addr:           address,
		Handler:        handler,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
