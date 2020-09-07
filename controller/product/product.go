package product

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-demo/model"
	"go-gin-demo/util"
	"net/http"
)

//type ProductAdd struct {
//	Name  string `form:"name" json:"name" validate:"required"`
//	Intro string `form:"intro" json:"intro" validate:"required"`
//}

func Index(c *gin.Context) {
	util.Json(http.StatusOK, "success", nil)
	fmt.Println("Hello World")
}

func Store(c *gin.Context) {

	var product model.Product
	product.Name = c.Request.FormValue("name")
	product.Intro = c.Request.FormValue("intro")

	res, err := product.Insert()

	if err != nil {
		util.Json(http.StatusBadRequest, "fail", err.Error())
	}

	util.Json(http.StatusOK, "success", res)
}
