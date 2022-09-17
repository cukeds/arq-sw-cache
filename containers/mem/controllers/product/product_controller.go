package productController

import (
	"cache_test/dto"
	mcache "cache_test/memcached"
	service "cache_test/services"
	"encoding/json"
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
	l, _ := service.ProductService.GetProductById(1)
	log.Info(l)

	var productDto dto.ProductDto
	tmp_id := c.Param("product_id")

	if productDto, err := mcache.Get(tmp_id); err == nil {
		c.JSON(http.StatusOK, productDto)
		log.Info("Data found in memcached")
		return
	} else {
		log.Error(err)
	}

	id, _ := strconv.Atoi(tmp_id)
	log.Info(id)
	productDto, err := service.ProductService.GetProductById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	log.Info("Data found in database")
	c.JSON(http.StatusOK, productDto)

	res, _ := json.Marshal(productDto)
	mcache.Set(tmp_id, res)
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

	id := string(productRDto.ProductId)
	setter, _ := json.Marshal(productRDto)
	mcache.Set(id, setter)
	log.Info("Data saved in cache")

}
