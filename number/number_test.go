package number

import (
	"math/rand/v2"
	"testing"

	"github.com/khchehab/muzayaf/random"
)

// setupTest sets up a deterministic random source for testing
func setupTest(t *testing.T) {
	t.Helper()
	// Use a fixed seed for reproducible tests
	src := rand.NewPCG(42, 43)
	random.SetRandomSource(src)
}

// teardownTest resets the random source after testing
func teardownTest(t *testing.T) {
	t.Helper()
	random.ResetRandomSource()
}

// TestInt tests the Int function
func TestInt(t *testing.T) {
	setupTest(t)
	defer teardownTest(t)

	// Test multiple calls to ensure reproducibility
	int1 := Int()
	int2 := Int()

	// With a fixed seed, we should get consistent results
	expectedInt1 := 6910346261119195341
	expectedInt2 := 2231614027494634719

	if int1 != expectedInt1 {
		t.Errorf("Int() = %v, want %v", int1, expectedInt1)
	}

	if int2 != expectedInt2 {
		t.Errorf("Int() second call = %v, want %v", int2, expectedInt2)
	}

	// Test with min and max options
	boundedInt := Int(WithIntMin(10), WithIntMax(20))
	expectedBoundedInt := 18

	if boundedInt != expectedBoundedInt {
		t.Errorf("Int(WithIntMin(10), WithIntMax(20)) = %v, want %v", boundedInt, expectedBoundedInt)
	}

	// Test with multiple option
	multipleInt := Int(WithIntMin(0), WithIntMax(100), WithIntMultiple(5))
	expectedMultipleInt := 90

	if multipleInt != expectedMultipleInt {
		t.Errorf("Int(WithIntMin(0), WithIntMax(100), WithIntMultiple(5)) = %v, want %v", multipleInt, expectedMultipleInt)
	}
}

// TestFloat tests the Float function
func TestFloat(t *testing.T) {
	setupTest(t)
	defer teardownTest(t)

	// Test multiple calls to ensure reproducibility
	float1 := Float()
	float2 := Float()

	// With a fixed seed, we should get consistent results
	expectedFloat1 := 0.41
	expectedFloat2 := 0.52

	if float1 != expectedFloat1 {
		t.Errorf("Float() = %v, want %v", float1, expectedFloat1)
	}

	if float2 != expectedFloat2 {
		t.Errorf("Float() second call = %v, want %v", float2, expectedFloat2)
	}

	// Test with min and max options
	boundedFloat := Float(WithFloatMin(10.0), WithFloatMax(20.0))
	expectedBoundedFloat := 16.4

	if boundedFloat != expectedBoundedFloat {
		t.Errorf("Float(WithFloatMin(10.0), WithFloatMax(20.0)) = %v, want %v", boundedFloat, expectedBoundedFloat)
	}

	// Test with fraction digits option
	preciseFloat := Float(WithFloatFractionDigits(4))
	expectedPreciseFloat := 0.8828

	if preciseFloat != expectedPreciseFloat {
		t.Errorf("Float(WithFloatFractionDigits(4)) = %v, want %v", preciseFloat, expectedPreciseFloat)
	}
}

// TestBinary tests the Binary function
func TestBinary(t *testing.T) {
	setupTest(t)
	defer teardownTest(t)

	// Test multiple calls to ensure reproducibility
	binary1 := Binary()
	binary2 := Binary()

	// With a fixed seed, we should get consistent results
	expectedBinary1 := "1"
	expectedBinary2 := "1"

	if binary1 != expectedBinary1 {
		t.Errorf("Binary() = %v, want %v", binary1, expectedBinary1)
	}

	if binary2 != expectedBinary2 {
		t.Errorf("Binary() second call = %v, want %v", binary2, expectedBinary2)
	}

	// Test with min and max options
	boundedBinary := Binary(WithBinaryMin(10), WithBinaryMax(20))
	expectedBoundedBinary := "10010"

	if boundedBinary != expectedBoundedBinary {
		t.Errorf("Binary(WithBinaryMin(10), WithBinaryMax(20)) = %v, want %v", boundedBinary, expectedBoundedBinary)
	}

	// Test with prefix option
	prefixedBinary := Binary(WithBinaryPrefix(true))
	expectedPrefixedBinary := "0b0"

	if prefixedBinary != expectedPrefixedBinary {
		t.Errorf("Binary(WithBinaryPrefix(true)) = %v, want %v", prefixedBinary, expectedPrefixedBinary)
	}
}

