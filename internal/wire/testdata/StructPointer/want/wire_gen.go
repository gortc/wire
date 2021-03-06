// Code generated by Wire. DO NOT EDIT.

//go:generate go run gortc.io/wire/cmd/wire
//+build !wireinject

package main

// Injectors from wire.go:

func injectFooBar() *FooBar {
	foo := provideFoo()
	bar := provideBar()
	fooBar := &FooBar{
		Foo: foo,
		Bar: bar,
	}
	return fooBar
}

func injectEmptyStruct() *Empty {
	empty := &Empty{}
	return empty
}
