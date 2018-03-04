package controllers

import (
	"github.com/astaxie/beego"
	"shortrent/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type LoginHandler struct {
	beego.Controller
}

func (c *LoginHandler) Post() {
	var login_data models.Login_data

	json.Unmarshal(c.Ctx.Input.RequestBody, &login_data)
	fmt.Println(login_data)

	o := orm.NewOrm()
	o.Using("userinfo")

	var maps []orm.Params
	num, err := o.Raw("SELECT * FROM ih_user_profile WHERE up_mobile = ?", login_data.Mobile).Values(&maps)
	if err == nil && num > 0 {
		fmt.Println(maps) // slene
		fmt.Println(maps[0]["up_passwd"])
		if login_data.Password != maps[0]["up_passwd"] {
			fmt.Println("password is wrong")
			c.Data["json"] = map[string]interface{}{"errcode": "1", "errmsg": "password is wrong "}
			
		} else {
		c.Data["json"] = map[string]interface{}{"errcode": "0", "errmsg": "ok"}
		}		
	} else {
		fmt.Println("查询不到")
		c.Data["json"] = map[string]interface{}{"errcode": "1", "errmsg": "mobile does not exist"}
	}

	// this.Ctx.WriteString(rs)
	c.ServeJSON() //这个函数的作用是将上边的data按照json的方式进行传递，详见beego文档的多种格式输出部分

}