package routers

import (
	"shortrent/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    //这里定义了验证码图片的获取
    beego.Router("/api/piccode/", &controllers.YzmController{}, "Get:GetYzm")
    //这里定义了验证码的校验操作，通过该路由对提交的验证码进行校验
    beego.Router("/api/smscode/", &controllers.YzmController{}, "Post:JudgeYzm")
    //这里定义一个处理用户提交注册信息的功能
    beego.Router("/api/register/", &controllers.UserinfoHandler{})
    //这里定义一个处理用户登录的功能
    beego.Router("/api/login/", &controllers.LoginHandler{})
    //主页上用来处理检查用户是否登录的功能
    beego.Router("/api/check_login/", &controllers.IndexHandler{})
    //处理的是首页的房屋信息显示
    beego.Router("/api/house/index/", &controllers.HouseinfoHandler{})
}

