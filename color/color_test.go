package color

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

// TestRGBA tests the RGBA function
func TestRGBA(t *testing.T) {
	setupTest(t)
	defer teardownTest(t)

	// Test multiple calls to ensure reproducibility
	rgba1 := RGBA()
	rgba2 := RGBA()

	// With a fixed seed, we should get consistent results
	expectedRGBA1 := RGBAColor{Red: 157, Green: 191, Blue: 251, Alpha: 0.87}
	expectedRGBA2 := RGBAColor{Red: 164, Green: 129, Blue: 47, Alpha: 0.97}

	// Test struct values
	if rgba1.Red != expectedRGBA1.Red || rgba1.Green != expectedRGBA1.Green ||
		rgba1.Blue != expectedRGBA1.Blue || rgba1.Alpha != expectedRGBA1.Alpha {
		t.Errorf("RGBA() = %+v, want %+v", rgba1, expectedRGBA1)
	}

	if rgba2.Red != expectedRGBA2.Red || rgba2.Green != expectedRGBA2.Green ||
		rgba2.Blue != expectedRGBA2.Blue || rgba2.Alpha != expectedRGBA2.Alpha {
		t.Errorf("RGBA() second call = %+v, want %+v", rgba2, expectedRGBA2)
	}

	// Test String method for first call (Alpha < 1.0)
	expectedString1 := "rgba(157, 191, 251, 0.87)"
	if rgba1.String() != expectedString1 {
		t.Errorf("RGBA().String() = %v, want %v", rgba1.String(), expectedString1)
	}

	// Test String method for second call (Alpha = 0.97)
	expectedString2 := "rgba(164, 129, 47, 0.97)"
	t.Logf("RGBA second call: Alpha = %v, String() = %v", rgba2.Alpha, rgba2.String())
	if rgba2.String() != expectedString2 {
		t.Errorf("RGBA() second call .String() = %v, want %v", rgba2.String(), expectedString2)
	}
}

// TestRGB tests the RGB function
func TestRGB(t *testing.T) {
	setupTest(t)
	defer teardownTest(t)

	// Test multiple calls to ensure reproducibility
	rgb1 := RGB()
	rgb2 := RGB()

	// With a fixed seed, we should get consistent results
	expectedRGB1 := RGBAColor{Red: 157, Green: 191, Blue: 251, Alpha: 1.0}
	expectedRGB2 := RGBAColor{Red: 102, Green: 164, Blue: 129, Alpha: 1.0}

	// Test struct values
	if rgb1.Red != expectedRGB1.Red || rgb1.Green != expectedRGB1.Green ||
		rgb1.Blue != expectedRGB1.Blue || rgb1.Alpha != expectedRGB1.Alpha {
		t.Errorf("RGB() = %+v, want %+v", rgb1, expectedRGB1)
	}

	if rgb2.Red != expectedRGB2.Red || rgb2.Green != expectedRGB2.Green ||
		rgb2.Blue != expectedRGB2.Blue || rgb2.Alpha != expectedRGB2.Alpha {
		t.Errorf("RGB() second call = %+v, want %+v", rgb2, expectedRGB2)
	}

	// Test String method
	expectedString := "rgb(157, 191, 251)"
	if rgb1.String() != expectedString {
		t.Errorf("RGB().String() = %v, want %v", rgb1.String(), expectedString)
	}
}

// TestCMYK tests the CMYK function
func TestCMYK(t *testing.T) {
	setupTest(t)
	defer teardownTest(t)

	// Test multiple calls to ensure reproducibility
	cmyk1 := CMYK()
	cmyk2 := CMYK()

	// With a fixed seed, we should get consistent results
	expectedCMYK1 := CMYKColor{Cyan: 75, Magenta: 24, Yellow: 74, Key: 87}
	expectedCMYK2 := CMYKColor{Cyan: 76, Magenta: 99, Yellow: 44, Key: 97}

	// Test struct values
	if cmyk1.Cyan != expectedCMYK1.Cyan || cmyk1.Magenta != expectedCMYK1.Magenta ||
		cmyk1.Yellow != expectedCMYK1.Yellow || cmyk1.Key != expectedCMYK1.Key {
		t.Errorf("CMYK() = %+v, want %+v", cmyk1, expectedCMYK1)
	}

	if cmyk2.Cyan != expectedCMYK2.Cyan || cmyk2.Magenta != expectedCMYK2.Magenta ||
		cmyk2.Yellow != expectedCMYK2.Yellow || cmyk2.Key != expectedCMYK2.Key {
		t.Errorf("CMYK() second call = %+v, want %+v", cmyk2, expectedCMYK2)
	}

	// Test String method
	expectedString := "cmyk(75%, 24%, 74%, 87%)"
	if cmyk1.String() != expectedString {
		t.Errorf("CMYK().String() = %v, want %v", cmyk1.String(), expectedString)
	}
}

