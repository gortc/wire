// Code generated by Wire. DO NOT EDIT.

//go:generate go run gortc.io/wire/cmd/wire
//+build !wireinject

package main

import (
	"example.com/bar"
	"example.com/baz"
	"example.com/foo"
	"fmt"
)

// Injectors from wire.go:

func newBazService(config *baz.Config) *baz.Service {
	fooConfig := config.Foo
	service := foo.New(fooConfig)
	barConfig := config.Bar
	barService := bar.New(barConfig, service)
	bazService := &baz.Service{
		Foo: service,
		Bar: barService,
	}
	return bazService
}

// wire.go:

func main() {
	cfg := &baz.Config{
		Foo: &foo.Config{1},
		Bar: &bar.Config{2},
	}
	svc := newBazService(cfg)
	fmt.Println(svc.String())
}
