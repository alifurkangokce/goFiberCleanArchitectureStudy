package model

type User struct {
	AcilanFirma uint   `json:"acilanFirma"`
	Adsoyad     string `json:"adsoyad"`
}

func (User) TableName() string { return "uye_kullanicilar" }
