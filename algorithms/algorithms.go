package algorithms

import (
	"time"
)

type AlgorithmFunc func([]int) ([]int, int, int, int64)

func BubbleSort(arr []int) ([]int, int, int, int64) {
	comparisons := 0
	assignments := 0
	startTime := time.Now()

	n := len(arr)
	for i := range n {
		for j := range n - i - 1 {
			comparisons += 1
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				assignments += 1
			}
		}
	}

	executionTime := time.Since(startTime)

	return arr, comparisons, assignments, executionTime.Nanoseconds()
}

func MergeSort(arr []int) ([]int, int, int, int64) {
	comparisons := 0

	merge := func(left, right []int) []int {
		merged := []int{}

		for len(left) > 0 && len(right) > 0 {
			comparisons += 1
			if left[0] >= right[0] {
				merged = append(merged, left[0])
				left = append(left[:0], left[1:]...)
			} else {
				merged = append(merged, right[0])
				right = append(right[:0], right[1:]...)
			}
		}

		merged = append(merged, left...)
		merged = append(merged, right...)

		return merged
	}

	var sort func(arr []int) []int
	sort = func(arr []int) []int {
		comparisons += 1
		if len(arr) <= 1 {
			return arr
		}

		mid := len(arr) / 2
		left := arr[:mid]
		right := arr[mid:]

		return merge(sort(left), sort(right))

	}

	startTime := time.Now()
	sortedArr := sort(arr)
	executionTime := time.Since(startTime).Nanoseconds()

	return sortedArr, comparisons, 0, executionTime
}

func HeapSort(arr []int) ([]int, int, int, int64) {
	comparisons := 0
	assgnments := 0
	startTime := time.Now()

	n := len(arr)

	var heapify func([]int, int, int)
	heapify = func(arr []int, n, i int) {
		tmin := i
		left := 2*i + 1
		right := 2*i + 2
		comparisons += 3

		if (left < n) && (arr[left] < arr[i]) {
			tmin = left
		}
		if (right < n) && (arr[right] < arr[tmin]) {
			tmin = right
		}
		if tmin != i {
			arr[i], arr[tmin] = arr[tmin], arr[i]
			assgnments += 1
			heapify(arr, n, tmin)
		}
	}

	for i := n / 2; i > -1; i-- {
		heapify(arr, n, i)
	}

	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		assgnments += 1
		heapify(arr, i, 0)
	}

	executionTime := time.Since(startTime).Nanoseconds()

	return arr, comparisons, assgnments, executionTime

}

func SelectionSort(arr []int) ([]int, int, int, int64) {
	comparisons := 0
	assignments := 0
	startTime := time.Now()

	n := len(arr)
	for i := range n {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			comparisons += 1
			if arr[j] > arr[maxIdx] {
				maxIdx = j
			}
		}
		arr[i], arr[maxIdx] = arr[maxIdx], arr[i]
		assignments += 1
	}

	executionTime := time.Since(startTime).Nanoseconds()

	return arr, comparisons, assignments, executionTime
}

func InsertionSort(arr []int) ([]int, int, int, int64) {
	comparisons := 0
	assignments := 0
	startTime := time.Now()

	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && key > arr[j] {
			comparisons++
			assignments++
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
		assignments++
	}

	executionTime := time.Since(startTime).Nanoseconds()

	return arr, comparisons, assignments, executionTime
}

func QuickSort(arr []int) ([]int, int, int, int64) {
	comparisons := 0
	assignments := 0

	partition := func(arr []int, low, high int) int {
		i := low - 1
		pivot := arr[high]

		for j := low; j < high; j++ {
			comparisons++
			if arr[j] >= pivot {
				i += 1
				arr[i], arr[j] = arr[j], arr[i]
				assignments++
			}
		}

		arr[i+1], arr[high] = arr[high], arr[i+1]
		assignments++

		return i + 1
	}

	var helper func([]int, int, int)
	helper = func(arr []int, low, high int) {
		comparisons++
		if low < high {
			pi := partition(arr, low, high)
			helper(arr, low, pi-1)
			helper(arr, pi+1, high)
		}
	}

	startTime := time.Now()
	helper(arr, 0, len(arr)-1)
	executionTime := time.Since(startTime).Nanoseconds()

	return arr, comparisons, assignments, executionTime
}

func TestAlgorithm(algorithm AlgorithmFunc, sequences [][]int, n int) (int64, int) {

	var avgTime int64 = 0
	avgOperations := 0

	for _, seq := range sequences {
		copiedSeq := make([]int, n)
		copy(copiedSeq, seq)

		_, comparisons, assignments, timeTaken := algorithm(copiedSeq)

		avgTime += timeTaken
		avgOperations += comparisons + assignments
	}

	avgTime /= 10
	avgOperations /= 10

	return avgTime, avgOperations
}
