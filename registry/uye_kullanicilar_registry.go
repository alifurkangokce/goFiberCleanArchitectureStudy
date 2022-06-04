package registry

import (
	"main/interface/controller"
	ip "main/interface/presenter"
	ir "main/interface/repository"
	"main/usecase/interactor"
	up "main/usecase/presenter"
	ur "main/usecase/repository"
)

func (r *registry) NewUyeKullanicilarController() controller.UyeKullanicilarController {
	return controller.NewUyeKullanicilarController(r.NewUyeKullanicilarInteractor())
}

func (r *registry) NewUyeKullanicilarInteractor() interactor.UyeKullanicilarInteractor {
	return interactor.NewUyeKullanicilarInteractor(r.NewUyeKullanicilarRepository(), r.NewUyeKullanicilarPresenter(), ir.NewDBRepository(r.db))
}

func (r *registry) NewUyeKullanicilarRepository() ur.UyeKullanicilarRepository {
	return ir.NewUyeKullanicilarRepository(r.db)
}

func (r *registry) NewUyeKullanicilarPresenter() up.UyeKullanicilarPresenter {
	return ip.NewUyeKullanicilarPresenter()
}
