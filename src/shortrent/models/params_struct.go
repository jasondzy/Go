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