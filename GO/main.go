package main

import (
	"fmt"
	"newageprogramming/utility"
)

func main() {
	var temp utility.Celsius = 35.67
	fmt.Print(temp.ToFarehheit())
}
