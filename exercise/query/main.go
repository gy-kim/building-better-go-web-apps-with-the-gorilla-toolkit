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

	r.Path("/").Queries("foo", "bar").HandlerFunc(func1)          // http://localhost:4000/?foo=bar
	r.Path("/").Queries("bar", "{bar:[0-9]+}").HandlerFunc(func2) // http://localhost:4000/?bar=12

	http.Handle("/", r)

	go http.ListenAndServe(":4000", nil)

	fmt.Scanln()

}
