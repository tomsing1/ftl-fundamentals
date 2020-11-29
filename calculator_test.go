package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
		name string
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 4, name: "Two postive numbers which sum to a positive."},
		{a: 1, b: -1, want: 0, name: "One postive and one negative umbers which sum to zero."},
		{a: 5, b: 0, want: 5, name: "One postive number and zero sum to the positive number."},
	}
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s: Add(%f, %f): want %f, got %f",
				tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 4, b: 2, want: 2},
		{a: 4, b: -2, want: 6},
		{a: 0, b: -1, want: 1},
		{a: 0, b: 1, want: -1},
	}
	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Subtract(%f, %f): want %f, got %f",
				tc.a, tc.b, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 1, b: 1, want: 1},
		{a: 1, b: -1, want: -1},
		{a: -1, b: -1, want: 1},
		{a: 2, b: 2, want: 4},
	}
	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Multiply(%f, %f): want %f, got %f",
				tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b        float64
		want        float64
		errExpected bool
	}
	testCases := []testCase{
		{a: 1, b: 1, want: 1, errExpected: false},
		{a: 1, b: 0, want: 999, errExpected: true},
		{a: 4, b: 2, want: 2, errExpected: false},
		{a: 4, b: -2, want: -2, errExpected: false},
	}
	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		errReceived := err != nil
		if tc.errExpected != errReceived {
			t.Fatalf("Divide (%f, %f): unexpected error status: %v",
				tc.a, tc.b, errReceived)
		}
		if !errReceived && tc.want != got {
			t.Errorf("Divide(%f, %f): want %f, got %f",
				tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a           float64
		want        float64
		errExpected bool
	}
	testCases := []testCase{
		{a: 4, want: 2, errExpected: false},
		{a: 1, want: 1, errExpected: false},
		{a: 0, want: 0, errExpected: false},
		{a: -1, want: 999, errExpected: true},
	}
	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.a)
		errReceived := err != nil
		if tc.errExpected != errReceived {
			t.Fatalf("Sqrt (%f): unexpected error status: %v",
				tc.a, errReceived)
		}
		if !errReceived && tc.want != got {
			t.Errorf("Sqrt(%f): want %f, got %f",
				tc.a, tc.want, got)
		}
	}
}
