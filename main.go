package main

import (
	"fmt"
	"net/http"
)

func BarHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "bar")
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "root")
}

func WebHandler() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/bar", BarHandler)
	mux.HandleFunc("/", RootHandler)

	return mux
}

func main() {
	http.ListenAndServe(":6001", WebHandler())
}
