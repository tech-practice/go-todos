package handlers

import (
	"go-todos/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateTodo(db database.TodoInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		var todo interface{}
		id := c.Param("id")
		err := c.BindJSON(&todo)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		res, err := db.Update(id, todo)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}
