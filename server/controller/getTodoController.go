package controller

import (
	"net/http"
	"remindme/server/controller/utils"
	"remindme/server/data"

	"github.com/gin-gonic/gin"
)

func GetAllTodosController(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, data.TodoData)
}

func GetTodoByID(context *gin.Context) {
	idString := context.Param("id")

	id, error := utils.ParseStringToUUID(idString)
	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "There was a problem converting id to int"})
		return
	}

	todo, error := utils.GetTodoID(id)
	if error != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	context.JSON(http.StatusOK, todo)
}
