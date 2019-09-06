package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestMain(t *testing.M) {
	rand.Seed(time.Now().UTC().UnixNano())
}

func TestNewCodeIsGeneratedOnEachInvocation(t *testing.T) {
	// Arrange
	codes := make(map[string]int)
	iterations := 1000000
	duplicate := false

	// Act
	for i := 0; i < iterations; i++ {
		code := NewShortCode()
		if _, ok := codes[code]; ok {
			duplicate = true
			fmt.Printf("\n%v\n", code)
			//break
			codes[code]++
		} else {
			codes[code] = 1
		}
	}

	// Assert
	if duplicate {
		t.Errorf("Non-unique code was generated")
	}
}

func BenchmarkNewShortCode(b *testing.B) {
	NewShortCode()
}
