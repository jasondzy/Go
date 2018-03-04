package models

import (
    "github.com/astaxie/beego/orm"
)

/***********如下实现了一个struct结构体的定义，对应的是数据库中的表test_tbl
//   需要注意的是结构体的名称需要以大写字符开头，其名称要和数据库中的表格名称对应
type Test_tbl struct {
	Id int
	Test_title string
	Test_author string
	// Submission_date string
}
********************************************************************/

//这里定义的是user info结构，对应的是数据库中的ih_user_profile表
type Ih_user_profile struct {
	Id int
	Up_name string
	Up_mobile string
	Up_passwd string
	Up_real_name string
	Up_id_card string
	Up_avatar string 
	Up_admin int
}

func init() {
    // 需要在init中注册定义的model
    orm.RegisterModel(new(Ih_user_profile))
}
