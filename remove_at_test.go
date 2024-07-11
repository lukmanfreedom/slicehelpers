package slicehelpers

import "testing"

func TestRemoveAt(t *testing.T) {
	intSlice := []int{1, 2, 3, 4, 5}
	strSlice := []string{"a", "b", "c"}
	floatSlice := []float64{1.1, 2.2, 3.3}

	removedInt := RemoveAt(intSlice, 2)
	expectedInt := []int{1, 2, 4, 5}

	for i, v := range removedInt {
		if v != expectedInt[i] {
			t.Errorf("Expected %v, got %v", expectedInt, removedInt)
		}
	}

	removedStr := RemoveAt(strSlice, 1)
	expectedStr := []string{"a", "c"}

	for i, v := range removedStr {
		if v != expectedStr[i] {
			t.Errorf("Expected %v, got %v", expectedStr, removedStr)
		}
	}

	removedFloat := RemoveAt(floatSlice, 1)
	expectedFloat := []float64{1.1, 3.3}

	for i, v := range removedFloat {
		if v != expectedFloat[i] {
			t.Errorf("Expected %v, got %v", expectedFloat, removedFloat)
		}
	}
}
