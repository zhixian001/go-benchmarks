# Go Benchmakrs

## Usage

```go
tests := make(map[string]*tester.TestEntry)

tests["Test A"] = &tester.TestEntry{
  Function: func() { myFuncA(argA, argB) },
}

tests["Test B"] = &tester.TestEntry{
  Function: func() { myFuncB(argA, argB) },
}

tester.BenchmarkFunctions(&tests, 100, true)
```

## Useful Commands

### Build with gccgo auto vectorization

```bash
go build -compiler=gccgo -gccgoflags="-O3 -march=native -ffast-math"
go build -compiler=gccgo -gccgoflags="-O3 -march=core-avx2 -ffast-math"
go build -compiler=gccgo -gccgoflags="-O3 -march=haswell -ffast-math"
```