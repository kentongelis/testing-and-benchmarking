package main

import (
	"fmt"
	"testing"
)

// Simple Test
func TestCalculate(t *testing.T) {
	fmt.Println("Test Calculate")
	expected := 4
	result := Calculate(2)
	if expected != result {
		t.Error("Failed")
	}
}

// Table Driven Testing
func TestTableCalculate(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{-5, -3},
		{99999, 100001},
	}

	for _, test := range tests {
		if output := Calculate(test.input); output != test.expected {
			t.Errorf("Test Failed: %d inputted, %d expected, received: %d", test.input, test.expected, output)
		}
	}
}

// Subtests with t.Run()
func TestCalculateSubtests(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"positive", 2, 4},
		{"negative", -1, 1},
		{"zero", 0, 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if output := Calculate(test.input); output != test.expected {
				t.Errorf("expected %d, got %d", test.expected, output)
			}
		})
	}
}

// Parallel Test Execution
func TestCalculateParallel(t *testing.T) {
	t.Parallel()
	if Calculate(2) != 4 {
		t.Error("Expected 2 + 2 to equal 4")
	}
}

// Fuzzing Tests
func FuzzCalculate(f *testing.F) {
	f.Add(1)
	f.Add(0)
	f.Add(-5)

	f.Fuzz(func(t *testing.T, input int) {
		result := Calculate(input)
		if result != input+2 {
			t.Errorf("calculate(%d) = %d, wnat %d", input, result, input+2)
		}
	})
}

// Use the -cover flag to see how much of your code your test covers

// You can visualisze this by running go test --coverprofile=coverage.out

// Then generate an html file for this by running go tool cover -html=coverage.out

// Simple benchmark test
func BenchmarkCalculate(b *testing.B) {
	for b.Loop() {
		Calculate(2)
	}
}

// run go test -bench=

func TestOther(t *testing.T) {
	fmt.Println("Testing something else")
	fmt.Println("This shoudn't run with -run=calc")
}

// If we wanted to run our Calculate tests, run go test -run=Calculate -bench=.

// If you just wanted to run the Benchmark Calculate function that run go test -run=Bench -bench=.

func benchmarkCalculate(input int, b *testing.B) {
	for b.Loop() {
		Calculate(input)
	}
}

func BenchmarkCalcualte100(b *testing.B)         { benchmarkCalculate(100, b) }
func BenchmarkCalculateNegative100(b *testing.B) { benchmarkCalculate(-100, b) }
func BenchmarkCalculateNegative1(b *testing.B)   { benchmarkCalculate(-1, b) }
