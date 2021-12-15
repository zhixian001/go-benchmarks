package tests

import (
	"benchmark/tester"
	"math/rand"
)

func multiplyIntMatrix(size int, left, right *[][]int) *[][]int {
	result := make([][]int, size)

	for i := 0; i < size; i++ {
		result[i] = make([]int, size)
		for j := 0; j < size; j++ {
			for k := 0; k < size; k++ {
				result[i][j] += (*left)[i][k] * (*right)[k][j]
			}
		}
	}

	return &result
}

func multiplyFloatMatrix(size int, left, right *[][]float64) *[][]float64 {
	result := make([][]float64, size)

	for i := 0; i < size; i++ {
		result[i] = make([]float64, size)
		for j := 0; j < size; j++ {
			for k := 0; k < size; k++ {
				result[i][j] += (*left)[i][k] * (*right)[k][j]
			}
		}
	}

	return &result
}

func randIntMatrix(size int) *[][]int {
	result := make([][]int, size)

	for i := range result {
		result[i] = make([]int, size)

		for j := range result {
			result[i][j] = rand.Int()
		}
	}

	return &result
}

func randFloatMatrix(size int) *[][]float64 {
	result := make([][]float64, size)

	for i := range result {
		result[i] = make([]float64, size)

		for j := range result {
			result[i][j] = rand.Float64()
		}
	}

	return &result
}

func RunMatrixMultiplicationTest() {
	tests := make(map[string]*tester.TestEntry)

	matrixSize := 100

	tests["intMatrix"] = &tester.TestEntry{
		Function: func() {
			left := randIntMatrix(matrixSize)
			right := randIntMatrix(matrixSize)

			multiplyIntMatrix(matrixSize, left, right)
		},
	}

	tests["floatMatrix"] = &tester.TestEntry{
		Function: func() {
			left := randFloatMatrix(matrixSize)
			right := randFloatMatrix(matrixSize)

			multiplyFloatMatrix(matrixSize, left, right)
		},
	}

	tester.BenchmarkFunctions(&tests, 1000, true, true)
}
