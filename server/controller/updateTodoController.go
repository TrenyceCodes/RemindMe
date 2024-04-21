package controller

import (
	"net/http"
	"remindme/server/controller/utils"
	"remindme/server/models"

	"github.com/gin-gonic/gin"
)

func UpdateTodoByNameController(context *gin.Context) {
	//we will updateTodo by name
	idString, ok := context.GetQuery("id")

	if !ok {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Please provide an id"})
		return
	}

	uuidID, error := utils.ParseStringToUUID(idString)
	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "There was a problem converting id to int"})
		return
	}

	todo, error := utils.GetTodoID(uuidID)
	if error != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	var updatedTodo models.Todo
	if error := context.ShouldBindJSON(&updatedTodo); error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	todo.Name = updatedTodo.Name
	context.JSON(http.StatusOK, todo)
}

func UpdateTodoByCompletionController(context *gin.Context) {
	idString, ok := context.GetQuery("id")

	if !ok {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Please provide an id"})
		return
	}

	uuidID, error := utils.ParseStringToUUID(idString)
	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "There was an problem converting id string to id int"})
		return
	}

	todo, error := utils.GetTodoID(uuidID)
	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Todo not found"})
	}

	var updatedTodo models.Todo
	if error := context.ShouldBindJSON(&updatedTodo); error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
	}

	todo.Completed = updatedTodo.Completed
	context.JSON(http.StatusOK, todo)
}

func UpdateTodo(context *gin.Context) {
	idStr, ok := context.GetQuery("id")

	if !ok {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Please provide an id"})
		return
	}

	id, error := utils.ParseStringToUUID(idStr)
	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "There was an problem converting id string to id int"})
		return
	}

	todo, error := utils.GetTodoID(id)
	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Todo not found"})
	}

	var updatedTodo models.Todo
	if error := context.ShouldBindJSON(&updatedTodo); error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
	}

	todo.Name = updatedTodo.Name
	todo.Completed = updatedTodo.Completed
	context.JSON(http.StatusOK, todo)
}
