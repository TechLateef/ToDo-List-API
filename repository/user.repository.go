package repository

import (
	"github.com/techlateef/tech-lateef-gol/models"
	"gorm.io/gorm"
)

type ToDoRepository interface {
	CreateToDo(t models.ToDoModels) models.ToDoModels
	UpdateTodo(t models.ToDoModels, todoId uint64) models.ToDoModels
	DeleteTodo(todoId uint64) models.ToDoModels
	FindTodoById(todoId uint64) models.ToDoModels
	GetALL() []models.ToDoModels
}

type todoRepository struct {
	connection *gorm.DB
}

func NewToDoRepository(todoRepo *gorm.DB) ToDoRepository {
	return &todoRepository{
		connection: todoRepo,
	}

}

func (db *todoRepository) CreateToDo(t models.ToDoModels) models.ToDoModels {
	db.connection.Save(&t)
	db.connection.Preload("Todo").Find(&t)
	return t
}

func (db *todoRepository) UpdateTodo(t models.ToDoModels, todoId uint64) models.ToDoModels {
	db.connection.Save(&t)
	db.connection.Preload("Todo").First(&t)
	return t
}
func (db *todoRepository) DeleteTodo(todoId uint64) models.ToDoModels {
	var ToDoDelete models.ToDoModels
	db.connection.Delete(&ToDoDelete, todoId)
	return ToDoDelete
}

func (db *todoRepository) GetALL() []models.ToDoModels {
	var Todos []models.ToDoModels
	db.connection.Preload("Todo").Find(&Todos)
	return Todos
}

func (db *todoRepository) FindTodoById(todoId uint64) models.ToDoModels {
	var todoWId models.ToDoModels
	db.connection.Preload("Todo").Find(&todoWId, todoId)
	return todoWId
}
