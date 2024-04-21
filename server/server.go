package server

import (
	"remindme/server/controller"

	"github.com/gin-gonic/gin"
)

func MainServer() {
	serverPort := gin.Default()

	serverPort.POST("/todo", controller.CreateTodoController)
	serverPort.GET("/todo", controller.GetAllTodosController)
	serverPort.GET("/todo/:id", controller.GetTodoByID)
	serverPort.PATCH("/todo/updateByName", controller.UpdateTodoByNameController)
	serverPort.PATCH("todo/updateByCompletion", controller.UpdateTodoByCompletionController)
	serverPort.PATCH("/todo/update", controller.UpdateTodo)
	serverPort.DELETE("todo/deleteTodo/:id", controller.DeletedTodoById)
	serverPort.DELETE("todo/deleteTodo", controller.DeleteAllTodos)
	serverPort.Run("localhost:3001")
}
