// viewParameter.go
package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type ParaPage struct {
	Page
	Options []*Option
}

func viewParameter(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("runstat method:", r.Method, r.URL.Path, r.Form) //获取请求的方法
	page := new(ParaPage)
	page.Initialize("参数配置")
	count := gParaOptMgr.GetOptionCount()
	page.Options = make([]*Option, count)
	for i := 0; i < count; i++ {
		page.Options[i] = gParaOptMgr.GetOptionByIndex(i)
	}
	if r.Method == "GET" {
		t, _ := template.ParseFiles("template/parameter.html")
		t.Execute(w, page)
	} else if r.Method == "POST" {

	}
}
