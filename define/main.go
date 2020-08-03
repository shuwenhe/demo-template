package main

import (
	"html/template"
	"net/http"
)

func testWith(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("with.html"))
	t.Execute(w, "狸猫")
}

func testTemplate(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("template1.html", "template2.html"))
	t.Execute(w, "我能在两个文件中显示吗？")
}

func main() {
	http.HandleFunc("/testWith", testWith)
	http.HandleFunc("/testTemplate", testTemplate)
	http.ListenAndServe(":8080", nil)
}
