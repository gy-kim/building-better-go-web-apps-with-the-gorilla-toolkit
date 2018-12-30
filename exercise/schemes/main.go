package main

import (
	"fmt"
	"net/http"

	"github.com/gy-kim/mux"
)

func main() {
	r := mux.NewRouter()

	func1 := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("function 1"))
	}

	func2 := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("function 2"))
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.TLS == nil {
			func1(w, r)
		} else {
			func2(w, r)
		}
	}

	r.HandleFunc("/", handler)

	http.Handle("/", r)

	go http.ListenAndServe(":4000", nil)
	go http.ListenAndServeTLS(":4443", "../cert.pem", "../key.pem", nil)

	fmt.Scanln()
}
