package routers

import (
	"quickstart/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() { //init函数是go语言中在main函数之前运行的函数，为go默认调用的函数，所以这个函数不需要手动进行调用
    beego.Router("/", &controllers.MainController{}) //设置路由的路径，这里将路由的路径设置成/
    beego.Router("/user", &controllers.MainController{})
    beego.Router("/api/?:id", &controllers.MainController{}) //这里设置了参数匹配方式 匹配/api/123,其中 :id=123,具体参数的获取可见路由处理函数中的处理
    //###############如下的方式是一种比较简单的注册路由方式######
    beego.Get("/test", func(ctx *context.Context){
    	ctx.Output.Body([]byte("hello world"))
    	})
}
