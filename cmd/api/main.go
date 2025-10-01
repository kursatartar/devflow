package main

import (
	"context"
	handlers2 "devflow/internal/presentation/api/handlers"
	"log"
	"os"
	"time"

	"devflow/internal/config"
	"devflow/internal/db"
	repo "devflow/internal/persistence/mongodb/repositories"
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
	handlers2.InitUserService(userSvc)

	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		secretKey = "devflow-secret-key"
	}
	authSvc := services.NewAuthService(secretKey)
	handlers2.InitAuthService(authSvc)

	taskRepo := repo.NewTaskRepository(mongo.Database)
	taskSvc := services.NewTaskService(taskRepo)
	handlers2.InitTaskService(taskSvc)

	teamRepo := repo.NewTeamRepository(mongo.Database)
	teamSvc := services.NewTeamService(teamRepo)
	handlers2.InitTeamService(teamSvc)

	projectRepo := repo.NewProjectRepository(mongo.Database)
	projectSvc := services.NewProjectService(projectRepo)
	handlers2.InitProjectService(projectSvc)

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
	auth.Post("/register", handlers2.Register)
	auth.Post("/login", handlers2.Login)

	users := api.Group("/users")
	users.Post("/", handlers2.CreateUser)
	users.Get("/", handlers2.ListUsers)
	users.Get("/:id", handlers2.GetUser)
	users.Put("/:id", handlers2.UpdateUser)
	users.Delete("/:id", handlers2.DeleteUser)

	projects := api.Group("/projects")
	projects.Get("/", handlers2.ListProjects)
	projects.Post("/", handlers2.CreateProject)
	projects.Get("/:id", handlers2.GetProject)
	projects.Put("/:id", handlers2.UpdateProject)
	projects.Delete("/:id", handlers2.DeleteProject)

	tasks := api.Group("/tasks")
	tasks.Post("", handlers2.CreateTask)
	tasks.Get("", handlers2.ListTasks)
	tasks.Get("/:id", handlers2.GetTask)
	tasks.Put("/:id", handlers2.UpdateTask)
	tasks.Delete("/:id", handlers2.DeleteTask)

	teams := api.Group("/teams")
	teams.Get("/", handlers2.ListTeams)
	teams.Post("/", handlers2.CreateTeam)
	teams.Get("/:id", handlers2.GetTeam)
	teams.Put("/:id", handlers2.UpdateTeam)
	teams.Delete("/:id", handlers2.DeleteTeam)
	teams.Post("/:id/members", handlers2.AddTeamMember)

	log.Fatal(app.Listen(":8080"))
}
