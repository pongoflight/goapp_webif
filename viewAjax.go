// viewAjax.go
package main

import (
	"encoding/json"
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
			count := gStatOptMgr.GetOptionCount()
			options := make([]*Option, count)
			for i := 0; i < count; i++ {
				options[i] = gStatOptMgr.GetOptionByIndex(i)
				//options[i].Value = time.Now().Format(layout)
			}
			buf, _ := json.Marshal(options)
			w.Write(buf)
		}
		if r.Form["act"][0] == "getparameter" {
			count := gParaOptMgr.GetOptionCount()
			options := make([]*Option, count)
			for i := 0; i < count; i++ {
				options[i] = gParaOptMgr.GetOptionByIndex(i)
				options[i].Value = time.Now().Format(layout)
			}
			buf, _ := json.Marshal(options)
			w.Write(buf)
		}
	} else if r.Method == "POST" {
		if r.Form["act"][0] == "postparameter" {
			option := gParaOptMgr.GetOptionByToken(r.Form["oid"][0])
			option.Value = r.Form["value"][0]
			option = gParaOptMgr.PutOption(option)
			if option.Token == "" {
				w.Write([]byte("fail"))
			} else {
				w.Write([]byte("success"))
			}

		}
	}
}
