package main

import (
	"demo-template/range/model"
	"html/template"
	"net/http"
)

func testRange(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("range.html"))
	var emps []*model.Employee
	emp := &model.Employee{
		ID:       1,
		LastName: "李小路",
		Email:    "lxl@qq.com",
	}
	emps = append(emps, emp)

	emp2 := &model.Employee{
		ID:       2,
		LastName: "白百何",
		Email:    "bbh@qq.com",
	}
	emps = append(emps, emp2)

	emp3 := &model.Employee{
		ID:       3,
		LastName: "马蓉",
		Email:    "mr@wbq.com",
	}
	emps = append(emps, emp3)

	t.Execute(w, emps)
}

func main() {
	http.HandleFunc("/testRange", testRange)
	http.ListenAndServe(":8080", nil)
}
