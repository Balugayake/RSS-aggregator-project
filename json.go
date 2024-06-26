package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func repsonseWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		fmt.Println("error 5XX level", msg)
	}
	type Response struct {
		Error string `json:"error"`
	}

	responseWithJson(w, code, Response{Error: msg})
}

func responseWithJson(w http.ResponseWriter, code int, payload interface{}) {

	data, err := json.Marshal(payload)

	if err != nil {
		log.Printf("fail the marshal repsonse %v", payload)
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
