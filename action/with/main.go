package main

import (
	"net/http"
	"text/template"
)

func testWith(w http.ResponseWriter, r *http.Request) { // test range
	t := template.Must(template.ParseFiles("with.html"))
	t.Execute(w, "狸猫")
}

func main() {
	http.HandleFunc("/testWith", testWith)

	http.ListenAndServe(":8080", nil)
}
