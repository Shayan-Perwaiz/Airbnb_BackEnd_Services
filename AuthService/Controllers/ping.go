package controller

import "net/http"

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ping Pong"))
}