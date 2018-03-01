package routers

import (
	"shortrent/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/api/piccode/", &controllers.YzmController{}, "Get:GetYzm")
    beego.Router("/api/smscode/", &controllers.YzmController{}, "Post:JudgeYzm")
}
