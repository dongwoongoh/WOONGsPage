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

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/bar", barHandler)
	mux.HandleFunc("/", rootHandler)
	http.ListenAndServe(":6000", mux)
}
