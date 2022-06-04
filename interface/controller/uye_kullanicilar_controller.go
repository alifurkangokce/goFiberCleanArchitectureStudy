package controller

import (
	"main/domain/model"
	"main/usecase/interactor"

	"github.com/gofiber/fiber/v2"
)

type uyeKullanicilarController struct {
	uyeKullanicilarInteractor interactor.UyeKullanicilarInteractor
}

// GetSatisEkipOzeti implements UyeKullanicilarController
func (uk *uyeKullanicilarController) GetSatisEkipOzeti(c *fiber.Ctx) error {
	var user []*model.User
	p, err := uk.uyeKullanicilarInteractor.GetSatisEkipOzeti(user)
	if err != nil {
		return err
	}
	return c.JSON(p)
}

type UyeKullanicilarController interface {
	GetSatisEkipOzeti(c *fiber.Ctx) error
}

func NewUyeKullanicilarController(uk interactor.UyeKullanicilarInteractor) UyeKullanicilarController {
	return &uyeKullanicilarController{uk}
}
