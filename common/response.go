package common

import "net/http"

func SendError(w http.ResponseWriter, err error, statusCode int) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(statusCode)
	w.Write([]byte(err.Error()))
	return
}