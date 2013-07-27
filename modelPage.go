// types.go
package main

type NaviItem struct {
	Uri    string
	Title  string
	Active bool
}
type Page struct {
	AppTitle  string
	PageTitle string
	Navi      []NaviItem
}

func (r *Page) Initialize(pageTitle string) {
	r.AppTitle = "XHK-II-X(61850)装置参数配置工具"
	r.PageTitle = pageTitle
	r.Navi = []NaviItem{
		NaviItem{Uri: "/runstat", Title: "运行状态", Active: pageTitle == "运行状态"},
		NaviItem{Uri: "/parameter", Title: "参数配置", Active: pageTitle == "参数配置"},
		NaviItem{Uri: "/loginfo", Title: "日志信息", Active: pageTitle == "日志信息"},
		NaviItem{Uri: "/logout", Title: "退出登录", Active: pageTitle == "退出登录"},
	}
}
