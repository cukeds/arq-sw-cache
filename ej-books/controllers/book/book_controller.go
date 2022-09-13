package book

import (
	"github.com/emikohmann/ucc-arqsoft-2/ej-books/dtos"
	service "github.com/emikohmann/ucc-arqsoft-2/ej-books/services"
	"github.com/emikohmann/ucc-arqsoft-2/ej-books/services/repositories"
	e "github.com/emikohmann/ucc-arqsoft-2/ej-books/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var (
	bookService = service.NewServiceImpl(
		repositories.NewRepositoryCache(1000, 100, 30*time.Second),
		repositories.NewRepositoryMemcached("localhost", 11211),
		repositories.NewRepositoryMongo("localhost", 27017, "books"),
	)
)

func Get(c *gin.Context) {
	book, apiErr := bookService.Get(c.Param("id"))
	if apiErr != nil {
		c.JSON(apiErr.Status(), apiErr)
		return
	}
	c.JSON(http.StatusOK, book)
}

func Insert(c *gin.Context) {
	var book dtos.BookDTO
	if err := c.BindJSON(&book); err != nil {
		apiErr := e.NewBadRequestApiError(err.Error())
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	book, apiErr := bookService.Insert(book)
	if apiErr != nil {
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	c.JSON(http.StatusCreated, book)
}
