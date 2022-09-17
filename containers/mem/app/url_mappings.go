package app

import (
	productController "cache_test/controllers/product"
	mcache "cache_test/memcached"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Body struct {
	Name string `json:"name"`
}

func mapUrls() {
	// Products Mapping
	router.GET("/product/:product_id", productController.GetProductById)
	router.PUT("/product", productController.UpdateProduct)

	router.GET("/flush", func(context *gin.Context) {
		mcache.Flush()
		context.JSON(http.StatusOK, "Cache Flushed")
	})

	log.Info("Finishing mappings configurations")
}
