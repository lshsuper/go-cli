package repository

import (

	"github.com/lshsuper/go-cli/cmd/tmpl/api_ddd/domain/irepository/base"
	baseRepository"github.com/lshsuper/go-cli/cmd/tmpl/api_ddd/infrastructure/repository/base"
)

type Repository struct {
	DefaultRepository base.IDefaultRepository
}

func Register()*Repository  {

	return &Repository{
		DefaultRepository:baseRepository.NewDefaultRepository(),
	}

}