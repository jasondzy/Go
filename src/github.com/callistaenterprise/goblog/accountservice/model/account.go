package model

type Account struct {
	Id string `json:"id"`
	Name string `json:"name"`
	ServedBy string `json:"servedBy"`
}

type VipNotification struct {
	AccountId 	string 	`jason:"accountid"`
	ReadAt		string	`json:"readat"`
}