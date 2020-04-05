package calculator

import "errors"

func Subtract(x int, y int) int {
	return x - y
}

func Multiply(x int, y int) int {
	return x * y
}

func Divide(x int, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("division by zero")
	}
	return (x / y), nil
}

func Add(x int, y int) int {
	return x + y
}
