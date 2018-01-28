package main 

import (
	_ "quickstart/routers"
	"github.com/astaxie/beego"
)

func main() { //这里是这个gobee的入口函数，main包中的mian函数是go语言的入口
	//#################### 设置静态文件的路径 ############################
	//################### beego框架的默认静态文件的路径在static文件下下 ，如下是想通过手动的方式添加其他的一些静态文件路径
	//################## 注意，如下的设置方式一定要在 run函数之前
	
	// beego.SetStaticPath("/", "/home/ubuntu/user_jason/temp") //这条语句的作用就是设置静态文件的路径为/home/ubuntu/user_jason/temp，当浏览器请求为http://xxx/index.html
	
	//#################################################################
	//这里的Run是在beego.go的函数中被调用，该run函数实现了run函数中的参数解析，和一些运行前的初始化操作
	beego.Run() //这里调用beego中的RUN函数，
}

