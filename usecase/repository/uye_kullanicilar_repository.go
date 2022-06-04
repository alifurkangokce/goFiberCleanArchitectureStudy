package repository

import "main/domain/model"

type UyeKullanicilarRepository interface {
	GetSatisEkipOzeti(p []*model.User) ([]*model.User, error)
}
