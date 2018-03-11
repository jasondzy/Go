package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type SerachareaShow struct {
	beego.Controller
}

type SearchinfoShow struct {
	beego.Controller
}

func(c *SerachareaShow) Get() {

		o := orm.NewOrm()
		o.Using("userinfo")

		var maps []orm.Params
		var data = make([]map[string]interface{}, 0)

		num, err := o.Raw("select ai_area_id,ai_name from ih_area_info").Values(&maps)
		if err == nil && num > 0 {
			fmt.Println(maps) // slene
			for _, value := range maps {
				// area := make(map[string]interface{})
				// area["area_id"] = value["ai_area_id"]
				// area["name"] = value["ai_name"]

				data = append(data, map[string]interface{}{"area_id":value["ai_area_id"], "name":value["ai_name"]})

				fmt.Println("data:",data)
			}

			c.Data["json"] = map[string]interface{}{"errcode": "0", "errmsg": "ok", "data": data}
		} else {
			fmt.Println("查询不到")
			c.Data["json"] = map[string]interface{}{"errcode": "4101", "errmsg": "can not find data from database "}
		}

		c.ServeJSON() //这个函数的作用是将上边的data按照json的方式进行传递，详见beego文档的多种格式输出部分

}

func (c *SearchinfoShow) Get() {
	//search.html?aid=2&aname=徐汇区&sd=2018-03-13&ed=2018-03-14
	var aid string
	c.Ctx.Input.Bind(&aid, "aid")
	var sd string
	c.Ctx.Input.Bind(&sd, "sd")
	var ed string
	c.Ctx.Input.Bind(&ed, "ed")

	fmt.Println("aid,aname,sd,ed",aid,sd,ed)

	o := orm.NewOrm()
	o.Using("userinfo")

	var maps []orm.Params
	var data = make([]map[string]interface{}, 0)
	num, err := o.Raw("select hi_title,hi_house_id,hi_price,hi_room_count,hi_address,hi_order_count,up_avatar,hi_index_image_url,hi_ctime from ih_house_info inner join ih_user_profile on hi_user_id=up_user_id left join ih_order_info on hi_house_id=oi_house_id where hi_area_id=?", aid).Values(&maps)
	if err == nil && num > 0 {
		fmt.Println(maps) // slene	
		for _, value := range maps {
			data = append(data, map[string]interface{}{"house_id":value["hi_house_id"], "title":value["hi_title"], "price":value["hi_price"], "room_count":value["hi_room_count"], "address":value["hi_address"], "order_count":value["hi_order_count"], "avatar":value["up_avatar"], "image_url":value["hi_index_image_url"]})
		}

		total_page := len(data)
		c.Data["json"] = map[string]interface{}{"errcode": "0", "errmsg": "ok", "data":data, "total_page":total_page}

	} else {
		fmt.Println("查询不到")
		c.Data["json"] = map[string]interface{}{"errcode": "0", "errmsg": "can not find data from database ", "total_page":0}
	}


	c.ServeJSON() //这个函数的作用是将上边的data按照json的方式进行传递，详见beego文档的多种格式输出部分
}