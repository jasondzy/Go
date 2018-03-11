package models

type Param_smscode struct {
	Mobile string `json:"mobile"`
	Piccode string `json:"piccode"`
	Piccode_id string `json:"piccode_id"`
}

type Register_data struct {
	Mobile string `json:"mobile"`
	Password string `json:"password"`
	Password2 string `json:"password2"`
}

type Login_data struct {
	Mobile string `json:"mobile"`
	Password string `json:"password"`
}

type Real_name struct {
	Real_name string `json:"real_name"`
	Id_card string `json:"id_card"`
}

type Name_modify struct {
	Name string `json:"name"`
}

type Order_data struct {
	House_id string `json:"house_id"`
	Start_data string `json:"start_date"`
	End_date string `json:"end_date"`
}
