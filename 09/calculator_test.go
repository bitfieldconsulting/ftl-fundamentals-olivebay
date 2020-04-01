package calculator_test

import (
	"testing"
	"calculator"
)

func TestSubtract(t *testing.T) {
	testCases := []struct {
		a, b int
		want int
	}{
		{a: 1, b: 1, want: 0},
		{a: 2, b: 1, want: 1},
		{a: 6, b: 7, want: -1},
		{a: 7, b: 6, want: 1},
		{a: -6, b: -7, want: 1},
		{a: 6, b: -7, want: 13},
		{a: 0, b: -7, want: 7},
		{a: 0, b: 0, want: 0},
		{a: -1, b: 0, want: -1},
		{a: -1, b: 7, want: -8},
	}
	for _, testCase := range testCases {
		got := calculator.Subtract(testCase.a, testCase.b)
		if testCase.want != got {
			t.Errorf("Subtract(%d, %d): want %d, got %d", testCase.a, testCase.b, testCase.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	testCases := []struct {
		a, b int
		want int
	}{
		{a: 1, b: 1, want: 1},
		{a: 2, b: 1, want: 2},
		{a: 6, b: 7, want: 42},
		{a: -6, b: -7, want: 42},
		{a: 6, b: -7, want: -42},
		{a: 0, b: -7, want: 0},
		{a: 0, b: 0, want: 0},
		{a: -1, b: 0, want: 0},
		{a: -1, b: 7, want: -7},
	}
	for _, testCase := range testCases {
		got := calculator.Multiply(testCase.a, testCase.b)
		if testCase.want != got {
			t.Errorf("Multiply(%d, %d): want %d, got %d", testCase.a, testCase.b, testCase.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	testCases := []struct {
		a, b int
		want int
	}{
		{a: 1, b: 1, want: 1},
		{a: 2, b: 1, want: 2},
		{a: 6, b: -7, want: 0},
		{a: -1, b: 7, want: 0},
		
	}
	for _, testCase := range testCases {
		got := calculator.Divide(testCase.a, testCase.b)
		if testCase.want != got {
			t.Errorf("Divide(%d, %d): want %d, got %d", testCase.a, testCase.b, testCase.want, got)
		}
	}
}
