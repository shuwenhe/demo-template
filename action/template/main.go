package main

import (
	"net/http"
	"text/template"
)

func testTemplate(w http.ResponseWriter, r *http.Request) { // test range
	t := template.Must(template.ParseFiles("template.html", "template2.html"))
	t.Execute(w, "我能在两个文件中显示吗?")
}

func main() {
	http.HandleFunc("/testTemplate", testTemplate)

	http.ListenAndServe(":8080", nil)
}
