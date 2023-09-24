package Strategy

import "fmt"

// ISort defines the sorting strategy interface.
type ISort interface {
	Sort()
	SortedArray() []int64
}

// InsertionSort is a concrete sorting strategy using the insertion sort algorithm.
type InsertionSort struct {
	array []int64
}

func NewInsertionSort(arr []int64) ISort {
	return &InsertionSort{array: arr}
}

func (is *InsertionSort) SortedArray() []int64 {
	return is.array
}

func (is *InsertionSort) Sort() {
	var i, j int
	var key int64
	for i = 1; i < len(is.array); i++ {
		key = is.array[i]
		for j = i - 1; j >= 0 && is.array[j] > key; j-- {
			is.array[j+1] = is.array[j]
		}
		is.array[j+1] = key
	}
}

// RadixSort is a concrete sorting strategy using the radix sort algorithm.
type RadixSort struct {
	array []int64
}

func NewRadixSort(arr []int64) ISort {
	return &RadixSort{array: arr}
}

func (rs *RadixSort) SortedArray() []int64 {
	return rs.array
}

func (rs *RadixSort) getMax() int64 {
	max := rs.array[0]
	for _, num := range rs.array {
		if num > max {
			max = num
		}
	}
	return max
}

func (rs *RadixSort) countingSort(exp int64) {
	n := len(rs.array)
	output := make([]int64, n)
	count := make([]int64, 10)

	for i := 0; i < n; i++ {
		count[(rs.array[i]/exp)%10]++
	}

	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	for i := n - 1; i >= 0; i-- {
		output[count[(rs.array[i]/exp)%10]-1] = rs.array[i]
		count[(rs.array[i]/exp)%10]--
	}

	for i := 0; i < n; i++ {
		rs.array[i] = output[i]
	}
}

func (rs *RadixSort) Sort() {
	max := rs.getMax()
	for exp := int64(1); max/exp > 0; exp *= 10 {
		rs.countingSort(exp)
	}
}

// Sorter is a context that uses a sorting strategy.
type Sorter struct {
	sortingAlgo ISort
}

func (s *Sorter) setSorter(sortAlg ISort) {
	s.sortingAlgo = sortAlg
}

func (s *Sorter) PrintSortedArray() {
	s.sortingAlgo.Sort()
	fmt.Printf("%v", s.sortingAlgo.SortedArray())
}