// TestHSLA tests the HSLA function
func TestHSLA(t *testing.T) {
	setupTest(t)
	defer teardownTest(t)

	// Test multiple calls to ensure reproducibility
	hsla1 := HSLA()
	hsla2 := HSLA()

	// With a fixed seed, we should get consistent results
	expectedHSLA1 := HSLAColor{Hue: 270, Saturation: 24, Lightness: 74, Alpha: 0.87}
	expectedHSLA2 := HSLAColor{Hue: 272, Saturation: 99, Lightness: 44, Alpha: 0.97}

	// Test struct values
	if hsla1.Hue != expectedHSLA1.Hue || hsla1.Saturation != expectedHSLA1.Saturation ||
		hsla1.Lightness != expectedHSLA1.Lightness || hsla1.Alpha != expectedHSLA1.Alpha {
		t.Errorf("HSLA() = %+v, want %+v", hsla1, expectedHSLA1)
	}

	if hsla2.Hue != expectedHSLA2.Hue || hsla2.Saturation != expectedHSLA2.Saturation ||
		hsla2.Lightness != expectedHSLA2.Lightness || hsla2.Alpha != expectedHSLA2.Alpha {
		t.Errorf("HSLA() second call = %+v, want %+v", hsla2, expectedHSLA2)
	}

	// Test String method for first call (Alpha < 1.0)
	expectedString1 := "hsla(270°, 24%, 74%, 0.87)"
	if hsla1.String() != expectedString1 {
		t.Errorf("HSLA().String() = %v, want %v", hsla1.String(), expectedString1)
	}

	// Test String method for second call (Alpha = 0.97)
	expectedString2 := "hsla(272°, 99%, 44%, 0.97)"
	t.Logf("HSLA second call: Alpha = %v, String() = %v", hsla2.Alpha, hsla2.String())
	if hsla2.String() != expectedString2 {
		t.Errorf("HSLA() second call .String() = %v, want %v", hsla2.String(), expectedString2)
	}
}

// TestHSL tests the HSL function
func TestHSL(t *testing.T) {
	setupTest(t)
	defer teardownTest(t)

	// Test multiple calls to ensure reproducibility
	hsl1 := HSL()
	hsl2 := HSL()

	// With a fixed seed, we should get consistent results
	expectedHSL1 := HSLAColor{Hue: 270, Saturation: 24, Lightness: 74, Alpha: 1.0}
	expectedHSL2 := HSLAColor{Hue: 313, Saturation: 76, Lightness: 99, Alpha: 1.0}

	// Test struct values
	if hsl1.Hue != expectedHSL1.Hue || hsl1.Saturation != expectedHSL1.Saturation ||
		hsl1.Lightness != expectedHSL1.Lightness || hsl1.Alpha != expectedHSL1.Alpha {
		t.Errorf("HSL() = %+v, want %+v", hsl1, expectedHSL1)
	}

	if hsl2.Hue != expectedHSL2.Hue || hsl2.Saturation != expectedHSL2.Saturation ||
		hsl2.Lightness != expectedHSL2.Lightness || hsl2.Alpha != expectedHSL2.Alpha {
		t.Errorf("HSL() second call = %+v, want %+v", hsl2, expectedHSL2)
	}

	// Test String method
	expectedString := "hsl(270°, 24%, 74%)"
	if hsl1.String() != expectedString {
		t.Errorf("HSL().String() = %v, want %v", hsl1.String(), expectedString)
	}
}

// TestColorName tests the ColorName function
func TestColorName(t *testing.T) {
	setupTest(t)
	defer teardownTest(t)

	// Test with default locale
	name1 := ColorName()
	name2 := ColorName()

	// With a fixed seed, we should get consistent results
	expectedName1 := "paleturquoise"
	expectedName2 := "darkseagreen"

	if name1 != expectedName1 {
		t.Errorf("ColorName() = %v, want %v", name1, expectedName1)
	}

	if name2 != expectedName2 {
		t.Errorf("ColorName() second call = %v, want %v", name2, expectedName2)
	}
}

