package controller

import "net/http"

func Ping(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Ping Pong"))
}