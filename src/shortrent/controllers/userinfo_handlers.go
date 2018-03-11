package controllers

import (
	"github.com/astaxie/beego"
	"shortrent/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
    orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:dzy@/test?charset=utf8")
	orm.RegisterDataBase("userinfo", "mysql", "root:dzy@/Ihome?charset=utf8")
	//如下是一个测试数据表
    // orm.RegisterDataBase("test", "mysql", "root:dzy@/test?charset=utf8")
}

type UserinfoHandler struct {
	beego.Controller
}

func (c *UserinfoHandler) Post() {
	var userinfo models.Register_data

	json.Unmarshal(c.Ctx.Input.RequestBody, &userinfo)
	fmt.Println(userinfo)

	if userinfo.Password != userinfo.Password2 {
		fmt.Println("密码不匹配")
		c.Data["json"] = map[string]interface{}{"errcode": "1", "errmsg": "password mismatch"}
		// this.Ctx.WriteString(rs)
		c.ServeJSON() //这个函数的作用是将上边的data按照json的方式进行传递，详见beego文档的多种格式输出部分
		return

	} else {
/***********如下实现的是对一个数据库的验证操作，实现了一条数据插入到test表中
	o := orm.NewOrm()
	o.Using("test") // 默认使用 default，你可以指定为其他数据库
	
	test_data := new(models.Test_tbl)
	test_data.Test_title = "test111"
	test_data.Test_author = "dzy"

	fmt.Println(o.Insert(test_data))
************************************************************/

// 如下进行 用户信息的注册，将用户的信息注册到数据库中去
	o := orm.NewOrm()
	o.Using("userinfo")

	// 这里用来判断注册的手机号是否已经注册过 start

	var maps []orm.Params
	num, err := o.Raw("SELECT * FROM ih_user_profile WHERE up_mobile = ?", userinfo.Mobile).Values(&maps)
	if err == nil && num > 0 {

		c.Data["json"] = map[string]interface{}{"errcode": "1", "errmsg": "mobile number existed"}
		// this.Ctx.WriteString(rs)
		c.ServeJSON() //这个函数的作用是将上边的data按照json的方式进行传递，详见beego文档的多种格式输出部分
		return 
	}

	// 如下是设置 系统的session
	//该处的session的初始化代码在verify_code.go文件中的init函数
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)

	session_data := sess.Get("username")
	if session_data != nil {
		fmt.Println("username has existed , now, set it again ")
		sess.Delete("username")
		sess.Set("username", userinfo.Mobile)
	} else {
		fmt.Println("set username session")
		sess.Set("username", userinfo.Mobile)
	}
	// 设置 session end


	// 判断手机号是否注册 end

	userinfo_data := new(models.Ih_user_profile)
	userinfo_data.Up_name = userinfo.Mobile
	userinfo_data.Up_mobile = userinfo.Mobile
	userinfo_data.Up_passwd = userinfo.Password

	fmt.Println(o.Insert(userinfo_data))

	c.Data["json"] = map[string]interface{}{"errcode": "0", "errmsg": "ok"}
	// this.Ctx.WriteString(rs)
	c.ServeJSON() //这个函数的作用是将上边的data按照json的方式进行传递，详见beego文档的多种格式输出部分


}

}