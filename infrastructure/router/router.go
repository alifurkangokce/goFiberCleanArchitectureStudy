package router

import (
	"main/interface/controller"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(f fiber.Router, c controller.AppController) {

	f.Get("/projects", func(ctx *fiber.Ctx) error {
		return c.Project.GetProject(ctx)
	})
	f.Post("/projects", func(ctx *fiber.Ctx) error {
		return c.Project.CreateProject(ctx)
	})

	f.Get("/satisEkipOzeti", func(ctx *fiber.Ctx) error {
		return c.User.GetSatisEkipOzeti(ctx)
	})
}
