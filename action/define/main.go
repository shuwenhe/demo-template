package main

import (
	"net/http"
	"text/template"
)

func testDefine(w http.ResponseWriter, r *http.Request) { // test range
	t := template.Must(template.ParseFiles("define.html"))
	t.ExecuteTemplate(w, "model", "")
}

func main() {
	http.HandleFunc("/testDefine", testDefine)

	http.ListenAndServe(":8080", nil)
}
