// viewAjax.go
package main

import (
	"fmt"
	"net/http"
	"time"
)

const layout = "2006-01-02 15:04:05 Mon"

func viewAjax(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("ajax method:", r.Method, r.URL.Path, r.Form) //获取请求的方法
	if r.Method == "GET" {
		if r.Form["act"][0] == "gettime" {
			w.Write([]byte(time.Now().Format(layout)))
		}
		if r.Form["act"][0] == "getrunstat" {
			w.Write([]byte("{\"serverName\":\"SY\"}"))
		}
	} else if r.Method == "POST" {

	}
}
