package controllers

import (
	"github.com/astaxie/beego"
	"shortrent/models"
	"encoding/json"
	"fmt"
	"io"
	"crypto/sha1"
	"encoding/base32"
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
		//密码加密验证
		h := sha1.New()
		io.WriteString(h, login_data.Password)
		password_sh1 := base32.StdEncoding.EncodeToString(h.Sum(nil)) //进行base32编码操作
		
		if password_sh1 != maps[0]["up_passwd"] {
			// fmt.Println("password is wrong")
			c.Data["json"] = map[string]interface{}{"errcode": "1", "errmsg": "password is wrong "}
			
		} else {
			// fmt.Println(" verity correct ", password_sh1)
			c.Data["json"] = map[string]interface{}{"errcode": "0", "errmsg": "ok"}
		}		
	} else {
		fmt.Println("查询不到")
		c.Data["json"] = map[string]interface{}{"errcode": "1", "errmsg": "mobile does not exist"}
	}

	// 如下是设置 系统的session
	//该处的session的初始化代码在verify_code.go文件中的init函数
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)

	session_data := sess.Get("username")
	if session_data != nil {
		fmt.Println("username has existed , now, set it again ")
		sess.Delete("username")
		sess.Set("username", maps[0]["up_name"])
	} else {
		fmt.Println("set username session")
		sess.Set("username", maps[0]["up_name"])
	}
	// 设置 session end

	// this.Ctx.WriteString(rs)
	c.ServeJSON() //这个函数的作用是将上边的data按照json的方式进行传递，详见beego文档的多种格式输出部分

}