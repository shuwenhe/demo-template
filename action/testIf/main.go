package main

import (
	"html/template"
	"net/http"
)

func testIf(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("if.html")) // Parse the template
	age := 17
	t.Execute(w, age > 18) // The second parameter is the action returned to the template
}

func main() {
	http.HandleFunc("/testIf", testIf)

	http.ListenAndServe(":8080", nil)
}
