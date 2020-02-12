package main

import (
	"fmt"

	"github.com/runeimp/breese"
)

const Version = "0.1.0-alpha"

func main() {
	fmt.Printf("Breese CLI v%s\n", Version)
	fmt.Printf("Breese Lib v%s\n", breese.Version)
}