// TestOctal tests the Octal function
func TestOctal(t *testing.T) {
	setupTest(t)
	defer teardownTest(t)

	// Test multiple calls to ensure reproducibility
	octal1 := Octal()
	octal2 := Octal()

	// With a fixed seed, we should get consistent results
	expectedOctal1 := "5"
	expectedOctal2 := "7"

	if octal1 != expectedOctal1 {
		t.Errorf("Octal() = %v, want %v", octal1, expectedOctal1)
	}

	if octal2 != expectedOctal2 {
		t.Errorf("Octal() second call = %v, want %v", octal2, expectedOctal2)
	}

	// Test with min and max options
	boundedOctal := Octal(WithOctalMin(10), WithOctalMax(20))
	expectedBoundedOctal := "22"

	if boundedOctal != expectedBoundedOctal {
		t.Errorf("Octal(WithOctalMin(10), WithOctalMax(20)) = %v, want %v", boundedOctal, expectedBoundedOctal)
	}

	// Test with prefix option
	prefixedOctal := Octal(WithOctalPrefix(true))
	expectedPrefixedOctal := "06"

	if prefixedOctal != expectedPrefixedOctal {
		t.Errorf("Octal(WithOctalPrefix(true)) = %v, want %v", prefixedOctal, expectedPrefixedOctal)
	}
}

// TestHex tests the Hex function
func TestHex(t *testing.T) {
	setupTest(t)
	defer teardownTest(t)

	// Test multiple calls to ensure reproducibility
	hex1 := Hex()
	hex2 := Hex()

	// With a fixed seed, we should get consistent results
	expectedHex1 := "d"
	expectedHex2 := "f"

	if hex1 != expectedHex1 {
		t.Errorf("Hex() = %v, want %v", hex1, expectedHex1)
	}

	if hex2 != expectedHex2 {
		t.Errorf("Hex() second call = %v, want %v", hex2, expectedHex2)
	}

	// Test with min and max options
	boundedHex := Hex(WithHexMin(10), WithHexMax(20))
	expectedBoundedHex := "12"

	if boundedHex != expectedBoundedHex {
		t.Errorf("Hex(WithHexMin(10), WithHexMax(20)) = %v, want %v", boundedHex, expectedBoundedHex)
	}

	// Test with prefix option
	prefixedHex := Hex(WithHexPrefix(true))
	expectedPrefixedHex := "0x6"

	if prefixedHex != expectedPrefixedHex {
		t.Errorf("Hex(WithHexPrefix(true)) = %v, want %v", prefixedHex, expectedPrefixedHex)
	}
}

// TestRoman tests the Roman function
func TestRoman(t *testing.T) {
	setupTest(t)
	defer teardownTest(t)

	// Test multiple calls to ensure reproducibility
	roman1 := Roman()
	roman2 := Roman()

	// With a fixed seed, we should get consistent results
	expectedRoman1 := "MMCMXCVII"
	expectedRoman2 := "CMLXVIII"

	if roman1 != expectedRoman1 {
		t.Errorf("Roman() = %v, want %v", roman1, expectedRoman1)
	}

	if roman2 != expectedRoman2 {
		t.Errorf("Roman() second call = %v, want %v", roman2, expectedRoman2)
	}

	// Test with min and max options
	boundedRoman := Roman(WithRomanMin(10), WithRomanMax(20))
	expectedBoundedRoman := "XVIII"

	if boundedRoman != expectedBoundedRoman {
		t.Errorf("Roman(WithRomanMin(10), WithRomanMax(20)) = %v, want %v", boundedRoman, expectedBoundedRoman)
	}
}

// BenchmarkInt benchmarks the Int function
func BenchmarkInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int()
	}
}

// BenchmarkFloat benchmarks the Float function
func BenchmarkFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Float()
	}
}

// BenchmarkBinary benchmarks the Binary function
func BenchmarkBinary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Binary()
	}
}

// BenchmarkOctal benchmarks the Octal function
func BenchmarkOctal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Octal()
	}
}

// BenchmarkHex benchmarks the Hex function
func BenchmarkHex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hex()
	}
}

// BenchmarkRoman benchmarks the Roman function
func BenchmarkRoman(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Roman()
	}
}
