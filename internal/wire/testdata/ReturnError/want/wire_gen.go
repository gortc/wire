// Code generated by Wire. DO NOT EDIT.

//go:generate go run gortc.io/wire/cmd/wire
//+build !wireinject

package main

// Injectors from wire.go:

func injectFoo() (Foo, error) {
	foo, err := provideFoo()
	if err != nil {
		return 0, err
	}
	return foo, nil
}
