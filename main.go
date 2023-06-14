package main

import (
	"fmt"
	"net/http"
)

func barHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "bar")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "root")
}

func webHandler() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/bar", barHandler)
	mux.HandleFunc("/", rootHandler)

	return mux
}

func main() {
	http.ListenAndServe(":6001", webHandler())
}
