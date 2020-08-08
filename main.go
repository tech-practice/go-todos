package main

import (
	"context"
	"go-todos/config"
	"go-todos/database"
	"go-todos/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	conf := config.GetConfig()
	ctx := context.TODO()

	db := database.ConnectDB(ctx, conf.Mongo)
	collection := db.Collection(conf.Mongo.Collection)

	client := &database.TodoClient{
		Col: collection,
		Ctx: ctx,
	}

	r := gin.Default()

	todos := r.Group("/todos")
	todos.Use(Authorization(conf.Token))
	{
		todos.GET("/", handlers.SearchTodos(client))
		todos.GET("/:id", handlers.GetTodo(client))
		todos.POST("/", handlers.InsertTodo(client))
		todos.PATCH("/:id", handlers.UpdateTodo(client))
		todos.DELETE("/:id", handlers.DeleteTodo(client))
	}

	r.POST("/graphql", handlers.GraphqlTodos(client))

	r.Run(":8080")
}

func Authorization(token string) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if token != auth {
			c.AbortWithStatusJSON(401, gin.H{"message": "Invalid authorization token"})
			return
		}
		c.Next()
	}
}
