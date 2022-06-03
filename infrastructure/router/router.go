package router

import (
	"github.com/gofiber/fiber/v2"
	"main/interface/controller"
)

func NewRouter(f fiber.Router, c controller.AppController) {

	f.Get("/projects", func(ctx *fiber.Ctx) error {
		return c.Project.GetProject(ctx)
	})
	f.Post("/projects", func(ctx *fiber.Ctx) error {
		return c.Project.CreateProject(ctx)
	})
}
