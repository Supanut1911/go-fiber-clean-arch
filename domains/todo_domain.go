package domains

import "go-fiber-clean-arch/models"

type TodoUseCase interface {
	GetAllToDos() (t []models.Todo, err error)
	CreateAToDo() (t *models.Todo, err error)
	GetAToDo() (t *models.Todo, id string, err error)
	UpdateAToDo() (t *models.Todo, id string)
}