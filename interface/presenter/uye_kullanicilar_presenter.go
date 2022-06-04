package presenter

import (
	"main/domain/model"
	"main/usecase/presenter"
)

type UyeKullanicilarPresenter struct {
}

// ResponseGetSatisEkipOzeti implements presenter.UyeKullanicilarPresenter
func (uk *UyeKullanicilarPresenter) ResponseGetSatisEkipOzeti(u []*model.User) []*model.User {
	//değişiklik yapabilirsin
	return u
}

func NewUyeKullanicilarPresenter() presenter.UyeKullanicilarPresenter {
	return &UyeKullanicilarPresenter{}
}
