// Type ServeMux Handle

// func (mux *ServeMux) Handle(pattern string, handler Handler)

// mux := http.NewServeMux()
// mux.Handle("/v1/api/ping", http.HandlerFunc(Ping))

package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	// our function
	pingHandler := func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("\nDevops BH for Golang mux Handle()!"))
	}

	// handlerFunc
	mux.Handle("/v1/api/ping", http.HandlerFunc(pingHandler)) // ok
	// mux.Handle("/v1/api/ping2", pingHandler) // error
	// mux.Handle("/v1/api/ping", mux.HandlerFunc(pingHandler)) // error

	// mux.Handle("/", func(w http.ResponseWriter, r *http.Request) {  // error
	// 	w.WriteHeader(http.StatusNotFound)
	// 	fmt.Fprintln(w, "You're lost, go home devopsBH!")
	// })

	log.Printf("\nServer run 8080\n")
	// Listen
	log.Fatal(http.ListenAndServe(":8080", mux))
}
