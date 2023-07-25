package main

import (
	"TheFirst/config"
	"TheFirst/route"
	"log"
	"net/http"
)

func main() {
	startWebServer()
}

// 通过指定端口启动 Web 服务器
func startWebServer() {
	r := route.NewRouter()

	// 处理静态资源文件
	assets := http.FileServer(http.Dir(config.Viperconfig.App.Static))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", assets))

	http.Handle("/", r)

	log.Println("Starting HTTP service at " + config.Viperconfig.App.Address)
	err := http.ListenAndServe(config.Viperconfig.App.Address, nil)

	if err != nil {
		log.Println("An error occured starting HTTP listener at " + config.Viperconfig.App.Address)
		log.Println("Error: " + err.Error())
	}
}
