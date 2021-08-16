package main

import (
	"log"

	"github.com/go-god/gdi"
	"github.com/go-god/gdi/factory"
)

type service struct {
	App     string
	Config  *appConfig `inject:""`
	digInfo digInfo
}

type digInfo struct {
	app string
}

// AppConfig app config
type appConfig struct {
	Name string
}

// New create server entry
func New() *service {
	return &service{}
}

func createConfig() *appConfig {
	return &appConfig{
		Name: "dig-inject-demo",
	}
}

func createAbc(conf *appConfig) *digInfo {
	return &digInfo{
		app: conf.Name,
	}
}

func (s *service) GetConfig() *appConfig {
	return s.Config
}

func main() {
	di := factory.CreateDI(factory.DigInject)

	err := di.Provide(
		&gdi.Object{Value: New},
		&gdi.Object{Value: createConfig},
		&gdi.Object{Value: createAbc},
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
