package utility

import "fmt"

func IfElseCondition() {
	var score int = 89

	if score >= 60 {
		fmt.Println("PASS")
	} else if score == 0 {
		fmt.Println("GET OUT")
	} else {
		fmt.Println("FAIL")
	}
}
