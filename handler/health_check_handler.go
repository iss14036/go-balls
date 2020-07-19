package handler

import "net/http"

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Pong"))
}