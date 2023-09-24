package Strategy

import (
	"testing"
)

func TestRadixSort_Sort(t *testing.T) {
	// Test with a large dataset
	arrLarge := []int64{170, 45, 75, 90, 802, 24, 2, 66, 1000, 98765, 4321, 99999, 123}
	sorterLarge := NewRadixSort(arrLarge)
	sorterLarge.Sort()
	sortedLarge := sorterLarge.SortedArray()

	expectedLarge := []int64{2, 24, 45, 66, 75, 90, 123, 170, 802, 1000, 4321, 98765, 99999}
	if !equal(sortedLarge, expectedLarge) {
		t.Errorf("RadixSort failed (Large dataset): got %v, expected %v", sortedLarge, expectedLarge)
	}
}

func TestInsertionSort_Sort(t *testing.T) {
	// Test with a small dataset
	arrSmall := []int64{12, 11, 13, 5, 6}
	sorterSmall := NewInsertionSort(arrSmall)
	sorterSmall.Sort()
	sortedSmall := sorterSmall.SortedArray()

	expectedSmall := []int64{5, 6, 11, 12, 13}
	if !equal(sortedSmall, expectedSmall) {
		t.Errorf("InsertionSort failed (Small dataset): got %v, expected %v", sortedSmall, expectedSmall)
	}
}

func equal(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
