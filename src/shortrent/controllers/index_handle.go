package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type IndexHandler struct {
	beego.Controller
}

type HouseinfoHandler struct {
	beego.Controller
}

func (c *IndexHandler) Get() {

	// 如下是设置 系统的session
	//该处的session的初始化代码在verify_code.go文件中的init函数
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)

	session_data := sess.Get("username")
	if session_data != nil {
		fmt.Println("username exist ")
		var data = map[string]string{"name":session_data.(string)}
		c.Data["json"] = map[string]interface{}{"errcode": "0", "data": data}

	} else {
		fmt.Println("user not login ")
		c.Data["json"] = map[string]interface{}{"errcode": "1", "data": nil}
	}
	// 设置 session end
	c.ServeJSON() //这个函数的作用是将上边的data按照json的方式进行传递，详见beego文档的多种格式输出部分
}

func (c *HouseinfoHandler) Get() {

	o := orm.NewOrm()
	o.Using("userinfo")

	var maps []orm.Params
	var houses = make([]map[string]interface{}, 0)
	num, err := o.Raw("select hi_house_id,hi_title,hi_index_image_url from ih_house_info order by hi_order_count desc limit 3;").Values(&maps)
	// fmt.Println(maps)

	//如下查询首页的房屋信息
	if err == nil && num > 0 {
		fmt.Println("查询到数据")
		for _,value := range maps {
			// fmt.Println(value)
			houses = append(houses, map[string]interface{}{"house_id":value["hi_house_id"], "img_url":value["hi_index_image_url"], "title":value["hi_title"]})
		}
		// fmt.Println(houses)

	} else {
		fmt.Println("查询不到房屋信息")
	}
	//end
	//如下查询地区信息
	var maps_area []orm.Params
	var area = make([]map[string]interface{}, 0)
	num, err = o.Raw("select * from ih_area_info").Values(&maps_area)
	// fmt.Println(maps_area)
	if err == nil && num > 0 {
	fmt.Println("查询到数据")
	for _,value := range maps_area {
		// fmt.Println(value)
		area = append(area, map[string]interface{}{"area_id":value["ai_area_id"], "name":value["ai_name"]})
		}
		// fmt.Println(area)

	} else {
		fmt.Println("查询不到地区信息")
	}

	c.Data["json"] = map[string]interface{}{"errcode": "0", "errmsg": "ok", "houses":houses, "areas":area}
	c.ServeJSON() //这个函数的作用是将上边的data按照json的方式进行传递，详见beego文档的多种格式输出部分
}