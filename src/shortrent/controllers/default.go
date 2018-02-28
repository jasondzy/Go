package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	c.Ctx.Redirect(302, "/template/index.html") //这里将主页路由/重定向到/template/目录下去
}
