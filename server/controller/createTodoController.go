package controller

import (
	"net/http"
	"remindme/server/data"
	"remindme/server/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateTodoController(context *gin.Context) {
	//create a todo
	var todo models.Todo
	error := context.BindJSON(&todo)

	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}

	todo.ID = uuid.New()
	data.TodoData = append(data.TodoData, todo)
	context.IndentedJSON(http.StatusOK, todo)
}
