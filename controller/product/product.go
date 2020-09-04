package product

import (
	"github.com/gin-gonic/gin"
)

type ProductAdd struct {
	Name  string `form:"name" json:"name" validate:"required"`
	Intro string `form:"intro" json:"intro" validate:"required"`
}

func Store(c *gin.Context) {
	var productAdd ProductAdd



}
