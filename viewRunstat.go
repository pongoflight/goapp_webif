// viewRunstat.go
package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type RunstatPage struct {
	Page
	Fields []string
}

func viewRunstat(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("runstat method:", r.Method, r.URL.Path, r.Form) //获取请求的方法
	page := new(RunstatPage)
	page.Initialize("运行状态")
	page.Fields = []string{"Running", "Uptime"}
	if r.Method == "GET" {
		t, _ := template.ParseFiles("template/runstat.html")
		t.Execute(w, page)
	} else if r.Method == "POST" {

	}
}
