package controllers

import (
	"dborm/lib/database"
	"dborm/models"

	"net/http"

	"github.com/labstack/echo/v4"
)

func GetBooksController(c echo.Context) error {
	books, err := database.GetBooks()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succses",
		"data":    books,
	})
}

func GetBookController(c echo.Context) error {

	judulBook := c.Param("judul")
	book, err := database.GetBook(judulBook)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succses",
		"data":    book,
	})

}

func CreateBookController(c echo.Context) error {

	var book = models.Book{}
	c.Bind(&book)
	

	bookCreate, err := database.CreateBook(book)
	if err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succses",
		"data":    bookCreate,
	})

}

func DeleteBookController(c echo.Context) error {

	judulBook := c.Param("judul")
	_, err := database.DeleteBook(judulBook)
	if err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succses",
	})

}

func UpdateBookController(c echo.Context) error {

	var book = models.Book{}
	c.Bind(&book)
	judulBook := c.Param("judul")

	bookCreate, err := database.UpdateBook(book, judulBook)
	if err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succses",
		"data":    bookCreate,
	})

}