// TestColorNameWithLocale tests the ColorName function with a specific locale
func TestColorNameWithLocale(t *testing.T) {
	setupTest(t)
	defer teardownTest(t)

	// Test with explicit "en" locale
	name := ColorName(WithLocale("en"))
	expectedName := "paleturquoise"

	if name != expectedName {
		t.Errorf("ColorName(WithLocale(\"en\")) = %v, want %v", name, expectedName)
	}

	// Test with non-existent locale (should fall back to "en")
	name = ColorName(WithLocale("non-existent"))
	expectedName = "darkseagreen" // Same as "en" with our fixed seed

	if name != expectedName {
		t.Errorf("ColorName(WithLocale(\"non-existent\")) = %v, want %v", name, expectedName)
	}
}

// TestRGBAColorString tests the String method of RGBAColor
func TestRGBAColorString(t *testing.T) {
	// Test with alpha < 1.0
	rgba := RGBAColor{Red: 255, Green: 0, Blue: 0, Alpha: 0.5}
	expected := "rgba(255, 0, 0, 0.50)"
	if rgba.String() != expected {
		t.Errorf("RGBAColor{255, 0, 0, 0.5}.String() = %v, want %v", rgba.String(), expected)
	}

	// Test with alpha = 1.0
	rgba = RGBAColor{Red: 255, Green: 0, Blue: 0, Alpha: 1.0}
	expected = "rgb(255, 0, 0)"
	if rgba.String() != expected {
		t.Errorf("RGBAColor{255, 0, 0, 1.0}.String() = %v, want %v", rgba.String(), expected)
	}
}

// TestHSLAColorString tests the String method of HSLAColor
func TestHSLAColorString(t *testing.T) {
	// Test with alpha < 1.0
	hsla := HSLAColor{Hue: 0, Saturation: 100, Lightness: 50, Alpha: 0.5}
	expected := "hsla(0°, 100%, 50%, 0.50)"
	if hsla.String() != expected {
		t.Errorf("HSLAColor{0, 100, 50, 0.5}.String() = %v, want %v", hsla.String(), expected)
	}

	// Test with alpha = 1.0
	hsla = HSLAColor{Hue: 0, Saturation: 100, Lightness: 50, Alpha: 1.0}
	expected = "hsl(0°, 100%, 50%)"
	if hsla.String() != expected {
		t.Errorf("HSLAColor{0, 100, 50, 1.0}.String() = %v, want %v", hsla.String(), expected)
	}
}

// TestCMYKColorString tests the String method of CMYKColor
func TestCMYKColorString(t *testing.T) {
	cmyk := CMYKColor{Cyan: 0, Magenta: 100, Yellow: 100, Key: 0}
	expected := "cmyk(0%, 100%, 100%, 0%)"
	if cmyk.String() != expected {
		t.Errorf("CMYKColor{0, 100, 100, 0}.String() = %v, want %v", cmyk.String(), expected)
	}
}

// BenchmarkRGBA benchmarks the RGBA function
func BenchmarkRGBA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RGBA()
	}
}

// BenchmarkRGB benchmarks the RGB function
func BenchmarkRGB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RGB()
	}
}

// BenchmarkCMYK benchmarks the CMYK function
func BenchmarkCMYK(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CMYK()
	}
}

// BenchmarkHSLA benchmarks the HSLA function
func BenchmarkHSLA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HSLA()
	}
}

// BenchmarkHSL benchmarks the HSL function
func BenchmarkHSL(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HSL()
	}
}

// BenchmarkColorName benchmarks the ColorName function
func BenchmarkColorName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ColorName()
	}
}

// BenchmarkColorNameWithLocale benchmarks the ColorName function with locale option
func BenchmarkColorNameWithLocale(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ColorName(WithLocale("en"))
	}
}

// BenchmarkRGBAColorString benchmarks the String method of RGBAColor
func BenchmarkRGBAColorString(b *testing.B) {
	rgba := RGBA()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = rgba.String()
	}
}

// BenchmarkHSLAColorString benchmarks the String method of HSLAColor
func BenchmarkHSLAColorString(b *testing.B) {
	hsla := HSLA()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = hsla.String()
	}
}

// BenchmarkCMYKColorString benchmarks the String method of CMYKColor
func BenchmarkCMYKColorString(b *testing.B) {
	cmyk := CMYK()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = cmyk.String()
	}
}
