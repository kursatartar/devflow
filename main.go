package main

import (
	"fmt"

	"github.com/kursatartar/devflowv2/handlers"
)

func main() {
	// USERS
	handlers.CreateUser("u1", "kursatartar", "kursat@example.com", "hash1", "admin", "Kürşat", "Artar", "avatar1.png")
	handlers.CreateUser("u2", "burakgurbuz", "burak@example.com", "hash2", "user", "Burak", "Gürbüz", "avatar2.png")
	handlers.CreateUser("u3", "hasanyilmaz", "hasan@example.com", "hash3", "user", "Hasan", "Yılmaz", "avatar3.png")
	handlers.CreateUser("u4", "coskunates", "coskun@example.com", "hash4", "user", "Coşkun", "Ateş", "avatar4.png")

	handlers.ListUsers()
	fmt.Println()

	// PROJECTS
	handlers.CreateProject(
		"p1", "DevFlow", "İlk CLI tabanlı proje", "u1", "active",
		[]string{"u1", "u2"}, false, []string{"todo", "in-progress", "done"},
	)
	handlers.CreateProject(
		"p2", "DevFlowV2", "Struct ve Mongo uyumlu versiyon", "u2", "planning",
		[]string{"u2", "u3"}, true, []string{"backlog", "review", "done"},
	)

	handlers.ListProjects()
	fmt.Println()
	handlers.CreateTask(
		"t1", "Struct Kullanımı", "Go'da struct nasıl tanımlanır ve kullanılır?",
		"p1", "u1", "u1", "todo", "medium", "2025-07-30T12:00:00Z",
		[]string{"go", "struct"}, 2.5, 0.5,
	)
	handlers.CreateTask(
		"t2", "Maps ile Veri Saklama", "Go'da map yapısıyla verileri nasıl organize ederiz?",
		"p2", "u2", "u2", "in-progress", "high", "2025-08-01T18:00:00Z",
		[]string{"go", "map"}, 3.0, 1.2,
	)
	handlers.CreateTask(
		"t3", "Pointer Mantığı", "Go'da pointer nasıl çalışır, neden kullanılır?",
		"p2", "u3", "u2", "todo", "low", "2025-08-03T15:30:00Z",
		[]string{"go", "pointer"}, 1.5, 0.0,
	)

	handlers.ListTasks()
}
