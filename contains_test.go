package slicehelpers

import "testing"

func TestContains(t *testing.T) {
	intSlice := []int{1, 2, 3, 4, 5}
	strSlice := []string{"a", "b", "c"}
	floatSlice := []float64{1.1, 2.2, 3.3}

	if !Contains(intSlice, 3) {
		t.Errorf("Expected intSlice to contain 3")
	}

	if Contains(intSlice, 6) {
		t.Errorf("Expected intSlice to not contain 6")
	}

	if !Contains(strSlice, "b") {
		t.Errorf("Expected strSlice to contain 'b'")
	}

	if Contains(strSlice, "d") {
		t.Errorf("Expected strSlice to not contain 'd'")
	}

	if !Contains(floatSlice, 2.2) {
		t.Errorf("Expected floatSlice to contain 2.2")
	}

	if Contains(floatSlice, 4.4) {
		t.Errorf("Expected floatSlice to not contain 4.4")
	}
}
