package main

import (
	"fmt"
	"runtime"

	"goLangTraining/003-packages/dog"
)

func main() {
	fmt.Printf("runtime: %s\narchitecture: %s\n", runtime.GOOS, runtime.GOARCH)
	dog.Sit()
	dog.Bark()
}
