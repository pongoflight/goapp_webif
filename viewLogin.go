// viewLogin.go
package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type LoginPage struct {
	Page
	PageTitle string
	Failed    bool
}

func viewLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("login method:", r.Method, r.URL.Path, r.Form) //获取请求的方法
	if r.Method == "GET" {
		page := LoginPage{Page: publicPage, PageTitle: "用户登录", Failed: false}
		t, _ := template.ParseFiles("template/login.html")
		t.Execute(w, page)
	} else if r.Method == "POST" {
		var tm = time.Now()
		var auth_pw1 = fmt.Sprintf("%02d%02d", tm.Month(), tm.Day()+1)
		var auth_pw2 = fmt.Sprintf("%02d%02d", tm.Month(), tm.Day()+2)
		form_pw, ok := r.Form["inputPassword"]
		if ok {
			if form_pw[0] == auth_pw1 {
				http.Redirect(w, r, "/runstat/", 302)
			} else if form_pw[0] == auth_pw2 {
				w.Write([]byte("S2"))
			} else {
				page := LoginPage{Page: publicPage, Failed: true}
				t, _ := template.ParseFiles("template/login.html")
				t.Execute(w, page)
			}
		} else {
			page := LoginPage{Page: publicPage, Failed: true}
			t, _ := template.ParseFiles("template/login.html")
			t.Execute(w, page)
		}
	}
}
