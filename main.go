package main

import (
	"go/build"
	"os"
)

func main() {
	context := build.Default
	wd, _ := os.Getwd()
	_, err := context.Import("knative.dev/eventing/pkg/apis", wd, build.ImportComment)
	if err != nil {
		panic(err)
	}
}
