package model

type Customers struct {
	Size      int              `json:"size"`
	Customers map[int]Customer `json:"customers"`
	CMsg      string           `json:"cMsg"`
}
