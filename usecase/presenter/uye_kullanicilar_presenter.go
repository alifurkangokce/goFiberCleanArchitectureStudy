package presenter

import "main/domain/model"

type UyeKullanicilarPresenter interface {
	ResponseGetSatisEkipOzeti(u []*model.User) []*model.User
}
