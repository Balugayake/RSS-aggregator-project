package main

import "net/http"

func handlerRediness(w http.ResponseWriter, r *http.Request) {
	responseWithJson(w, 200, struct{}{})
}
