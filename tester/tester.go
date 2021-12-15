package tester

import (
	"fmt"
	"math"
	"time"
)

type TestEntry struct {
	// TODO: Prepare, PrepareArg, FunctionArg
	Function func()
}

func BenchmarkFunctions(tests *map[string]*TestEntry, trialCount int, ignoreFirstRun, hideEachDuration bool) {
	testResult := make(map[string][]time.Duration)

	if ignoreFirstRun {
		trialCount++
	}

	for testNameKey, entryToTest := range *tests {
		testResultDurations := &[]time.Duration{}

		for i := 0; i < trialCount; i++ {
			startTime := time.Now()

			entryToTest.Function()

			elapsedTime := time.Since(startTime)

			if ignoreFirstRun && i == 0 {
				continue
			}

			*testResultDurations = append(*testResultDurations, elapsedTime)
		}

		testResult[testNameKey] = *testResultDurations
	}

	fmt.Println("🧀 Test Result 🐭")
	for k, v := range testResult {
		funcName := fmt.Sprintf("> %s : ", k)
		testResults := ""
		testStats := "📊 "

		var sum time.Duration = 0
		var avg float64

		for _, duration := range v {
			testResults += fmt.Sprintf("%s\t", duration)

			sum += duration
		}

		avg = float64(sum) / float64(trialCount)

		avgDuration, _ := time.ParseDuration(fmt.Sprintf("%fns", avg))

		testStats += fmt.Sprintf("(avg: %s, ", avgDuration)

		var stddevHelper float64 = 0.0

		for _, duration := range v {
			stddevHelper += math.Pow((avg - float64(duration)), 2)
		}

		stdDev := math.Sqrt(stddevHelper / float64(trialCount))

		testStats += fmt.Sprintf("stddev: %f)", stdDev)

		outputText := funcName + testStats

		if !hideEachDuration {
			outputText += "\n" + testResults
		}

		fmt.Println(outputText)
	}
}
