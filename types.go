// types.go
package main

type NaviItem struct {
	Uri   string
	Title string
}
type Page struct {
	AppTitle string
	Navi     [4]NaviItem
}

// 页面公共信息
var publicPage = Page{
	AppTitle: "XHK-II-X(61850)装置参数配置工具",
	Navi: [4]NaviItem{
		NaviItem{Uri: "runstat", Title: "运行状态"},
		NaviItem{Uri: "parameter", Title: "参数配置"},
		NaviItem{Uri: "loginfo", Title: "日志信息"},
		NaviItem{Uri: "logout", Title: "退出登录"},
	},
}
