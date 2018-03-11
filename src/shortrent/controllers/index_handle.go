package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"strconv"
	"shortrent/models"
	"encoding/json"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type IndexHandler struct {
	beego.Controller
}

type HouseinfoHandler struct {
	beego.Controller
}

type ShowDetail struct {
	beego.Controller
}

type OrderHandle struct {
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

//**************************************显示房间的详细信息***********************************************
func (c *ShowDetail) Get() {
	//获取URL中的参数
	var house_id string
	c.Ctx.Input.Bind(&house_id, "house_id")  //role变量保存传入的参数,从而获得url地址中传入的参数
	fmt.Println("house_id:",house_id)

	// 如下是获取系统的session
	//该处的session的初始化代码在verify_code.go文件中的init函数
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)

	session_data := sess.Get("username")
	if session_data != nil {
		fmt.Println("username has existed ")

		o := orm.NewOrm()
		o.Using("userinfo")

		var maps []orm.Params
		var maps_image []orm.Params
		var maps_houseinfo []orm.Params
		var maps_comment []orm.Params
		var house_data = make(map[string]interface{})
		var images = make([]string, 0)
		var facilities = make([]int, 0)
		var comment = make(map[string]interface{})
		var comments = make([]map[string]interface{}, 0)
		var integer int
		// var data = make([]map[string]interface{}, 0)
		num, err := o.Raw("select hi_title,hi_price,hi_address,hi_room_count,hi_acreage,hi_house_unit,hi_capacity,hi_beds,hi_deposit,hi_min_days,hi_max_days,up_name,up_avatar,hi_user_id from ih_house_info inner join ih_user_profile on hi_user_id=up_user_id where hi_house_id=?", house_id).Values(&maps)

		if err == nil && num > 0 {
			fmt.Println("maps=", maps) // slene	
			// for _, value := maps {
			house_data["hid"] = house_id
			house_data["user_id"] = maps[0]["hi_user_id"].(string)
			house_data["title"] = maps[0]["hi_title"].(string)
			house_data["price"] = maps[0]["hi_price"].(string)
			house_data["address"] = maps[0]["hi_address"].(string)
			house_data["room_count"] = maps[0]["hi_room_count"].(string)
			house_data["acreage"] = maps[0]["hi_acreage"].(string)
			house_data["unit"] = maps[0]["hi_house_unit"].(string)
			house_data["capacity"] = maps[0]["hi_capacity"].(string)
			house_data["beds"] = maps[0]["hi_beds"].(string)
			house_data["deposit"] = maps[0]["hi_deposit"].(string)
			house_data["min_days"] = maps[0]["hi_min_days"].(string)
			house_data["max_days"] = maps[0]["hi_max_days"].(string)
			house_data["user_name"] = maps[0]["up_name"].(string)
			house_data["user_avatar"] = maps[0]["up_avatar"].(string)

				// data = append(data, house_data)
			// }
			//如下查询的是房屋的全部图片相关信息
			num2, err2 := o.Raw("select hi_url from ih_house_image where hi_house_id=?", house_id).Values(&maps_image)
			if err2 == nil && num2 > 0 {
				for _,value2 := range maps_image {
					images = append(images, value2["hi_url"].(string))
				}
				
				house_data["images"] = images

			} else {
				fmt.Println("查询不到 房屋图片信息")
				c.Data["json"] = map[string]interface{}{"errcode": "4101", "errmsg": "can not find data from database "}
			}
			//如下查询的是房屋基本设施的信息
			num3, err3 := o.Raw("select hf_facility_id from ih_house_facility where hf_house_id=?", house_id).Values(&maps_houseinfo)
			if err3 == nil && num3 > 0 {
				
				for _,value3 := range maps_houseinfo {
					fmt.Println("=======")
					integer,_ = strconv.Atoi(value3["hf_facility_id"].(string))
					facilities = append(facilities, integer)
				}
				fmt.Println("facilities=",facilities)
				house_data["facilities"] = facilities

			} else {
				fmt.Println("查询不到 房屋基本设施信息")
				c.Data["json"] = map[string]interface{}{"errcode": "4101", "errmsg": "can not find data from database "}
			}

			//如下查询的是房屋的评论相关信息
			num4, err4 := o.Raw("select oi_comment,up_name,oi_utime,up_mobile from ih_order_info inner join ih_user_profile on oi_user_id=up_user_id where oi_house_id=? and oi_status=4 and oi_comment is not null", house_id).Values(&maps_comment)
			if err4 == nil && num4 > 0 {
				for _,value4 := range maps_comment {
					comment["user_name"] = value4["up_name"].(string)
					comment["content"] = value4["oi_comment"].(string)
					comment["ctime"] = value4["oi_utime"].(string)
					comment["user_id"] = value4["up_mobile"].(string)

					comments = append(comments, comment)
				}
				
				house_data["comments"] = comments

				c.Data["json"] = map[string]interface{}{"errcode": "0", "errmsg": "ok", "data":house_data, "user_id":comment["user_id"].(string)}

			} else {
				fmt.Println("查询不到 房屋基本设施信息")
				c.Data["json"] = map[string]interface{}{"errcode": "4101", "errmsg": "can not find data from database "}
			}



			
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

//******************************************end********************************************************

//******************************************订单预定****************************************************

func (c *OrderHandle) Post() {
	//解析传递的json数据
	var order_data models.Order_data
	json.Unmarshal(c.Ctx.Input.RequestBody, &order_data)
	fmt.Println(order_data)

	o := orm.NewOrm()
	o.Using("userinfo")

	var maps []orm.Params
	var user_info []orm.Params

	//获取用户的登录信息
	//该处的session的初始化代码在verify_code.go文件中的init函数
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)

	session_data := sess.Get("username")
	if session_data != nil {
		fmt.Println("username has existed  ")
		num2, err2 := o.Raw("select up_user_id from ih_user_profile where up_name=?", session_data).Values(&user_info)
		if err2 == nil && num2 > 0 {
			fmt.Println("user_id:",user_info[0]["up_user_id"].(string))
		} else {
			fmt.Println("username session does not exist")
			c.Data["json"] = map[string]interface{}{"errcode": "4101", "errmsg": "user not login"}
			c.ServeJSON()
			return 
		}

	} else {
		fmt.Println("username session does not exist")
		c.Data["json"] = map[string]interface{}{"errcode": "4101", "errmsg": "user not login"}
		c.ServeJSON()
		return 
	}
	// 设置 session end



	//用来判断预定的当前房屋是否存在
	num, err := o.Raw("select hi_price,hi_user_id from ih_house_info where hi_house_id=?", order_data.House_id).Values(&maps)
	if err == nil && num > 0 {
		fmt.Println("get house data suscess",maps) // slene	
		price,_ := strconv.Atoi(maps[0]["hi_price"].(string))

		res, err := o.Raw("insert into ih_order_info(oi_user_id,oi_house_id,oi_begin_date,oi_end_date,oi_days,oi_house_price,oi_amount) values(?, ?, ?, ?, ?, ?, ?);", user_info[0]["up_user_id"].(string), order_data.House_id, order_data.Start_data, order_data.End_date, 3, price, 3*price).Exec()
		if err == nil {
			num, _ := res.RowsAffected()
			fmt.Println("mysql row affected nums: ", num)
			c.Data["json"] = map[string]interface{}{"errcode": "0", "errmsg": "ok "}
		} else {
			fmt.Println("数据无法插入数据库")
			c.Data["json"] = map[string]interface{}{"errcode": "4004", "errmsg": "insert database wrong "}
		}
		
	} else {
		fmt.Println("查询不到")
		c.Data["json"] = map[string]interface{}{"errcode": "4101", "errmsg": "get house error"}
	}

	c.ServeJSON()

}

//*********************************************end**********************************************************