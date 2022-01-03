package base

import (
	"github.com/lshsuper/go-cli/cmd/tmpl/api_ddd/domain/entity/base"
)

type IDefaultRepository interface {

	GetDefault(id int) base.DefEntity

}
