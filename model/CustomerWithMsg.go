package model

type CustomerWithMsg struct {
	CId           int    `json:"cId"`
	CName         string `json:"cName"`
	CTel          int    `json:"cTel"`
	CAddress      string `json:"cAddress"`
	CRegisterDate string `json:"cRegisterDate"`
	CMsg          string `json:"msg"`
}
