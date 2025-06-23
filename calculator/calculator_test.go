package calculator_test

import (
	"calculator"
	"math"
	"testing"
)

type testCase struct {
	a, b float64
	want float64
}

func closeEnough(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestAdd(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 2},
		{a: 5, b: 0, want: 5},
	}
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("add(%f,%f) = %f, want %f", tc.a, tc.b, got, tc.want)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 5, b: 3, want: 2},
		{a: 10, b: 5, want: 5},
		{a: 0, b: 0, want: 0},
		{a: 2, b: 2, want: 0},
		{a: -2, b: 2, want: -4},
		{a: 2, b: -2, want: 4},
		{a: -2, b: -2, want: 0},
	}
	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("subtract(%f,%f) = %f, want %f", tc.a, tc.b, got, tc.want)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 2, b: 3, want: 6},
		{a: 5, b: 0, want: 0},
		{a: 1, b: 7, want: 7},
		{a: -2, b: 3, want: -6},
		{a: -3, b: -4, want: 12},
	}
	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("multiply(%f,%f) = %f, want %f", tc.a, tc.b, got, tc.want)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 2, b: 2, want: 1},
		{a: -1, b: -1, want: 1},
		{a: 10, b: 5, want: 2},
		{a: 1, b: 3, want: 0.333333},
	}
	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		if err != nil {
			t.Fatalf("invalid input found, error: %v", err)
		}
		if !closeEnough(tc.want, got, 0.001) {
			t.Errorf("divide(%f,%f) = %f, want %f", tc.a, tc.b, got, tc.want)
		}
		// if tc.want != got {
		// 	t.Errorf("divide(%f,%f) = %f, want %f", tc.a, tc.b, got, tc.want)
		// }
	}
}

func TestInvalidDivide(t *testing.T) {
	t.Parallel()
	_, err := calculator.Divide(1, 0)
	if err == nil {
		t.Error("want error,got nil")
	}
}
