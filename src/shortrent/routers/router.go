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
    //处理首页房屋显示的详细信息
    beego.Router("/api/house/info", &controllers.ShowDetail{})
    //处理用户登出操作
    beego.Router("/api/logout/", &controllers.LogoutHandler{})
    //处理用户的详细信息显示
    beego.Router("/api/profile/", &controllers.MyinfoHandler{})
    //处理用户真实信息提交
    beego.Router("/api/profile/auth/", &controllers.RealnameHandler{})
    //处理用户头像上传
    beego.Router("/api/profile/avatar/", &controllers.PersonimgUpload{})
    //处理用户修改用户名
    beego.Router("/api/profile/name/", &controllers.PersonnameModify{})
    //订单显示
    beego.Router("/api/order/my", &controllers.ShowOrder{})
    //订单处理
    beego.Router("/api/order/", &controllers.OrderHandle{})
    //处理搜索信息中的地区显示
    beego.Router("/api/house/area/", &controllers.SerachareaShow{})
    //处理搜索信息要显示的信息
    beego.Router("/api/house/list2", &controllers.SearchinfoShow{})

}

