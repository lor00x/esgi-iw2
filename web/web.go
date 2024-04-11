package main

import (
	"fmt"
	"net/http"
)

func Log(next http.Handler) http.Handler {

	logger := func(w http.ResponseWriter,
		r *http.Request) {
		fmt.Printf("Request : %s %s\n",
			r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(logger)
}
func main() {

	s := http.NewServeMux()

	// Conversion en http.HandlerFunc
	hello := http.HandlerFunc(helloWorld)
	s.Handle("/hello", Log(hello))

	http.ListenAndServe(":8080", s)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}
