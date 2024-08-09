package main

import (
	TodoController "api/src/aggregates/todo/entrypoint/rest"

	"github.com/labstack/echo/v4"
)

func main() {
	s := echo.New()

	s.POST("/todo", TodoController.Store)
	s.GET("/todo/show", TodoController.Show)
	s.GET("/todo", TodoController.Index)
	s.PUT("/todo", TodoController.Update)
	s.DELETE("/todo", TodoController.Destroy)

	s.Logger.Fatal(s.Start(":8081"))
}
