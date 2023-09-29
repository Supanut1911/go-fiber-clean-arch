package models

type Todo struct {
	ID uint
	Tittle string
	Description string
}

func (b *Todo) TableName() string {
	return "todo"
}