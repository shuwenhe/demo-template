package main

import (
	"html/template"
	"net/http"
)

func testDefine(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("define.html"))
	t.ExecuteTemplate(w, "model", "")
}

func main() {
	http.HandleFunc("/testDefine", testDefine)

	http.ListenAndServe(":8080", nil)
}
