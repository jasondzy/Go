package service

import (
	"net/http"
	"log"
)

func StartWebServer(port string) {
	log.Println("Starting HTTP service at" + port)

	r := NewRouter() //调用一个第三方的路由库，对标准的net/http的路由功能进行封装
	http.Handle("/", r) //使用http.Handle来注册处理 函数，因为r中实现了serverhttp这个方法，根据鸭子模型，这里可以直接传入r这个参数，其实这里的r是对标准库的handle的一种封装，类似中间件的作用
		  //在匹配了/这个路径后，r中的mux库会继续对剩下的路径进行解析处理
		  //所以这里只需要注册一个根路由/即可，剩下的路由路径在r中通过mux库来注册
		  //如若这里不使用mux库来对路由进行容易管理，那么这里会调用众多的http.Handle("/accounts/xx",xx)函数来注册各个路由
	err := http.ListenAndServe(":" + port, nil)//这里监听并处理请求，一开始通过调用defaultservermux来处理/这个路径，接着用r来处理剩下的路径并调用相应的函数

	if err != nil {
		log.Println("An error occured starting HTTPlistener at port", port)
		log.Println("Error: " + err.Error())
	}
}