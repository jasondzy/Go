package controllers

import (
	"github.com/astaxie/beego"
	"shortrent/models"
	"fmt"
	"encoding/json"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type MyinfoHandler struct {
	beego.Controller
}

type LogoutHandler struct {
	beego.Controller
}

type RealnameHandler struct {
	beego.Controller
}

func (c *MyinfoHandler) Get() {
	// 如下是设置 系统的session
	//该处的session的初始化代码在verify_code.go文件中的init函数
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)

	session_data := sess.Get("username")
	if session_data != nil {
		fmt.Println("username has existed ")
		// fmt.Println(session_data)

		o := orm.NewOrm()
		o.Using("userinfo")

		var maps []orm.Params
		var imagepath string
		num, err := o.Raw("select up_mobile,up_avatar from ih_user_profile where up_name=?", session_data).Values(&maps)
		if err == nil && num > 0 {
			// fmt.Println(maps) // slene	
			if maps[0]["up_avatar"] == "" {
				imagepath = "/static/images/landlord01.jpg"
			} else {
				// fmt.Println("avatar",maps[0]["up_avatar"])
				imagepath = maps[0]["up_avatar"].(string)
			}
			c.Data["json"] = map[string]interface{}{"errcode": "0", "data": map[string]interface{}{"name":session_data, "mobile":maps[0]["up_mobile"], "avatar":imagepath}}
		} else {
			fmt.Println("查询不到")
			c.Data["json"] = map[string]interface{}{"errcode": "4101", "errmsg": "can not find data from database "}
		}
	
		} else {
		fmt.Println("username session does not exist")
		c.Data["json"] = map[string]interface{}{"errcode": "4101", "errmsg": "user not login "}

	}
	// 设置 session end
	c.ServeJSON() //这个函数的作用是将上边的data按照json的方式进行传递，详见beego文档的多种格式输出部分

}

func (c *RealnameHandler) Get() {
	// 如下是设置 系统的session
	//该处的session的初始化代码在verify_code.go文件中的init函数
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)

	session_data := sess.Get("username")
	if session_data != nil {
		fmt.Println("username has existed ", session_data)
		o := orm.NewOrm()
		o.Using("userinfo")

		var maps []orm.Params
		num, err := o.Raw("select up_real_name,up_id_card from ih_user_profile where up_name=?", session_data).Values(&maps)
		if err == nil && num > 0 {
			fmt.Println(maps) // slene	
			c.Data["json"] = map[string]interface{}{"errcode": "0", "data": map[string]interface{}{"real_name":maps[0]["up_real_name"], "id_card":maps[0]["up_id_card"]}}
		} else {
			fmt.Println("查询不到")
			c.Data["json"] = map[string]interface{}{"errcode": "4101", "errmsg": "can not find data from database "}
		}
		
	} else {
		fmt.Println("username session does not exist")
		c.Data["json"] = map[string]interface{}{"errcode": "4101", "errmsg": "user not login, please login "}
	}
	// 设置 session end
	c.ServeJSON() //这个函数的作用是将上边的data按照json的方式进行传递，详见beego文档的多种格式输出部分

}

func (c *RealnameHandler) Post() {
	// 如下是设置 系统的session
	//该处的session的初始化代码在verify_code.go文件中的init函数
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)

	session_data := sess.Get("username")
	if session_data != nil {
		fmt.Println("username has existed ", session_data)

		//获取用户的输入信息 start
		var real_name models.Real_name
		json.Unmarshal(c.Ctx.Input.RequestBody, &real_name)
		fmt.Println("login data=====", real_name)
		//获取用户输入信息 end
		o := orm.NewOrm()
		o.Using("userinfo")

		res, err := o.Raw("update ih_user_profile set up_real_name=?, up_id_card=? where up_mobile=?", real_name.Real_name, real_name.Id_card, session_data).Exec()
		if err == nil {
			num, _ := res.RowsAffected()
			fmt.Println("mysql row affected nums: ", num)
			c.Data["json"] = map[string]interface{}{"errcode": "0", "errmsg": "ok "}
		} else {
			fmt.Println("查询不到")
			c.Data["json"] = map[string]interface{}{"errcode": "1", "errmsg": "update database wrong "}
		}
		
	} else {
		fmt.Println("username session does not exist")
		c.Data["json"] = map[string]interface{}{"errcode": "4101", "errmsg": "user not login, please login "}
	}
	// 设置 session end
	c.ServeJSON() //这个函数的作用是将上边的data按照json的方式进行传递，详见beego文档的多种格式输出部分
}

func (c *LogoutHandler) Get() {

	// 如下是设置 系统的session
	//该处的session的初始化代码在verify_code.go文件中的init函数
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)

	session_data := sess.Get("username")
	if session_data != nil {
		fmt.Println("username has existed , now, set it again ")
		sess.Delete("username")
	} else {
		fmt.Println("username session does not exist")
	}
	// 设置 session end

	c.Data["json"] = map[string]interface{}{"errcode": "0", "errmsg": "ok "}

	c.ServeJSON() //这个函数的作用是将上边的data按照json的方式进行传递，详见beego文档的多种格式输出部分


}