package models

// import "gorm.io/gorm"

type Book struct {
	Judul     string `json:"judul" form:"judul" validate:"required" gorm:"primaryKey;unique;size:100"`
	Tahun    string `json:"tahun" form:"tahun" validate:"required" gorm:"size:100;not null"`
	Tempat string `json:"tempat" form:"tempat" validate:"required" gorm:"size:100;not null"`
}
