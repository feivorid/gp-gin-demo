package route

import "github.com/gin-gonic/gin"

func SetupRouter(e *gin.Engine) {
	ProductRouter := e.Group("/products")
	{
		ProductRouter.POST("")
		ProductRouter.PUT("")
		ProductRouter.GET("")
		ProductRouter.DELETE("")
		ProductRouter.GET("")
	}
}
