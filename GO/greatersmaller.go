package main

import "fmt"

func greatersmaller() {
	var a int = 20
	var b int = 22

	if a > b {
		fmt.Println(a, "is greater")
	} else {
		fmt.Println(b, "is greater")
	}
}
