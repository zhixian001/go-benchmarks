package main

import "benchmark/tests"

func main() {
	// tests.RunMatrixMultiplicationTest()
	// tests.RunVectorOuterProductTest()
	tests.RunVectorInnerProductTest()
	tests.RunVectorInnerProductCGoAVXTest()
}
