package repository

import (
	"errors"
	"gorm.io/gorm"
	"main/domain/model"
	"main/usecase/repository"
)

type projectRepository struct {
	db *gorm.DB
}

func (p2 projectRepository) FindAll(p []*model.Project) ([]*model.Project, error) {
	err := p2.db.Find(&p).Error

	if err != nil {
		return nil, err
	}

	return p, nil
}

func (p2 projectRepository) Create(p *model.Project) (*model.Project, error) {
	if err := p2.db.Create(p).Error; !errors.Is(err, nil) {
		return nil, err
	}
	
	return p, nil
}

func NewUserRepository(db *gorm.DB) repository.ProjectRepository {
	return &projectRepository{db}
}
