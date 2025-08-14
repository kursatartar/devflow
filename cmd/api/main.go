package main

import (
	"devflow/internal/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
)

func ListResp(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
		"message": "Rules fetched successfully",
	})
}

func ItemResp(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
		"message": "Rule :id fetched successfully",
	})
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error { return c.SendString("ok") })

	api := app.Group("/api")

	auth := api.Group("/auth")
	auth.Post("/register", ListResp)
	auth.Post("/login", ListResp)
	auth.Post("/refresh", ListResp)
	auth.Post("/logout", ListResp)

	users := api.Group("/users")
	users.Post("/", handlers.CreateUser)
	users.Get("/", ListResp)
	users.Get("/:id", ItemResp)
	users.Put("/:id", ItemResp)
	users.Delete("/:id", ItemResp)

	projects := api.Group("/projects")
	projects.Get("/", ListResp)
	projects.Post("/", ListResp)
	projects.Get("/:id", ItemResp)
	projects.Put("/:id", ItemResp)
	projects.Delete("/:id", ItemResp)

	projects.Get("/:id/tasks", ListResp)
	projects.Post("/:id/tasks", ListResp)

	tasks := api.Group("/tasks")
	tasks.Get("/:id", ItemResp)
	tasks.Put("/:id", ItemResp)
	tasks.Delete("/:id", ItemResp)

	teams := api.Group("/teams")
	teams.Get("/", ListResp)
	teams.Post("/", ListResp)
	teams.Get("/:id", ItemResp)
	teams.Put("/:id", ItemResp)
	teams.Post("/:id/members", ListResp)

	log.Fatal(app.Listen(":8080"))
}
