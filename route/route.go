package route

import (
	"github.com/gin-gonic/gin"
	"go-gin-demo/controller/product"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	ProductRouter := router.Group("/products")
	ProductRouter.POST("", product.Store)
	ProductRouter.GET("", product.Index)
	//ProductRouter.GET("")
	//ProductRouter.DELETE("")
	//ProductRouter.GET("")

	return router
}
