package controllers

import (
	"github.com/astaxie/beego"
)

/**************这里定义MainController作为路由的处理函数****************************/
type MainController struct {
	beego.Controller //这里采用组合的方式包含了beego中的Controller结构
}

func (c *MainController) Get() { //这里对beego中的Controller结构进行重新定义
	c.Data["Website"] = "beego.me" //这里定义的Data是要传入如下定义的模板中
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl" //这里定义所用的模板
	// c.Ctx.WriteString("hello world") //这里的方式可以不使用上边的模板方式，可以直接使用该方式往客户端写入一个字符串，该方式和tornado比较相似
}

/***************可在如下参考上方的方式定义路由的处理函数****************************/