package utility

import "fmt"

func SwitchCase() {
	// var day string
	// fmt.Println("Enter the day: ")
	// fmt.Scan(&day)

	// switch day {
	// case "Sat", "Sun":
	// 	fmt.Println("Holiday!!!")
	// default:
	// 	fmt.Println("Working Day")
	// }

	var marks int
	fmt.Println("Enter marks: ")
	fmt.Scan(&marks)

	switch {
	case marks >= 50:
		fmt.Println("Passed!!")
	default:
		fmt.Println("Failed!!")
	}
}
