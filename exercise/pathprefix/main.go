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

	r.PathPrefix("/foo").HandlerFunc(func1)
	r.PathPrefix("/bar").HandlerFunc(func2)

	http.Handle("/", r)

	go http.ListenAndServe(":4000", nil)

	fmt.Scanln()
}
