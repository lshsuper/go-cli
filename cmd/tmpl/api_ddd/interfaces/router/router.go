package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lshsuper/go-cli/cmd/tmpl/api_ddd/interfaces/handler/base"
)


var groupName=""
func Register(r *gin.Engine)  {

	g:=r.Group(groupName)
	//基础模块
	base.Register(g)


}
