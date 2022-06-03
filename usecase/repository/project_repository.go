package repository

import "main/domain/model"

type ProjectRepository interface {
	FindAll(p []*model.Project) ([]*model.Project, error)
	Create(p *model.Project) (*model.Project, error)
}
