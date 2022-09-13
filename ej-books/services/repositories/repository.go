package repositories

import (
	"github.com/emikohmann/ucc-arqsoft-2/ej-books/dtos"
	"github.com/emikohmann/ucc-arqsoft-2/ej-books/utils/errors"
)

type Repository interface {
	Get(id string) (dtos.BookDTO, errors.ApiError)
	Insert(book dtos.BookDTO) (dtos.BookDTO, errors.ApiError)
	Update(book dtos.BookDTO) (dtos.BookDTO, errors.ApiError)
	Delete(id string) errors.ApiError
}
