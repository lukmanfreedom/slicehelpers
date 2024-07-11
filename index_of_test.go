package slicehelpers

import "testing"

func TestIndexOf(t *testing.T) {
	intSlice := []int{1, 2, 3, 4, 5}
	strSlice := []string{"a", "b", "c"}
	floatSlice := []float64{1.1, 2.2, 3.3}

	if idx := IndexOf(intSlice, 3); idx != 2 {
		t.Errorf("Expected index of 3 to be 2, got %d", idx)
	}

	if idx := IndexOf(intSlice, 6); idx != -1 {
		t.Errorf("Expected index of 6 to be -1, got %d", idx)
	}

	if idx := IndexOf(strSlice, "b"); idx != 1 {
		t.Errorf("Expected index of 'b' to be 1, got %d", idx)
	}

	if idx := IndexOf(strSlice, "d"); idx != -1 {
		t.Errorf("Expected index of 'd' to be -1, got %d", idx)
	}

	if idx := IndexOf(floatSlice, 2.2); idx != 1 {
		t.Errorf("Expected index of 2.2 to be 1, got %d", idx)
	}

	if idx := IndexOf(floatSlice, 4.4); idx != -1 {
		t.Errorf("Expected index of 4.4 to be -1, got %d", idx)
	}
}
