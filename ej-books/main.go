package main

import (
	"fmt"
	"github.com/emikohmann/ucc-arqsoft-2/ej-books/router"
	"github.com/gin-gonic/gin"
)

var (
	ginRouter *gin.Engine
)

func main() {
	ginRouter = gin.Default()
	router.MapUrls(ginRouter)
	if err := ginRouter.Run(":8090"); err != nil {
		fmt.Println(fmt.Sprintf("Error running application: %v", err))
	}
}
