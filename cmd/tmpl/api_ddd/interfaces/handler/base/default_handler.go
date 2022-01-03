package base

import (
	"github.com/gin-gonic/gin"
	
)

type DefaultHandler struct {



}

func NewDefaultHandler() *DefaultHandler {
	return &DefaultHandler{}
}




func(c *DefaultHandler)Get(ctx *gin.Context){

	//ctx.JSON(http.StatusOK,c.defService.Get(1))
}

