// viewRunstat.go
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type RunstatPage struct {
	Page
	PageTitle string
}

func viewRunstat(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("runstat method:", r.Method, r.URL.Path, r.Form) //获取请求的方法
	if r.Method == "GET" {
		page := RunstatPage{Page: publicPage, PageTitle: "运行状态"}
		t, _ := template.ParseFiles("template/runstat.html")
		t.Execute(w, page)
	} else if r.Method == "POST" {

	}
}
