package app

import (
	productController "cache_test/controllers/product"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

type Body struct {
	Name string `json:"name"`
}

func mapUrls() {
	// Products Mapping
	router.GET("/product/:product_id", productController.GetProductById)
	router.PUT("/product", productController.UpdateProduct)

	router.GET("/ejecucion", func(context *gin.Context) {
		body := Body{}

		body.Name = os.Getenv("HOSTNAME")
		context.JSON(http.StatusAccepted, &body)
	})

	log.Info("Finishing mappings configurations")
}
