package models

type Task struct {
	ID        string
	Title     string
	Status    string
	ProjectID string
}

var Tasks = map[string]Task{}
