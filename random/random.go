// Package random provides a centralized way to generate random numbers throughout the library.
// It allows users to customize the random source used by the library, which is useful for
// testing or when specific random number generation behavior is required.
//
// Usage:
//
//	// Use the default random source
//	n := rand.IntN(100)
//	f := rand.Float64()
//
//	// Set a custom random source
//	customSource := rand.NewPCG(seed1, seed2)
//	rand.SetRandomSource(customSource)
//
//	// Reset to the default random source
//	rand.ResetRandomSource()
package random

import (
	"math/rand/v2"
	"sync"
	"time"
)

var (
	// defaultSource is the default random source using package default randomness
	defaultSource = rand.NewPCG(uint64(time.Now().UnixNano()), uint64(time.Now().UnixNano()))

	// currentRandom points to the active random
	currentRandom = rand.New(defaultSource)

	// mu protects concurrent access to currentRandom
	mu sync.RWMutex
)

// SetRandomSource sets a custom random source for the library
func SetRandomSource(src rand.Source) {
	mu.Lock()
	defer mu.Unlock()
	currentRandom = rand.New(src)
}

// ResetRandomSource resets to the default random source
func ResetRandomSource() {
	mu.Lock()
	defer mu.Unlock()
	currentRandom = rand.New(defaultSource)
}

// IntN generates a random int in [0,n)
func IntN(n int) int {
	mu.RLock()
	defer mu.RUnlock()
	return currentRandom.IntN(n)
}

// Float64 generates a random float64 in [0.0,1.0)
func Float64() float64 {
	mu.RLock()
	defer mu.RUnlock()
	return currentRandom.Float64()
}

// Int64N generates a random int64 in [0,n)
func Int64N(n int64) int64 {
	mu.RLock()
	defer mu.RUnlock()
	return currentRandom.Int64N(n)
}
