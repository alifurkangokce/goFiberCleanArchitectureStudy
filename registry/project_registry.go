package registry

import (
	"main/interface/controller"
	ip "main/interface/presenter"
	ir "main/interface/repository"
	"main/usecase/interactor"
	up "main/usecase/presenter"
	ur "main/usecase/repository"
)

func (r *registry) NewProjectController() controller.ProjectController {
	return controller.NewProjectController(r.NewProjectInteractor())
}

func (r *registry) NewProjectInteractor() interactor.ProjectInteractor {
	return interactor.NewProjectInteractor(r.NewProjectRepository(), r.NewProjectPresenter(), ir.NewDBRepository(r.db))
}

func (r *registry) NewProjectRepository() ur.ProjectRepository {
	return ir.NewUserRepository(r.db)
}

func (r *registry) NewProjectPresenter() up.ProjectPresenter {
	return ip.NewProjectPresenter()
}
