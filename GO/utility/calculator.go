package utility

import "fmt"

func Calculator() {
	var a int
	fmt.Print("Enter first integer: ")
	fmt.Scan(&a)

	var b int
	fmt.Print("Enter second integer: ")
	fmt.Scan(&b)

	var choice int
	fmt.Println("Enter Choice: ")
	fmt.Println("1.Addition")
	fmt.Println("2.Subtration")
	fmt.Println("3.Division")
	fmt.Println("4.Multiplication")
	fmt.Scan(&choice)

	if choice == 1 {
		var addition int = a + b
		fmt.Printf("Addition %d \n", addition)
	} else if choice == 2 {
		var subtraction int = a - b
		fmt.Printf("Subtraction %d \n", subtraction)
	} else if choice == 3 {
		if a == 0 {
			fmt.Println("Division not possible with 0")
		}
		if b == 0 {
			fmt.Println("Division is not possible with 0")
		}
		var division int = a / b
		fmt.Printf("Division %d \n", division)
	} else if choice == 4 {
		var multiplication int = a * b
		fmt.Printf("Multiplication %d \n", multiplication)
	} else {
		fmt.Println("Invalid error")
	}

	var x float32
	fmt.Print("Enter first float: ")
	fmt.Scan(&x)

	var y float32
	fmt.Print("Enter second float: ")
	fmt.Scan(&y)

	var choice2 int
	fmt.Println("Enter Choice: ")
	fmt.Println("1.Addition")
	fmt.Println("2.Subtration")
	fmt.Println("3.Division")
	fmt.Println("4.Multiplication")
	fmt.Scan(&choice2)

	if choice2 == 1 {
		var add float32 = x + y
		fmt.Printf("Addition %f \n", add)
	} else if choice2 == 2 {
		var sub float32 = x - y
		fmt.Printf("Subtraction %f \n", sub)
	} else if choice2 == 3 {
		if x == 0 {
			fmt.Println("Division not possible with 0")
		}
		if y == 0 {
			fmt.Println("Division is not possible with 0")
		}
		var div float32 = x / y
		fmt.Printf("Division %f \n", div)
	} else if choice == 4 {
		var multi float32 = x * y
		fmt.Printf("Multiplication %f \n", multi)
	} else {
		fmt.Println("Invalid error")
	}

}
