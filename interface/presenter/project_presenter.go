package presenter

import (
	"main/domain/model"
	"main/usecase/presenter"
)

type projectPresenter struct {
}

func (p2 projectPresenter) ResponseProject(p []*model.Project) []*model.Project {
	//değişiklik yapabilirsin
	return p
}

func NewProjectPresenter() presenter.ProjectPresenter {
	return &projectPresenter{}
}
