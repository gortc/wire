// Code generated by Wire. DO NOT EDIT.

//go:generate go run gortc.io/wire/cmd/wire
//+build !wireinject

package main

// Injectors from wire.go:

func injectFooBar() FooBar {
	foo := provideFoo()
	bar := provideBar()
	fooBar := provideFooBar(foo, bar)
	return fooBar
}
