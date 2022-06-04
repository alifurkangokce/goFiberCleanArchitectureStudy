package interactor

import (
	"main/domain/model"
	"main/usecase/presenter"
	"main/usecase/repository"
)

type uyeKullanicilarInteractor struct {
	UyeKullanicilarRepository repository.UyeKullanicilarRepository
	UyeKullanicilarPresenter  presenter.UyeKullanicilarPresenter
	DBRepository              repository.DBRepository
}

type UyeKullanicilarInteractor interface {
	GetSatisEkipOzeti([]*model.User) ([]*model.User, error)
}

func NewUyeKullanicilarInteractor(r repository.UyeKullanicilarRepository, p presenter.UyeKullanicilarPresenter, d repository.DBRepository) *uyeKullanicilarInteractor {
	return &uyeKullanicilarInteractor{r, p, d}
}
func (pi *uyeKullanicilarInteractor) GetSatisEkipOzeti(p []*model.User) ([]*model.User, error) {
	p, err := pi.UyeKullanicilarRepository.GetSatisEkipOzeti(p)
	if err != nil {
		return nil, err
	}
	return pi.UyeKullanicilarPresenter.ResponseGetSatisEkipOzeti(p), nil
}
