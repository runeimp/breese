package main

import (
	"fmt"

	"github.com/runeimp/breese"
)

const Version = "0.1.0-alpha"

func main() {
	fmt.Printf("Breese CLI v%s\n", Version)
	fmt.Printf("Breese Lib v%s\n", breese.Version)

	fmt.Println()

	ituChars := "abcdefghijklmnopqrstuvwxyz0123456789áåéñ"
	gerkeChars := "äöü"

	fmt.Println("International")
	for _, r := range ituChars {
		c := string(r)
		fmt.Printf("%c: %v\n", r, breese.MorseCodeInternational[c].ValueUnicode)
	}
	fmt.Println()

	fmt.Println("Continental")
	for _, r := range gerkeChars {
		c := string(r)
		fmt.Printf("%c: %v\n", r, breese.MorseCodeContinental[c].ValueUnicode)
	}
	// for k, v := range breese.MorseCodeInternational {
	// 	fmt.Printf("%q: %v\n", k, v.ValueUnicode)
	// }

	fmt.Println()
}
