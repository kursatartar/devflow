package main

import (
	"context"
	"log"
	"time"

	"devflow/internal/config"
	"devflow/internal/db"
	"devflow/internal/handlers"
	repo "devflow/internal/persistence/mongodb"
	"devflow/internal/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.LoadConfig()

	if cfg.Database.DBName == "" {
		cfg.Database.DBName = "devflow"
	}
	if cfg.Database.MaxPool == 0 {
		cfg.Database.MaxPool = 10
	}
	if cfg.Database.Timeout == 0 {
		cfg.Database.Timeout = 10 * time.Second
	}

	mongo, err := db.NewMongo(
		cfg.Database.MongoURI,
		cfg.Database.DBName,
		cfg.Database.MaxPool,
		cfg.Database.Timeout,
	)
	if err != nil {
		log.Fatal("mongo connect error:", err)
	}
	defer mongo.Close(context.Background())

	userRepo := repo.NewUserRepository(mongo.Database)
	userSvc := services.NewUserService(userRepo)
	handlers.InitUserService(userSvc)

	taskRepo := repo.NewTaskRepository(mongo.Database)
	taskSvc := services.NewTaskService(taskRepo)
	handlers.InitTaskService(taskSvc)

	teamRepo := repo.NewTeamRepository(mongo.Database)
	teamSvc := services.NewTeamService(teamRepo)
	handlers.InitTeamService(teamSvc)

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			switch err {
			case services.ErrEmailExists:
				return c.Status(fiber.StatusConflict).JSON(fiber.Map{"success": false, "message": err.Error()})
			case services.ErrInvalidEmail:
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": err.Error()})
			default:
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "message": err.Error()})
			}
		},
	})

	app.Get("/", func(c *fiber.Ctx) error { return c.SendString("ok") })

	api := app.Group("/api")

	auth := api.Group("/auth")
	auth.Post("/register", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{"success": false, "message": "not implemented"})
	})
	auth.Post("/login", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{"success": false, "message": "not implemented"})
	})
	auth.Post("/refresh", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{"success": false, "message": "not implemented"})
	})
	auth.Post("/logout", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{"success": false, "message": "not implemented"})
	})

	users := api.Group("/users")
	users.Post("/", handlers.CreateUser)
	users.Get("/", handlers.ListUsers)
	users.Get("/:id", handlers.GetUser)
	users.Put("/:id", handlers.UpdateUser)
	users.Delete("/:id", handlers.DeleteUser)

	projects := api.Group("/projects")
	projects.Get("/", handlers.ListProjects)
	projects.Post("/", handlers.CreateProject)
	projects.Get("/:id", handlers.GetProject)
	projects.Put("/:id", handlers.UpdateProject)
	projects.Delete("/:id", handlers.DeleteProject)

	tasks := api.Group("/tasks")
	tasks.Post("", handlers.CreateTask)
	tasks.Get("", handlers.ListTasks)
	tasks.Get("/:id", handlers.GetTask)
	tasks.Put("/:id", handlers.UpdateTask)
	tasks.Delete("/:id", handlers.DeleteTask)

	teams := api.Group("/teams")
	teams.Get("/", handlers.ListTeams)
	teams.Post("/", handlers.CreateTeam)
	teams.Get("/:id", handlers.GetTeam)
	teams.Put("/:id", handlers.UpdateTeam)
	teams.Delete("/:id", handlers.DeleteTeam)
	teams.Post("/:id/members", handlers.AddTeamMember)

	log.Fatal(app.Listen(":8080"))
}
