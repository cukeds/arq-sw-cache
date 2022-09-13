package services

import (
	"github.com/emikohmann/ucc-arqsoft-2/ej-books/dtos"
	e "github.com/emikohmann/ucc-arqsoft-2/ej-books/utils/errors"
)

type Service interface {
	Get(id string) (dtos.BookDTO, e.ApiError)
	Insert(book dtos.BookDTO) (dtos.BookDTO, e.ApiError)
}
