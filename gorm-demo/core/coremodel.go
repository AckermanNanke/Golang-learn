package core

import "net/http"

type Context struct {
	r *http.Request
	w http.ResponseWriter
}
