package main

import (
	"fmt"
	lib "github.com/padamstx/gomodlib/v3"
)

const version string = "2.0.0"

func main() {
	fmt.Printf("nomodapp main, version %s\n", version)
	fmt.Printf("   message from gomodlib: %s\n", lib.HelloWorld())
}