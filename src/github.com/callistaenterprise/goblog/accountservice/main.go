package main

import (
	"fmt"
	"github.com/callistaenterprise/goblog/accountservice/service"
	"github.com/callistaenterprise/goblog/accountservice/dbclient"
)

var appName = "accountservice"

func main() {
	fmt.Printf("starting %v\n", appName)
	//初始化database
	initializeBoltClient()

	service.StartWebServer("6767")
}


func initializeBoltClient() {
	service.DBClient = &dbclient.BoltClient{} //这里定义一个空的db结构
	service.DBClient.OpenBoltDb() // 调用上边定义的db结构中的Open方法来创建一个database文件
	service.DBClient.Seed() //调用db结构体中的seed方法，往创建的database文件中添加数据
}
