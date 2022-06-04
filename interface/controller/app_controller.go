package controller

type AppController struct {
	Project interface{ ProjectController }
	User    interface{ UyeKullanicilarController }
}
