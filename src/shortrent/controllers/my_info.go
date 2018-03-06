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

type PersonimgUpload struct {
	beego.Controller
}

type PersonnameModify struct {
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
				fmt.Println("imagepath",imagepath)
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

//***********************修改用户头像信息********************************************************
func (c *PersonimgUpload) Post() {

	// 如下是设置 系统的session
	//该处的session的初始化代码在verify_code.go文件中的init函数
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)

	session_data := sess.Get("username")
	if session_data != nil {
		fmt.Println("username has existed  ")
		
		//从数据库中获取数据
		o := orm.NewOrm()
		o.Using("userinfo")

		var maps []orm.Params
		num, err := o.Raw("select up_mobile,up_avatar from ih_user_profile where up_name=?", session_data).Values(&maps)
		if err == nil && num > 0 {
			fmt.Println(maps) // slene	

			//这里处理用户上传的图片源信息，使用的是beego的上传文件处理方法
			f, _, err := c.GetFile("avatar")//该函数是获取上传信息，保存在系统内存中，此处主要用来判定用户的上传文件是否正确
			if err != nil {
				fmt.Println(" getfile error")
			}
			defer f.Close()
			file_path := "/static/images/personal_images/" + maps[0]["up_mobile"].(string)
			c.SaveToFile("avatar", file_path) // 保存位置在 static/images/personal_images, 没有文件夹要先创建

			//如下将上传的图片路径保存到数据库中去
			res, err := o.Raw("update ih_user_profile set up_avatar=? where up_name=?", file_path, session_data).Exec()
			if err == nil {
				num, _ := res.RowsAffected()
				fmt.Println("mysql row affected nums: ", num)
				c.Data["json"] = map[string]interface{}{"errcode": "0", "data": file_path}
				fmt.Println(file_path)
			} else {
				fmt.Println("查询不到")
				c.Data["json"] = map[string]interface{}{"errcode": "4101", "errmsg": "update database wrong "}
			}


		} else {
			fmt.Println("查询不到")
			c.Data["json"] = map[string]interface{}{"errcode": "4101", "errmsg": "can not find data from database "}
		}


	} else {
		fmt.Println("username session does not exist")
		c.Data["json"] = map[string]interface{}{"errcode": "4101", "errmsg": "not login "}
	}
	// 设置 session end

	c.ServeJSON() //这个函数的作用是将上边的data按照json的方式进行传递，详见beego文档的多种格式输出部分
}
//**********************************end*************************************************************

//**********************************修改用户名功能***************************************************
func (c *PersonnameModify) Post() {
	//获取用户的输入信息 start
	var modify_name models.Name_modify
	json.Unmarshal(c.Ctx.Input.RequestBody, &modify_name)
	fmt.Println("name to change =====", modify_name)
	//获取用户输入信息 end

	// 如下是设置 系统的session
	//该处的session的初始化代码在verify_code.go文件中的init函数
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)

	session_data := sess.Get("username")
	if session_data != nil {
		fmt.Println("username has existed ")
		o := orm.NewOrm()
		o.Using("userinfo")

		var maps []orm.Params
		num, err := o.Raw("select * from ih_user_profile where up_name=?", modify_name.Name).Values(&maps)
		if err == nil && num > 0 {
			fmt.Println("查询到数据：", maps) // slene	
			c.Data["json"] = map[string]interface{}{"errcode": "4001", "errmsg": "change name already exist "}
		
		} else {
			fmt.Println("查询不到")
			//如下将上传的图片路径保存到数据库中去
			res, err := o.Raw("update ih_user_profile set up_name=? where up_name=?", modify_name.Name, session_data).Exec()
			if err == nil {
				num, _ := res.RowsAffected()
				fmt.Println("mysql row affected nums: ", num)
				c.Data["json"] = map[string]interface{}{"errcode": "0", "errmsg": "ok"}
			} else {
				fmt.Println("查询不到")
				c.Data["json"] = map[string]interface{}{"errcode": "4101", "errmsg": "update database wrong "}
			}
		}
		
	} else {
		fmt.Println("username session does not exist")
		c.Data["json"] = map[string]interface{}{"errcode": "4101", "errmsg": "not login "}
	}
	// 设置 session end

	c.ServeJSON() //这个函数的作用是将上边的data按照json的方式进行传递，详见beego文档的多种格式输出部分
}
//**********************************end*************************************************************


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