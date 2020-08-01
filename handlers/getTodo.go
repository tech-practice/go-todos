package handlers

import (
	"go-todos/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTodo(db database.TodoInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		res, err := db.Get(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}
