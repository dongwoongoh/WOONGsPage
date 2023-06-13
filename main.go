package main

import (
	"fmt"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	fmt.Println("w", w)
	fmt.Println("r", r)
	fmt.Println("query:", query)
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":6001", nil)
}
