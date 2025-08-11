package main

import (
	"devflow/internal/handlers"
	"devflow/internal/models"
	"devflow/tools/generics"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	section := flag.String("section", "", "users, projects, tasks, generics")
	action := flag.String("action", "", "create, list")

	username := flag.String("username", "", "")
	email := flag.String("email", "", "")
	password := flag.String("password", "", "")
	role := flag.String("role", "user", "")
	firstName := flag.String("firstName", "", "")
	lastName := flag.String("lastName", "", "")
	avatar := flag.String("avatar", "", "")

	projectID := flag.String("projectID", "", "")
	projectName := flag.String("projectName", "", "")
	projectDesc := flag.String("projectDesc", "", "")
	projectOwner := flag.String("ownerID", "", "")
	projectStatus := flag.String("projectStatus", "planning", "")
	projectPrivate := flag.Bool("private", false, "")
	projectMembers := flag.String("team", "", "")
	projectWorkflow := flag.String("workflow", "", "")

	flag.Parse()

	switch *section {
	case "users":
		handleUserSection(*action, *username, *email, *password, *role, *firstName, *lastName, *avatar)
	case "projects":
		handleProjectSection(*action, *projectID, *projectName, *projectDesc, *projectOwner, *projectStatus, *projectPrivate, *projectMembers, *projectWorkflow)
	case "generics":
		runGenericsTest()
	default:
		fmt.Println("gecerli bir section girilmeli")
	}
}

func handleUserSection(action, username, email, password, role, firstName, lastName, avatar string) {
	switch action {
	case "create":
		handlers.CreateUser(username, username, email, password, role, firstName, lastName, avatar)
	case "list":
		handlers.ListUsers()
	default:
		fmt.Println("gecerli bir action girilmeli")
	}
}

func handleProjectSection(action, id, name, desc, ownerID, status string, isPrivate bool, membersRaw, workflowRaw string) {
	switch action {
	case "create":
		if id == "" || name == "" || ownerID == "" {
			fmt.Println("eksik bilgi")
			os.Exit(1)
		}
		team := splitCSV(membersRaw)
		workflow := splitCSV(workflowRaw)
		handlers.CreateProject(id, name, desc, ownerID, status, team, isPrivate, workflow)
	case "list":
		handlers.ListProjects()
	default:
		fmt.Println("gecerli bir action girilmeli")
	}
}

func splitCSV(input string) []string {
	if input == "" {
		return []string{}
	}
	parts := strings.Split(input, ",")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return parts
}

func runGenericsTest() {
	users := []*models.User{
		{ID: "u1", Username: "kursat"},
		{ID: "u2", Username: "burak"},
	}
	found := generics.FindByID(users, "u2")
	if found != nil {
		fmt.Println("bulundu:", found.Username)
	} else {
		fmt.Println("bulunamadı")
	}
}
