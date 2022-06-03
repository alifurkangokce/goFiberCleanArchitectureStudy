package presenter

import "main/domain/model"

type ProjectPresenter interface {
	ResponseProject(p []*model.Project) []*model.Project
}
