package main

import (
	"html/template"
	"net/http"
)

func testDefine(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("define.html"))
	t.ExecuteTemplate(w, "model", "")
}

func testDefine2(w http.ResponseWriter, r *http.Request) {
	age := 17
	var t *template.Template
	if age < 18 {
		t = template.Must(template.ParseFiles("define2.html", "content2.html"))
	} else {
		t = template.Must(template.ParseFiles("define2.html", "content1.html"))
	}
	t.ExecuteTemplate(w, "model", "")
}

func main() {
	http.HandleFunc("/testDefine", testDefine)
	http.HandleFunc("/testDefine2", testDefine2)

	http.ListenAndServe(":8080", nil)
}
