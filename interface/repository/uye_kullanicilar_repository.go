package repository

import (
	"main/domain/model"
	"main/usecase/repository"

	"gorm.io/gorm"
)

type UyeKullanicilarRepository struct {
	db *gorm.DB
}

// GetSatisEkipOzeti implements repository.UyeKullanicilarRepository
func (u *UyeKullanicilarRepository) GetSatisEkipOzeti(p []*model.User) ([]*model.User, error) {

	u.db.Raw("call satisEkipOzetiSP()").Scan(&p)

	return p, nil
}

func NewUyeKullanicilarRepository(db *gorm.DB) repository.UyeKullanicilarRepository {
	return &UyeKullanicilarRepository{db}
}
