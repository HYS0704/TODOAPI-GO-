package main

import (
	"net/http"
	"errors"
	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string	`json:"id"`
	Item      string	`json:"item"`
	Completed bool		`json:"completed"`
}

var todos = []todo{}

func getTodos(context *gin.Context)  {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context)  {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodo(context *gin.Context)  {
	id := context.Param("id")
	todoPtr, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message":"Todo not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todoPtr)
}

func toggleTodoStatus(context *gin.Context)  {
	id := context.Param("id")
	todoPtr, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message":"Todo not found"})
		return
	}

	todoPtr.Completed = !todoPtr.Completed

	context.IndentedJSON(http.StatusOK, todoPtr)
}

func getTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID ==id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("todo not found")
}

func deleteTodo(context *gin.Context)  {
	id := context.Param("id")
	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			context.IndentedJSON(http.StatusOK, gin.H{"message":"todo deleted"})
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message":"Todo not found"})
}

func main()  {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.POST("/todos", addTodo)
	router.DELETE("/todos/:id", deleteTodo)
	router.Run(":8080")
}
