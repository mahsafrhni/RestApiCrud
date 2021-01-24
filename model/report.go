package model

type Report struct {
	TotalCustomers int    `json:"totalcustomers"`
	Period         int    `json:"period"`
	Msg            string `json:"msg"`
}
