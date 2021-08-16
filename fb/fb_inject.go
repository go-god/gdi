package fb

import (
	fbInject "github.com/facebookgo/inject"

	"github.com/go-god/gdi"
)

// InjectName inject type name
const InjectName = "fb-inject"

type injectImpl struct {
	graph fbInject.Graph
}

// New return inject entry
func New() gdi.Injector {
	di := &injectImpl{}

	return di
}

// Provide objects to the Graph.
func (di *injectImpl) Provide(objects ...*gdi.Object) error {
	fbObjects := make([]*fbInject.Object, 0, len(objects)+1)
	for _, object := range objects {
		if object == nil {
			continue
		}

		fbObjects = append(fbObjects, &fbInject.Object{Name: object.Name, Value: object.Value})
	}

	return di.graph.Provide(fbObjects...)
}

// Invoke the incomplete Objects.
// values must be func() error
func (di *injectImpl) Invoke(values ...interface{}) error {
	err := di.graph.Populate()
	if err != nil {
		return err
	}

	for _, fn := range values {
		if callback, ok := fn.(func() error); ok {
			err = callback()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// String return inject type name.
func (di *injectImpl) String() string {
	return "di inject name: " + InjectName
}
