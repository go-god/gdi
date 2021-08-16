package factory

import (
	"github.com/go-god/gdi"
	"github.com/go-god/gdi/dig"
	"github.com/go-god/gdi/fb"
)

type constructor func() gdi.Injector

// InjectType inject type
type InjectType int

const (
	// FbInject fb inject
	FbInject InjectType = iota + 1
	// DigInject dig inject
	DigInject
)

var injectTypeMap = map[InjectType]constructor{
	FbInject:  fb.New,
	DigInject: dig.New,
}

// Register register gdi.Injector
func Register(diType InjectType, c constructor) {
	_, ok := injectTypeMap[diType]
	if ok {
		panic("registered injector already exists")
	}

	injectTypeMap[diType] = c
}

// CreateDI create an injector interface.
func CreateDI(diType InjectType) gdi.Injector {
	factory, ok := injectTypeMap[diType]
	if !ok {
		panic("di type not exist")
	}

	return factory()
}
