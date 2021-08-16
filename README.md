# gdi

    go inject tools
    This is a di tool encapsulated by the dependency injection framework and 
    shielding the underlying details. 
    It is mainly used to build go projects and solve the dependencies of go 
    project components

# create di inject interface

## use fb-inject
```go
package main

import (
	"log"

	"github.com/go-god/gdi"
	"github.com/go-god/gdi/factory"
)

func main() {
	di := factory.CreateDI(factory.FbInject)
	// fn is func or interface
	err := di.Provide(gdi.Object{Value: fn})
	if err != nil {
		log.Fatalln("provide error: ",err)
	}
	
	err = di.Invoke()
	if err != nil{
		log.Fatalln("invoke error: ", err)
    }
    
    // some code...
}

```
## use dig-inject
```go
package main

import (
	"log"

	"github.com/go-god/gdi"
	"github.com/go-god/gdi/factory"
)

func main() {
	di := factory.CreateDI(factory.DigInject)
	// fn is func or interface
	err := di.Provide(gdi.Object{Value: fn})
	if err != nil {
		log.Fatalln("provide error: ", err)
	}
	
	err = di.Invoke(func() {
		// start service...
    })
	if err != nil{
		log.Fatalln("invoke error: ",err)
    }
    
    // some code...
}

```

# Dependency injection framework

    facebook inject: github.com/facebookgo/inject
    uber-go/dig: go.uber.org/dig
