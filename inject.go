package gdi

// Injector di inject interface
type Injector interface {
	// Provide objects to the Graph. The Object documentation describes
	// the impact of various fields.
	Provide(objects ...*Object) error

	// Invoke the incomplete Objects.
	Invoke(values ...interface{}) error

	// String return inject type name.
	String() string
}

// Object an Object in the Graph.
type Object struct {
	Value interface{}
	Name  string // di name optional
	Group string // di group name optional
}
