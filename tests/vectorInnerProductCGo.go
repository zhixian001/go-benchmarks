package tests

// #cgo CFLAGS: -Wall -mavx
// #include "vectorInnerProductCGo.h"
import "C"

import (
	"benchmark/tester"
	"fmt"
	"unsafe"
)

func multiplyFixed4VectorCGo(left, right *[4]float64) float64 {
	return float64(
		C.innerProductFixed4Vector(
			(*[4]C.double)(unsafe.Pointer(left)),
			(*[4]C.double)(unsafe.Pointer(right)),
		),
	)
}

// for debugging cgo simd function
func testIntrinsic() {
	left := randFloatVectorFixed4()
	right := randFloatVectorFixed4()

	fmt.Println(left)
	fmt.Println(right)

	ret := multiplyFixed4VectorCGo(left, right)

	fmt.Println(ret)
}

func RunVectorInnerProductCGoAVXTest() {
	tests := make(map[string]*tester.TestEntry)

	tests["innerProductFloatVector4CGoAVX"] = &tester.TestEntry{
		Function: func() {
			left := randFloatVectorFixed4()
			right := randFloatVectorFixed4()

			multiplyFixed4VectorCGo(left, right)
		},
	}

	tester.BenchmarkFunctions(&tests, 1000, true, true)
}
