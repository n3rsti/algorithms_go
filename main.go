package main

import (
	"fmt"
	"main/algorithms"
	"main/lists"
)

type Algorithm struct {
	Name string
	Func algorithms.AlgorithmFunc
}

type Sequence struct {
	Name      string
	Generator lists.SequenceGenerator
}

var algorithmsList = []Algorithm{
	{"Quick sort", algorithms.QuickSort},
	{"Merge sort", algorithms.MergeSort},
	{"Heap sort", algorithms.HeapSort},
	{"Bubble sort", algorithms.BubbleSort},
	{"Insertion sort", algorithms.InsertionSort},
	{"Selection sort", algorithms.SelectionSort},
}

var sequenceGenerators = []Sequence{
	{"Random", lists.GenerateRandomSequence},
	{"Increasing", lists.GenerateIncreasingSequence},
	{"Decreasing", lists.GenerateDecreasingSequence},
	{"A-Shape", lists.GenerateAShapeSequence},
	{"V-Shape", lists.GenerateVShapeSequence},
}

var nList = []int{1000, 2000, 3000, 5000, 8000, 10000, 15000, 20000, 30000, 50000}

func PrintTableHeader() {
	fmt.Print("Size;")
	for _, alg := range algorithmsList {
		fmt.Printf("%s;", alg.Name)
	}
	fmt.Println()
}

func PrintTableData[T int | int64](sequenceName string, data [][]T) {
	for i, sizeData := range data {
		fmt.Print(nList[i], ";")
		for _, val := range sizeData {
			fmt.Printf("%d;", val)
		}
		fmt.Println()
	}
}

func main() {
	for _, sequence := range sequenceGenerators {
		fmt.Printf("Results for %s sequence:\n", sequence.Name)

		timeResults := make([][]int64, len(nList))

		operationsResults := make([][]int, len(nList))

		for i, n := range nList {
			sequences := lists.GenerateSequenceFromGenerator(sequence.Generator, n)
			timeResults[i] = make([]int64, len(algorithmsList))
			operationsResults[i] = make([]int, len(algorithmsList))

			for j, algorithm := range algorithmsList {
				avgTime, avgOperations := algorithms.TestAlgorithm(algorithm.Func, sequences, n)
				timeResults[i][j] = avgTime
				operationsResults[i][j] = avgOperations
			}
		}

		PrintTableHeader()
		PrintTableData(sequence.Name, timeResults)

		PrintTableHeader()
		PrintTableData(sequence.Name, operationsResults)
	}
}
