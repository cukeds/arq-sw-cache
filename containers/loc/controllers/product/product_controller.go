package productController

import (
	cache "cache_test/cache"
	"cache_test/dto"
	service "cache_test/services"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Body struct {
	Name string `json:"name"`
}

func GetProductById(c *gin.Context) {

	body := Body{}
	body.Name = os.Getenv("HOSTNAME")
	log.Info("Container: " + body.Name)

	var productDto dto.ProductDto
	tmp_id := c.Param("product_id")

	if productDto, err := cache.Get(tmp_id); err == nil {
		c.JSON(http.StatusOK, productDto)
		log.Info("Data found in cache")
		return
	} else {
		log.Error(err)
	}

	id, _ := strconv.Atoi(tmp_id)
	productDto, err := service.ProductService.GetProductById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	log.Info("Data found in database")
	c.JSON(http.StatusOK, productDto)
	cache.Set(productDto)
}

func UpdateProduct(c *gin.Context) {

	body := Body{}
	body.Name = os.Getenv("HOSTNAME")
	log.Info("Container: " + body.Name)

	var productDto dto.ProductDto
	if err := c.BindJSON(&productDto); err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var productRDto dto.ProductDto
	productRDto, err := service.ProductService.UpdateProduct(productDto.ProductId, productDto.Description)
	if err != nil {
		log.Error(err.Error())
		c.JSON(err.Status(), err.Error())
		return
	}

	c.JSON(http.StatusCreated, productRDto)

	cache.Set(productRDto)
	log.Info("Data saved in cache")

}
