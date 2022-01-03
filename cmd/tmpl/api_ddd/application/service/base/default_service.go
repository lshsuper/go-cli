package base

import (
	"github.com/lshsuper/go-cli/cmd/tmpl/api_ddd/domain/entity/base"
	baseReponsitory "github.com/lshsuper/go-cli/cmd/tmpl/api_ddd/domain/irepository/base"
)

type DefaultService struct {
	_defaultRepository baseReponsitory.IDefaultRepository `inject:""`
}

func NewDefaultService(_defaultRepository baseReponsitory.IDefaultRepository) *DefaultService {
	return &DefaultService{_defaultRepository: _defaultRepository}
}

func (s DefaultService)Get(id int)base.DefEntity  {
	return s._defaultRepository.GetDefault(id)
}
