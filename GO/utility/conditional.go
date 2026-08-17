package utility

import "fmt"

func Conditional() {
	var marks int
	fmt.Print("Enter marks: ")
	fmt.Scan(&marks)

	if marks >= 40 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}
}
