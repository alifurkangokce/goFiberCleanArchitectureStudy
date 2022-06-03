package controller

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"main/domain/model"
	"main/usecase/interactor"
)

type projectController struct {
	projectInteractor interactor.ProjectInteractor
}

func (pc projectController) GetProject(c *fiber.Ctx) error {
	var p []*model.Project
	p, err := pc.projectInteractor.Get(p)
	if err != nil {
		return err
	}
	return c.JSON(p)
}

func (pc projectController) CreateProject(c *fiber.Ctx) error {
	var params model.Project
	if err := c.BodyParser(&params); !errors.Is(err, nil) {
		return err
	}
	u, err := pc.projectInteractor.Create(&params)
	if !errors.Is(err, nil) {
		return err
	}
	return c.JSON(u)
}

type ProjectController interface {
	GetProject(c *fiber.Ctx) error
	CreateProject(c *fiber.Ctx) error
}

func NewProjectController(pi interactor.ProjectInteractor) ProjectController {
	return &projectController{pi}
}
