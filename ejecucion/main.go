package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	engine := gin.New()
	engine.GET("/start", func(c *gin.Context) {
		var t time.Duration
		t = 50 * time.Millisecond
		for true {
			time.Sleep(t)
			http.Get("http://localhost/mem/product/1")
		}
	})

	engine.Run(":3000")
}
