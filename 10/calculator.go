package calculator

import (
	"fmt"
)

func Subtract(x int, y int) int {
	return x - y
}

func Multiply(x int, y int) int {
	return x * y
}

func Divide(x int, y int) (int, error) {
	if y == 0 {
		return 0, fmt.Errorf("%d / %d division by zero", x, y)
	}
	return (x / y), nil
}

func Add(x int, y int) int {
	return x + y
}
