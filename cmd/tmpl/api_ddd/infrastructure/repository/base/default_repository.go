package base

import (
	"github.com/lshsuper/go-cli/cmd/tmpl/api_ddd/domain/entity/base"
)

type DefaultRepository struct {

}

func NewDefaultRepository() *DefaultRepository {
	return &DefaultRepository{}
}


func (dao *DefaultRepository)GetDefault(id int) base.DefEntity {
	return base.DefEntity{
		Id: id,
	}
}