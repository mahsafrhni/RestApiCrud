package model

type Customer struct {
	CId           int    `json:"cId"`
	CName         string `json:"cName"`
	CTel          int    `json:"cTel"`
	CAddress      string `json:"cAddress"`
	CRegisterDate string `json:"cRegisterDate"`
}
