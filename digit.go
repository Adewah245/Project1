package main

func digit(num1 int, signs string, num2 int) int {
	switch signs {
	case "+", "Add", "add", "1":
		return num1 + num2
	case "-", "minus", "Minus", "2":
		return num1 - num2
	case "*", "mulyiply", "Multiply", "3":
		return num1 * num2
	case "/", "Divide", "divide", "4":
		return num1 / num2
	default:
		return 0
	}
}
