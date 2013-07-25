// hello project main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func outputFile(w http.ResponseWriter, fileName string) {
	fin, err := os.Open(fileName)
	defer fin.Close()
	if err != nil {
		fmt.Println(fileName, err)
		return
	}
	buf := make([]byte, 1024)
	for {
		n, _ := fin.Read(buf)
		if 0 == n {
			break
		}
		w.Write(buf[:n])
	}
}

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
	if r.URL.Path == "/favicon.ico" {
		outputFile(w, "static/favicon.ico")
	} else {
		http.Redirect(w, r, "/login/", 302)
	}
}

func main() {
	http.Handle("/static/", http.FileServer(http.Dir("")))
	http.HandleFunc("/login/", viewLogin)
	http.HandleFunc("/ajax/", viewAjax)
	http.HandleFunc("/runstat/", viewRunstat)
	http.HandleFunc("/", viewRoot)
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
