package model

type Project struct {
	Id       uint    `json:"id"`
	Projeadi *string `json:"projeadi"`
	Mutadi   *string `json:"mutadi"`
}

func (Project) TableName() string { return "projeler" }
