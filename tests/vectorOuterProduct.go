package tests

import (
	"benchmark/tester"
	"math/rand"
)

func randIntVectorFixed4() *[4]int {
	result := [4]int{}

	for i := 0; i < 4; i++ {
		result[i] = rand.Int()
	}

	return &result
}

func randFloatVectorFixed4() *[4]float64 {
	result := [4]float64{}

	for i := 0; i < 4; i++ {
		result[i] = rand.Float64()
	}

	return &result
}

func outerProductIntVectorFixed4(left, right *[4]int) *[4][4]int {
	result := [4][4]int{}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			result[i][j] = (*left)[i] * (*right)[j]
		}
	}

	return &result
}

func outerProductFloatVectorFixed4(left, right *[4]float64) *[4][4]float64 {
	result := [4][4]float64{}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			result[i][j] = (*left)[i] * (*right)[j]
		}
	}

	return &result
}

func RunVectorOuterProductTest() {
	tests := make(map[string]*tester.TestEntry)

	tests["outerProductIntVector4"] = &tester.TestEntry{
		Function: func() {
			left := randIntVectorFixed4()
			right := randIntVectorFixed4()

			outerProductIntVectorFixed4(left, right)
		},
	}

	tests["outerProductFloatVector4"] = &tester.TestEntry{
		Function: func() {
			left := randFloatVectorFixed4()
			right := randFloatVectorFixed4()

			outerProductFloatVectorFixed4(left, right)
		},
	}

	tester.BenchmarkFunctions(&tests, 1000, true, true)
}
