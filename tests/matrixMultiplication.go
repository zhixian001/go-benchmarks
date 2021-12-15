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

func multiplyIntMatrixSize100(left, right *[100][100]int) *[100][100]int {
	var result [100][100]int

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			for k := 0; k < 100; k++ {
				result[i][j] += (*left)[i][k] * (*right)[k][j]
			}
		}
	}

	return &result
}

func multiplyFloatMatrixSize100(left, right *[100][100]float64) *[100][100]float64 {
	var result [100][100]float64

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			for k := 0; k < 100; k++ {
				result[i][j] += (*left)[i][k] * (*right)[k][j]
			}
		}
	}

	return &result
}

func randIntMatrixSize100() *[100][100]int {
	var result [100][100]int

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			result[i][j] = rand.Int()
		}
	}

	return &result
}

func randFloatMatrixSize100() *[100][100]float64 {
	var result [100][100]float64

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
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

	tests["intFixed100Matrix"] = &tester.TestEntry{
		Function: func() {
			left := randIntMatrixSize100()
			right := randIntMatrixSize100()

			multiplyIntMatrixSize100(left, right)
		},
	}

	tests["floatFixed100Matrix"] = &tester.TestEntry{
		Function: func() {
			left := randFloatMatrixSize100()
			right := randFloatMatrixSize100()

			multiplyFloatMatrixSize100(left, right)
		},
	}

	tester.BenchmarkFunctions(&tests, 1000, true, true)
}
