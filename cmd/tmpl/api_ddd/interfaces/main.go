package main

import (
	"github.com/facebookgo/inject"
	"github.com/gin-gonic/gin"
	"github.com/lshsuper/go-cli/cmd/tmpl/api_ddd/application/service"
	"github.com/lshsuper/go-cli/cmd/tmpl/api_ddd/infrastructure/repository"
	"github.com/lshsuper/go-cli/cmd/tmpl/api_ddd/interfaces/router"
)

func main()  {


	r:=gin.New()
	//注册路由
	router.Register(r)
	InitInject()
	r.Run(":10086")

}

func InitInject()  {
	graph := inject.Graph{}
	if err := graph.Provide(
		&inject.Object{
			Value: service.Register(),
		},
		&inject.Object{
			Value: repository.Register(),
		},
	); err != nil {
		panic(err)
	}

	if err := graph.Populate(); err != nil {
		panic(err)
	}
}
