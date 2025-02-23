package main

// Func http Handlefunc
// func HandleFunc(pattern string, handler func(ResponseWriter, *Request))

// http.HandleFunc("/v1/api/ping", pingHandler)
// http.HandleFunc("/v1/api/ping", func(w http.ResponseWriter, req *http.Request){})

import (
	"log"
	"net/http"
)

func main() {

	// our function
	pingHandler := func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("\nDevops BH for Golang HandleFunc!"))
	}

	// handleFunc
	http.HandleFunc("/v1/api/ping", pingHandler)
	http.HandleFunc("/v1/api/ping2", pingHandler)
	http.HandleFunc("/v1/api/ping3", pingHandler)

	// show run server
	log.Printf("\nServer run 8080\n")

	// Listen
	log.Fatal(http.ListenAndServe(":8080", nil))
}
