package main

import "github.com/kursatartar/devflowv2/handlers"

func main() {
	handlers.CreateUser("1", "Kürşat", "kursatartar@k.com")
	handlers.CreateUser("2", "Burak", "burakgurbuz@b.com")
	handlers.CreateUser("2", "Hasan", "hasanyilmaz@h.com")
	handlers.CreateUser("2", "Coşkun", "coskunates@c.com")
	handlers.ListUsers()

	handlers.UpdateUser("1", "Kürşat Artar", "kursatartar@devflow.com")
	handlers.DeleteUser("2")
	handlers.ListUsers()

	handlers.CreateProject("p1", "DevFlow", "improving development skills")
	handlers.CreateProject("p2", "Go-bot", "slack bot in go")
	handlers.ListProjects()

	handlers.UpdateProject("p2", "Go-bot 2.0", "updated go bot project")
	handlers.DeleteProject("p1")
	handlers.ListProjects()

	handlers.CreateProject("p1", "DevFlow", "...")
	handlers.CreateTask("t1", "initialize repo", "todo", "p1")
	handlers.CreateTask("t2", "write handlers", "in-progress", "p1")
	handlers.ListTasks()
	handlers.UpdateTask("t2", "write handlers v2", "done", "p1")
	handlers.DeleteTask("t1")
	handlers.ListTasks()

}
