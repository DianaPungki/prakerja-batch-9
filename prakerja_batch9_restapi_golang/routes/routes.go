package routes

import (
	"dborm/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	e.POST("/books", controllers.CreateBookController)
	e.GET("/books", controllers.GetBooksController)
	e.PUT("/books/:judul", controllers.UpdateBookController)
	e.DELETE("/books/:judul", controllers.DeleteBookController)
	e.GET("/books/:judul", controllers.GetBookController)

	return e
}