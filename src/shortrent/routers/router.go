package routers

import (
	"shortrent/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/api/piccode/", &controllers.YzmController{}, "Get:GetYzm")
    beego.Router("/yzm/judgeyzm", &controllers.YzmController{}, "Post:JudgeYzm")
}
