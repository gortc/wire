// Code generated by Wire. DO NOT EDIT.

//go:generate go run gortc.io/wire/cmd/wire
//+build !wireinject

package main

import (
	"io"
	"os"
)

// Injectors from wire.go:

func inject() io.Writer {
	file := _wireFileValue
	return file
}

var (
	_wireFileValue = os.Stdout
)
