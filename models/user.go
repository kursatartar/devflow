package models

type User struct {
	ID    string
	Name  string
	Email string
}

var Users = map[string]User{}
