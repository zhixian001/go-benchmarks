package tests

import (
	"benchmark/tester"
)

func innerProductIntVectorFixed4(left, right *[4]int) int {
	result := 0

	for i := 0; i < 4; i++ {
		result += (*left)[i] * (*right)[i]
	}

	return result
}

func innerProductFloatVectorFixed4(left, right *[4]float64) float64 {
	result := 0.0

	for i := 0; i < 4; i++ {
		result += (*left)[i] * (*right)[i]
	}

	return result
}

func RunVectorInnerProductTest() {
	tests := make(map[string]*tester.TestEntry)

	tests["innerProductIntVector4"] = &tester.TestEntry{
		Function: func() {
			left := randIntVectorFixed4()
			right := randIntVectorFixed4()

			innerProductIntVectorFixed4(left, right)
		},
	}

	tests["innerProductFloatVector4"] = &tester.TestEntry{
		Function: func() {
			left := randFloatVectorFixed4()
			right := randFloatVectorFixed4()

			innerProductFloatVectorFixed4(left, right)
		},
	}

	tester.BenchmarkFunctions(&tests, 1000, true, true)
}
