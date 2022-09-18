package app

import "github.com/DProject89/cmsfoto/app/models"

type Model struct {
	Model interface{}
}

func RegisterModel() []Model {
	return []Model{
		{Model: models.User{}},
	}
}
