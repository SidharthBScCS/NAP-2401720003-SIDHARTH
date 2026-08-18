package utility

import "fmt"

func Calculator() {

	for {
		var choice int
		fmt.Println("1.Perform integer operations")
		fmt.Println("2.Perform float operations")
		fmt.Println("3.Exit")
		fmt.Scan(&choice)

		if choice == 1 {
			var a int
			fmt.Print("Enter first integer: ")
			fmt.Scan(&a)

			var b int
			fmt.Print("Enter second integer: ")
			fmt.Scan(&b)

			var addition int = a + b
			fmt.Printf("Addition %d \n", addition)

			var subtraction int = a - b
			fmt.Printf("Subtraction %d \n", subtraction)

			if a == 0 {
				fmt.Println("Division not possible with 0")
			} else if b == 0 {
				fmt.Println("Division is not possible with 0")
			}
			var division int = a / b
			fmt.Printf("Division %d \n", division)
		} else if choice == 2 {
			var x float32
			fmt.Print("Enter first float: ")
			fmt.Scan(&x)

			var y float32
			fmt.Print("Enter second float: ")
			fmt.Scan(&y)

			var add float32 = x + y
			fmt.Printf("Addition %f \n", add)

			var sub float32 = x - y
			fmt.Printf("Subtraction %f \n", sub)

			if x == 0 {
				fmt.Println("Division not possible with 0")
			} else if y == 0 {
				fmt.Println("Division is not possible with 0")
			} else {
				var div float32 = x / y
				fmt.Printf("Division %f \n", div)
			}
			var multi float32 = x * y
			fmt.Printf("Multiplication %f \n", multi)
		} else if choice == 3 {
			break
		} else {
			fmt.Println("Invalid error")
		}
	}

}
