package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8000", nil)
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)
	fmt.Fprintln(w, req.URL.Path)
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}
