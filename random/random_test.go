package random

import (
	"math/rand/v2"
	"testing"
)

// setupTest sets up a deterministic random source for testing
func setupTest(t *testing.T) {
	t.Helper()
	// Use a fixed seed for reproducible tests
	src := rand.NewPCG(42, 43)
	SetRandomSource(src)
}

// teardownTest resets the random source after testing
func teardownTest(t *testing.T) {
	t.Helper()
	ResetRandomSource()
}

// TestSetRandomSource tests the SetRandomSource function
func TestSetRandomSource(t *testing.T) {
	// Create a custom source with a known seed
	src1 := rand.NewPCG(100, 200)

	// Set the custom source
	SetRandomSource(src1)

	// Generate some values
	val1 := IntN(1000)
	val2 := Float64()

	// Create a new source with the same seed
	src2 := rand.NewPCG(100, 200)

	// Set the new source
	SetRandomSource(src2)

	// Generate values again
	val1Again := IntN(1000)
	val2Again := Float64()

	// Values should be the same with sources created with the same seed
	if val1 != val1Again {
		t.Errorf("SetRandomSource failed: IntN values differ: %v != %v", val1, val1Again)
	}

	if val2 != val2Again {
		t.Errorf("SetRandomSource failed: Float64 values differ: %v != %v", val2, val2Again)
	}

	// Reset for other tests
	ResetRandomSource()
}

// TestResetRandomSource tests the ResetRandomSource function
func TestResetRandomSource(t *testing.T) {
	// Since the default source is time-based, we'll test that ResetRandomSource
	// returns to a different source than a custom one, rather than testing exact values

	// Set a custom source with a known seed
	customSrc := rand.NewPCG(100, 200)
	SetRandomSource(customSrc)

	// Generate some values with custom source
	customVal1 := IntN(1000)
	customVal2 := Float64()

	// Reset to default
	ResetRandomSource()

	// Generate values again after reset
	resetVal1 := IntN(1000)
	resetVal2 := Float64()

	// Values should be different from custom source values
	// This is a probabilistic test, but with high probability the values will differ
	if customVal1 == resetVal1 && customVal2 == resetVal2 {
		t.Errorf("ResetRandomSource failed: values should differ after reset")
	}

	// Set the same custom source again
	SetRandomSource(customSrc)

	// Generate values again
	customVal1Again := IntN(1000)
	customVal2Again := Float64()

	// Values should be different from the first custom source values
	// because we've used the source in between
	if customVal1 == customVal1Again && customVal2 == customVal2Again {
		t.Errorf("ResetRandomSource failed: custom source should advance")
	}

	// Reset for other tests
	ResetRandomSource()
}

// TestIntN tests the IntN function
func TestIntN(t *testing.T) {
	setupTest(t)
	defer teardownTest(t)

	// Test multiple calls to ensure reproducibility
	val1 := IntN(100)
	val2 := IntN(100)

	// With a fixed seed, we should get consistent results
	expectedVal1 := 74
	expectedVal2 := 24

	if val1 != expectedVal1 {
		t.Errorf("IntN(100) = %v, want %v", val1, expectedVal1)
	}

	if val2 != expectedVal2 {
		t.Errorf("IntN(100) second call = %v, want %v", val2, expectedVal2)
	}

	// Test with different bounds
	val3 := IntN(1000)
	expectedVal3 := 741

	if val3 != expectedVal3 {
		t.Errorf("IntN(1000) = %v, want %v", val3, expectedVal3)
	}
}

// TestFloat64 tests the Float64 function
func TestFloat64(t *testing.T) {
	setupTest(t)
	defer teardownTest(t)

	// Test multiple calls to ensure reproducibility
	val1 := Float64()
	val2 := Float64()

	// With a fixed seed, we should get consistent results
	expectedVal1 := 0.40510544537896787
	expectedVal2 := 0.5178550802041214

	if val1 != expectedVal1 {
		t.Errorf("Float64() = %v, want %v", val1, expectedVal1)
	}

	if val2 != expectedVal2 {
		t.Errorf("Float64() second call = %v, want %v", val2, expectedVal2)
	}
}

// TestInt64N tests the Int64N function
func TestInt64N(t *testing.T) {
	setupTest(t)
	defer teardownTest(t)

	// Test multiple calls to ensure reproducibility
	val1 := Int64N(100)
	val2 := Int64N(100)

	// With a fixed seed, we should get consistent results
	expectedVal1 := int64(74)
	expectedVal2 := int64(24)

	if val1 != expectedVal1 {
		t.Errorf("Int64N(100) = %v, want %v", val1, expectedVal1)
	}

	if val2 != expectedVal2 {
		t.Errorf("Int64N(100) second call = %v, want %v", val2, expectedVal2)
	}

	// Test with different bounds
	val3 := Int64N(1000)
	expectedVal3 := int64(741)

	if val3 != expectedVal3 {
		t.Errorf("Int64N(1000) = %v, want %v", val3, expectedVal3)
	}
}

// BenchmarkIntN benchmarks the IntN function
func BenchmarkIntN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntN(100)
	}
}

// BenchmarkFloat64 benchmarks the Float64 function
func BenchmarkFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Float64()
	}
}

// BenchmarkInt64N benchmarks the Int64N function
func BenchmarkInt64N(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int64N(100)
	}
}
