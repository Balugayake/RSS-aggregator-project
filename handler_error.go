package main

import "net/http"

func hanlerError(w http.ResponseWriter, r *http.Request) {

	repsonseWithError(w, 400, "something went wrong")
}
