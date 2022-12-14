package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/techlateef/tech-lateef-gol/configs"
	"github.com/techlateef/tech-lateef-gol/controllers"
	"github.com/techlateef/tech-lateef-gol/repository"
	"github.com/techlateef/tech-lateef-gol/services"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                   = configs.ConnectToDB()
	todoRepository repository.ToDoRepository  = repository.NewToDoRepository(db)
	service        services.ToDoService       = services.NewToDoService(todoRepository)
	controller     controllers.TodoController = controllers.NewToDoController(service)
)

// Route func to server endpoints
func Routes() {

	route := gin.Default()

	route.POST("/todo", controller.CreateToDo)
	route.GET("/todo", controller.GetALL)
	route.PUT("/todo/:id", controller.UpdateTodo)
	route.DELETE("/todo/:id", controller.DeleteTodo)
	route.GET("/todo/:id", controller.FindTodoById)

	//Run route whenever triggered
	route.Run()
}
