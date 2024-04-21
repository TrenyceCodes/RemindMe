package utils

import (
	"errors"
	"remindme/server/data"
	"remindme/server/models"

	"github.com/google/uuid"
)

func GetTodoID(id uuid.UUID) (*models.Todo, error) {
	for index, todo := range data.TodoData {
		if todo.ID == id {
			return &data.TodoData[index], nil
		}
	}

	return nil, errors.New("there was an error looking for todo by id")
}

func DeleteTodoById(id uuid.UUID) error {
	for index, todo := range data.TodoData {
		if todo.ID == id {
			data.TodoData = append(data.TodoData[:index], data.TodoData[index+1:]...)
			return nil
		}
	}
	return errors.New("there was an error deleting id")
}

func ParseStringToUUID(id string) (uuid.UUID, error) {
	uuidId, error := uuid.Parse(id)

	if error != nil {
		return uuid.Nil, error
	}

	return uuidId, nil
}
