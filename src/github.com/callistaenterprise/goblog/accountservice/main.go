package main

import (
	"fmt"
	"flag"
	"github.com/callistaenterprise/goblog/accountservice/service"
	"github.com/callistaenterprise/goblog/accountservice/dbclient"
	"github.com/callistaenterprise/goblog/accountservice/config"
	"github.com/spf13/viper"
	"github.com/callistaenterprise/goblog/accountservice/message"
	"github.com/streadway/amqp"
)

var appName = "accountservice"

func init() {
	profile := flag.String("profile", "test", "Environment profile, something similar to spring profiles")
	configServerUrl := flag.String("configServerUrl", "http://192.168.31.241:8888", "Address to config serv")
	configBranch := flag.String("configBranch", "master", "git branch to fetch configuration from")

	flag.Parse() //对传入的参数进行解析

	viper.Set("profile", *profile)  //使用viper来保存参数
	viper.Set("configServerUrl", *configServerUrl)
	viper.Set("configBranch", *configBranch)


}

func main() {
	fmt.Printf("starting %v\n", appName)

	config.LoadConfigurationFromBranch( //调用 loadconfig函数来从远程加载config参数
		viper.GetString("configServerUrl"),
		appName,
		viper.GetString("profile"),
		viper.GetString("configBranch"))

	//初始化database
	initializeBoltClient()

	initializeMessage()

	service.StartWebServer(viper.GetString("server_port")) //这里从远程github中解析出来的port口号，为7777
	//service.StartWebServer("6767")
}

//这里初始化database，在这里要掌握GO语言的面向对象编程的方法
func initializeBoltClient() {
	service.DBClient = &dbclient.BoltClient{} //这里定义一个空的db结构
	service.DBClient.OpenBoltDb() // 调用上边定义的db结构中的Open方法来创建一个database文件
	service.DBClient.Seed() //调用db结构体中的seed方法，往创建的database文件中添加数据
}


func initializeMessage() {
	if !viper.IsSet("amqp_server_url") {
		viper.Set("amqp_server_url", "amqp://guest:guest@192.168.31.241:5672")
	}

	service.MessagingClient = &message.MessagingClient{}
	service.MessagingClient.ConnectToBroker(viper.GetString("amqp_server_url"))
	service.MessagingClient.SubscribeToQueue("vipQueue", "topic", handleMessage)
}

func handleMessage(delivery amqp.Delivery) {
	fmt.Printf("Got a message: %v\n", string(delivery.Body))
}
