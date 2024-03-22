package lists

import "math/rand"

type SequenceGenerator func(int) []int

func GenerateRandomSequence(n int) []int {
	sequence := make([]int, n)

	for i := 0; i < n; i++ {
		sequence[i] = rand.Intn(n * 5)
	}

	return sequence
}

func GenerateIncreasingSequence(n int) []int {
	sequence := make([]int, n)
	for i := 0; i < n; i++ {
		sequence[i] = i + 1
	}
	return sequence
}

func GenerateDecreasingSequence(n int) []int {
	sequence := make([]int, n)
	for i := 0; i < n; i++ {
		sequence[i] = n - i
	}
	return sequence
}

func GenerateAShapeSequence(n int) []int {
	sequence := make([]int, n)
	mid := (n - 1) / 2

	for i := 0; i <= mid; i++ {
		sequence[i] = i + 1
	}

	for i := mid + 1; i < n; i++ {
		sequence[i] = sequence[i-1] - 1
	}

	return sequence
}
func GenerateVShapeSequence(n int) []int {
	sequence := make([]int, n)
	mid := (n - 1) / 2

	for i := 0; i <= mid; i++ {
		sequence[i] = mid - i + 1
	}

	for i := mid + 1; i < n; i++ {
		sequence[i] = sequence[i-1] + 1
	}

	return sequence
}

func GenerateSequenceFromGenerator(generator SequenceGenerator, n int) [][]int {
	sequences := make([][]int, 10)
	for i := range 10 {
		sequences[i] = generator(n)
	}

	return sequences
}
