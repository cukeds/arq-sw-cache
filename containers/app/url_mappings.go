package app

import (
	productController "cache_test/controllers/product"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"

	mcache "cache_test/memcached"
	log "github.com/sirupsen/logrus"
)

type Body struct {
	Name string `json:"name"`
}

func mapUrls() {
	// Products Mapping
	router.GET("/product/:product_id", productController.GetProductById)
	router.PUT("/product", productController.UpdateProduct)

	router.GET("/test", func(context *gin.Context) {
		body := Body{}

		body.Name = os.Getenv("HOSTNAME")
		context.JSON(http.StatusAccepted, &body)
	})

	router.GET("/flush", func(context *gin.Context) {
		mcache.Flush()
		context.JSON(http.StatusOK, "Cache Flushed")
	})

	log.Info("Finishing mappings configurations")
}
