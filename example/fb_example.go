package main

import (
	"log"

	"github.com/go-god/gdi"
	"github.com/go-god/gdi/factory"
)

type service struct {
	App    string
	Config *appConfig `inject:""`
}

// AppConfig app config
type appConfig struct {
	Name string
}

// New create server entry
func New() *service {
	return &service{}
}

func createConfig(name string) *appConfig {
	return &appConfig{
		Name: name,
	}
}

func (s *service) GetConfig() *appConfig {
	return s.Config
}

func main() {
	di := factory.CreateDI(factory.FbInject)

	err := di.Provide(
		&gdi.Object{Value: New()},
		&gdi.Object{Value: createConfig("gdi-demo")},
	)
	if err != nil {
		log.Fatalln("provide error: ", err)
	}

	err = di.Invoke(func(args ...interface{}) error {
		log.Println("service has run")
		return nil
	})

	if err != nil {
		log.Fatalln("invoke error: ", err)
	}

	// some code...
}
