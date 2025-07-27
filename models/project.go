package models

type Project struct {
	ID          string
	Name        string
	Description string
}

var Projects = map[string]Project{}
