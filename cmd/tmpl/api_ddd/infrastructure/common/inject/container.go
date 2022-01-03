package inject

import (
	"reflect"
	"sync"
)

type Inject interface {
     ObjectName()string
}

type IOCContainer struct {

	rwLock *sync.RWMutex
	m map[string]Inject

}

func NewIOCContainer() *IOCContainer {
	return &IOCContainer{m: map[string]Inject{},rwLock:new(sync.RWMutex)}
}

func (c*IOCContainer)AddSingleton(objs ...Inject)  {
	for _,obj:=range objs{
		obj=c.m[obj.ObjectName()]
	}

}

func (c*IOCContainer)Resolve(objs ...Inject)  {

	defer c.rwLock.RUnlock()
	c.rwLock.RLock()
	for _,obj:=range objs{
		name:=reflect.TypeOf(obj).Name()
		obj=c.m[name]
	}

}





