package main

import (
	"github.com/gofiber/fiber/v2"
	"main/infrastructure/datastore"
	"main/infrastructure/router"
	"main/registry"
)

func main() {
	db := datastore.NewDB()
	r := registry.NewRegistry(db)
	f := fiber.New()
	router.NewRouter(f, r.NewAppController())
	f.Listen(":8000")
}
