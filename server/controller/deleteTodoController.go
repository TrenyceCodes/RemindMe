package controller

import (
	"fmt"
	"net/http"
	"remindme/server/controller/utils"
	"remindme/server/data"
	"remindme/server/models"

	"github.com/gin-gonic/gin"
)

func DeletedTodoById(context *gin.Context) {
	idString := context.Param("id")

	uuidId, error := utils.ParseStringToUUID(idString)
	if error != nil {
		fmt.Println("Error parsing ID:", error)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "There was a problem parsing id"})
		return
	}

	err := utils.DeleteTodoById(uuidId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Todo not found"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}

func DeleteAllTodos(context *gin.Context) {
	data.TodoData = []models.Todo{}
	context.JSON(http.StatusOK, gin.H{"message": "Todos deleted successfully"})
}
