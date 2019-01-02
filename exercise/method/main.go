package main

import (
	"fmt"
	"net/http"

	"github.com/gy-kim/mux"
)

func main() {
	r := mux.NewRouter()

	html := "<html><head></head><body>" +
		"<form action='' method='POST'>" +
		" The field <input name='field'/>" +
		"</form></body></html>"

	func1 := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(html))
	}

	func2 := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(r.FormValue("field")))
	}

	r.HandleFunc("/", func1).Methods("GET")
	r.HandleFunc("/", func2).Methods("POST")

	http.Handle("/", r)

	go http.ListenAndServe(":4000", nil)

	fmt.Scanln()
}
