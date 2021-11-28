package main

import (
	"fmt"
	seb "github.com/superseb/go-libs"
)

func main() {
	string := seb.EpochString("yihaa")
	fmt.Println(string)
	// Yayayaya
	string2 := seb.EpochString("yihaa2")
	fmt.Println(string2)
	// Yayayaya
	// etc...
}
