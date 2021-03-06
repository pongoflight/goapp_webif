// hello project main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func viewRoot(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	http.Redirect(w, r, "/runstat/", 302)
}

var gStatOptMgr *OptionMgr
var gParaOptMgr *OptionMgr

func main() {
	// 启动状态option读写服务，及参数option读写服务
	gStatOptMgr = InitlizeStatOptMgr()
	gParaOptMgr = InitlizeParaOptMgr()

	http.Handle("/static/", http.FileServer(http.Dir("")))
	http.HandleFunc("/login/", viewLogin)
	http.HandleFunc("/ajax/", viewAjax)
	http.HandleFunc("/runstat/", viewRunstat)
	http.HandleFunc("/parameter/", viewParameter)
	http.HandleFunc("/", viewRoot)
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
