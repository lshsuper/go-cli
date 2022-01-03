package base

import "github.com/gin-gonic/gin"

//用户分组
var groupName="user"

func Register(r *gin.RouterGroup)  {

	g:=r.Group(groupName)

	//default
	defHandler:=NewDefaultHandler()
	g.GET("/get",defHandler.Get)
	
}
