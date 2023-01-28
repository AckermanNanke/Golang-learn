package initialize

import (
	"net/http"
	"time"
)

func InitServer(address string) *http.Server {
	return &http.Server{
		Addr:           address,
		Handler:        initRouter(),
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
