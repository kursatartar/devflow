package main

import (
	"devflow/internal/models"
	"fmt"

	"devflow/internal/handlers"
)

func main() {
	// USERS
	handlers.CreateUser("u1", "kursatartar", "kursat@example.com", "hash1", "admin",
		models.Profile{FirstName: "Kürşat", LastName: "Artar", AvatarURL: "avatar1.png"})

	handlers.CreateUser("u2", "burakgurbuz", "burak@example.com", "hash2", "user",
		models.Profile{FirstName: "Burak", LastName: "Gürbüz", AvatarURL: "avatar2.png"})

	handlers.CreateUser("u3", "hasanyilmaz", "hasan@example.com", "hash3", "user",
		models.Profile{FirstName: "Hasan", LastName: "Yılmaz", AvatarURL: "avatar3.png"})

	handlers.CreateUser("u4", "coskunates", "coskun@example.com", "hash4", "user",
		models.Profile{FirstName: "Coşkun", LastName: "Ateş", AvatarURL: "avatar4.png"})
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

	fmt.Println("only admins:")
	handlers.ListUsersByRole("admin")

}

/*
package main

import (
	"flag"
	"fmt"

	"devflow/internal/models"
	"devflow/internal/handlers"
)

func main() {
	action := flag.String("action", "", "create veya list")
	username := flag.String("username", "", "kullanıcı adı")
	email := flag.String("email", "", "e-posta")
	password := flag.String("password", "", "şifre hash")
	role := flag.String("role", "user", "kullanıcı rolü")
	firstName := flag.String("firstName", "", "ad")
	lastName := flag.String("lastName", "", "soyad")
	avatar := flag.String("avatar", "", "avatar url")

	flag.Parse()

	switch *action {
	case "create":
		if *username == "" || *email == "" || *password == "" || *firstName == "" || *lastName == "" {
			fmt.Println("eksik bilgi")
			return
		}
		id := *username
		profile := models.Profile{
			FirstName: *firstName,
			LastName:  *lastName,
			AvatarURL: *avatar,
		}
		handlers.CreateUser(id, *username, *email, *password, *role, profile)

	case "list":
		handlers.ListUsers()

	default:
		fmt.Println("geçerli bir action gir: create veya list")
	}
*/
