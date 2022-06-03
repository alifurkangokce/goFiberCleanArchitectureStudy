package interactor

import (
	"errors"
	"main/domain/model"
	"main/usecase/presenter"
	"main/usecase/repository"
)

type projectInteractor struct {
	ProjectRepository repository.ProjectRepository
	ProjectPresenter  presenter.ProjectPresenter
	DBRepository      repository.DBRepository
}

type ProjectInteractor interface {
	Get(p []*model.Project) ([]*model.Project, error)
	Create(p *model.Project) (*model.Project, error)
}

func NewProjectInteractor(r repository.ProjectRepository, p presenter.ProjectPresenter, d repository.DBRepository) *projectInteractor {
	return &projectInteractor{r, p, d}
}
func (pi *projectInteractor) Get(p []*model.Project) ([]*model.Project, error) {
	p, err := pi.ProjectRepository.FindAll(p)
	if err != nil {
		return nil, err
	}
	return pi.ProjectPresenter.ResponseProject(p), nil
}

func (pi *projectInteractor) Create(p *model.Project) (*model.Project, error) {
	data, err := pi.DBRepository.Transaction(func(i interface{}) (interface{}, error) {
		p, err := pi.ProjectRepository.Create(p)
		return p, err
	})
	project, ok := data.(*model.Project)
	if !ok {
		return nil, errors.New("Cast Error")
	}
	if !errors.Is(err, nil) {
		return nil, err
	}
	return project, nil

}
