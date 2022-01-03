package ctx

import (
	"github.com/lshsuper/go-cli/cmd/tmpl/api_ddd/infrastructure/common/inject"
)

//这里定义全局唯一变量
var(
	c *ctx

)

type ctx struct {

	UserID int
	IocContainer *inject.IOCContainer

}
//Register 注册
func Register()  {
	c=new(ctx)

	c.IocContainer=inject.NewIOCContainer()


}





