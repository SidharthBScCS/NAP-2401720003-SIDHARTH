package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20
	fmt.Println("Addition: ", a, "+", b, "=", a+b)
	fmt.Println("Subtraction: ", a, "-", b, "=", a-b)
	fmt.Println("Multiplication: ", a, "*", b, "=", a*b)
	fmt.Println("Division: ", a, "/", b, "=", a/b)

	var x float32 = 3.45
	var y float32 = 4.45
	fmt.Println("Addition: ", x, "+", y, "=", x+y)
	fmt.Println("Subtraction: ", x, "-", y, "=", x-y)
	fmt.Println("Multiplication: ", x, "*", y, "=", x*y)
	fmt.Println("Division: ", x, "/", y, "=", x/y)
}
