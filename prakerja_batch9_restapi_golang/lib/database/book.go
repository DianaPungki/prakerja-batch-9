package database

import (
	"dborm/config"
	"dborm/models"
	"errors"

	"fmt"

	"gorm.io/gorm"
)

func GetBooks() (interface{}, error) {
	var books []models.Book
	result := config.DB.Find(&books)
	err := result.Error
	if result.RowsAffected < 1 {
		err = fmt.Errorf("data book tidak ada")
		return nil, err
	}
	if result.Error != nil {
		return nil, err

	}

	return books, nil
}

func GetBook(judul string) (interface{}, error) {
	var book models.Book
	if err := config.DB.Where("judul = ?", judul).First(&book).Error; err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		return nil, err
	}
	return book, nil

}

func CreateBook(book models.Book) (interface{}, error) {

	if e := config.DB.Create(&book).Error; e != nil {
		return nil, e
	}
	return book, nil

}

func DeleteBook(judul string) (interface{}, error) {

	var book models.Book
	var err error
	result := config.DB.Delete(&book, "judul = ?", judul)

	err = result.Error

	if err != nil {
		return nil, err
	}

	if result.RowsAffected < 1 {
		err = fmt.Errorf("Baris id=%s tidak bisa dihapus, karena tidak ditemukan", judul)
		return nil, err
	}

	return book, nil

}

func UpdateBook(book models.Book, judul string) (interface{}, error) {

	var bookUpdated models.Book
	var err error
	result := config.DB.First(&bookUpdated, "judul = ?", judul)
	result = config.DB.Model(&bookUpdated).Updates(models.Book{Judul: book.Judul, Tahun: book.Tahun, Tempat: book.Tempat})

	err = result.Error

	if err != nil {
		return nil, err
	}

	return bookUpdated, nil

}
