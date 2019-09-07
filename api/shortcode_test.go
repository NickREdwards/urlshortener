package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestNewCodeIsGeneratedOnEachInvocation(t *testing.T) {
	// Arrange
	rand.Seed(time.Now().UTC().UnixNano())
	codes := make(map[string]int)
	iterations := 100000
	codeLength := 6
	duplicate := false

	// Act
	for i := 0; i < iterations; i++ {
		code := NewShortCode(codeLength)
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
	rand.Seed(time.Now().UTC().UnixNano())
	NewShortCode(6)
}
