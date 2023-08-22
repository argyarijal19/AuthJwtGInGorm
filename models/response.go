package models

type ReturnData struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Status  string `json:"status"`
}

type ListDataUser struct {
	ReturnData
	Data []UserSimgoa `json:"data"`
}
