package app

import (
	productController "cache_test/controllers/product"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

type Body struct {
	Name string `json:"name"`
}

func mapUrls() {
	// Products Mapping
	router.GET("/product/:product_id", productController.GetProductById)
	router.GET("/products", productController.GetProducts)
	router.GET("/products/:category_id", productController.GetProductsByCategoryId)
	router.GET("/products/search=:searchQuery", productController.GetProductsBySearch)

	router.GET("/test", func(context *gin.Context) {
		body := Body{}

		body.Name = os.Getenv("HOSTNAME")
		context.JSON(http.StatusAccepted, &body)
	})

	log.Info("Finishing mappings configurations")
}
