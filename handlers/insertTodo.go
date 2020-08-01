package handlers

import (
	"go-todos/database"
	"go-todos/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InsertTodo(db database.TodoInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		todo := models.Todo{}
		err := c.BindJSON(&todo)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		res, err := db.Insert(todo)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}
