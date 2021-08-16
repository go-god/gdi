package dig

import (
	"errors"

	"go.uber.org/dig"

	"github.com/go-god/gdi"
)

// InjectName inject type name
const InjectName = "fb-inject"

type injectImpl struct {
	container *dig.Container
}

// New return inject entry
func New() gdi.Injector {
	di := &injectImpl{
		container: dig.New(),
	}

	return di
}

// Provide objects to the Graph.
func (di *injectImpl) Provide(objects ...*gdi.Object) error {
	var err error
	for _, object := range objects {
		var opts []dig.ProvideOption
		if object.Name != "" {
			opts = append(opts, dig.Name(object.Name))
		}

		if object.Group != "" {
			opts = append(opts, dig.Name(object.Group))
		}

		if len(opts) > 0 {
			err = di.container.Provide(object.Value, opts...)
		} else {
			err = di.container.Provide(object.Value)
		}

		if err != nil {
			return errors.New("provide error: " + err.Error())
		}
	}

	return nil
}

// Invoke the incomplete Objects.
// values must be func() error
func (di *injectImpl) Invoke(values ...interface{}) error {
	if fnLen := len(values); fnLen == 0 || fnLen > 1 {
		return errors.New("the values parameter length must be 1")
	}

	return di.container.Invoke(values[0])
}

// String return inject type name.
func (di *injectImpl) String() string {
	return "di inject name: " + InjectName
}
