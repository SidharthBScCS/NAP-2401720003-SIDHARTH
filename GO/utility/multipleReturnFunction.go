package utility

func MultipleReturnFunction(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	} else {
		return a / b, true
	}
}
