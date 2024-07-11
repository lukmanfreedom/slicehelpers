package slicehelpers

import "testing"

func TestReduce(t *testing.T) {
	intSlice := []int{1, 2, 3, 4, 5}
	strSlice := []string{"a", "b", "c"}
	floatSlice := []float64{1.1, 2.2, 3.3}

	// Test for int slice
	expectedIntSum := 15
	intSum := Reduce(intSlice, func(acc, v int) int {
		return acc + v
	}, 0)
	if intSum != expectedIntSum {
		t.Errorf("Expected int sum %d, got %d", expectedIntSum, intSum)
	}

	expectedIntProduct := 120
	intProduct := Reduce(intSlice, func(acc, v int) int {
		return acc * v
	}, 1)
	if intProduct != expectedIntProduct {
		t.Errorf("Expected int product %d, got %d", expectedIntProduct, intProduct)
	}

	// Test for string slice
	expectedConcat := "abc"
	concatenated := Reduce(strSlice, func(acc, v string) string {
		return acc + v
	}, "")
	if concatenated != expectedConcat {
		t.Errorf("Expected concatenated string %s, got %s", expectedConcat, concatenated)
	}

	// Test for float slice
	expectedFloatSum := 6.6
	floatSum := Reduce(floatSlice, func(acc, v float64) float64 {
		return acc + v
	}, 0.0)
	if floatSum != expectedFloatSum {
		t.Errorf("Expected float sum %f, got %f", expectedFloatSum, floatSum)
	}
}
