package date

import (
	"math/rand/v2"
	"testing"
	"time"

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

// TestAny tests the Any function
func TestAny(t *testing.T) {
	setupTest(t)
	defer teardownTest(t)

	// Fix the relative date for testing
	relativeDate := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)

	// Test multiple calls to ensure reproducibility
	date1 := Any(WithRelative(relativeDate))
	date2 := Any(WithRelative(relativeDate))

	// With a fixed seed and relative date, we should get consistent results
	expectedDate1 := time.Date(2049, 11, 4, 20, 42, 6, 0, time.UTC)
	expectedDate2 := time.Date(1948, 5, 23, 8, 34, 32, 0, time.UTC)

	// Test date values
	if !date1.Equal(expectedDate1) {
		t.Errorf("Any(WithRelative) = %v, want %v", date1, expectedDate1)
	}

	if !date2.Equal(expectedDate2) {
		t.Errorf("Any(WithRelative) second call = %v, want %v", date2, expectedDate2)
	}

	// Test with custom options
	date3 := Any(WithRelative(relativeDate), WithYears(5))
	expectedDate3 := time.Date(2002, 5, 31, 0, 0, 49, 0, time.UTC)

	if !date3.Equal(expectedDate3) {
		t.Errorf("Any(WithRelative, WithYears) = %v, want %v", date3, expectedDate3)
	}
}

// TestBetween tests the Between function
func TestBetween(t *testing.T) {
	setupTest(t)
	defer teardownTest(t)

	from := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	to := time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC)

	// Test multiple calls to ensure reproducibility
	date1 := Between(from, to)
	date2 := Between(from, to)

	// With a fixed seed, we should get consistent results
	expectedDate1 := time.Date(2000, 10, 1, 5, 9, 34, 0, time.UTC)
	expectedDate2 := time.Date(2000, 3, 29, 13, 18, 25, 0, time.UTC)

	// Test date values
	if !date1.Equal(expectedDate1) {
		t.Errorf("Between() = %v, want %v", date1, expectedDate1)
	}

	if !date2.Equal(expectedDate2) {
		t.Errorf("Between() second call = %v, want %v", date2, expectedDate2)
	}

	// Test with reversed dates (should swap them)
	date3 := Between(to, from)
	expectedDate3 := time.Date(2000, 9, 28, 5, 15, 3, 0, time.UTC)

	if !date3.Equal(expectedDate3) {
		t.Errorf("Between(to, from) = %v, want %v", date3, expectedDate3)
	}

	// Test with same date
	sameDate := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	date4 := Between(sameDate, sameDate)

	if !date4.Equal(sameDate) {
		t.Errorf("Between(sameDate, sameDate) = %v, want %v", date4, sameDate)
	}
}

// TestFuture tests the Future function
func TestFuture(t *testing.T) {
	setupTest(t)
	defer teardownTest(t)

	// Fix the relative date for testing
	relativeDate := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)

	// Test multiple calls to ensure reproducibility
	date1 := Future(WithRelative(relativeDate))
	date2 := Future(WithRelative(relativeDate))

	// With a fixed seed, we should get consistent results
	expectedDate1 := time.Date(2074, 12, 3, 13, 21, 37, 0, time.UTC)
	expectedDate2 := time.Date(2024, 3, 13, 1, 23, 4, 0, time.UTC)

	// Test date values
	if !date1.Equal(expectedDate1) {
		t.Errorf("Future() = %v, want %v", date1, expectedDate1)
	}

	if !date2.Equal(expectedDate2) {
		t.Errorf("Future() second call = %v, want %v", date2, expectedDate2)
	}

	// Test with custom years
	date3 := Future(WithRelative(relativeDate), WithYears(5))
	expectedDate3 := time.Date(2003, 9, 16, 3, 6, 52, 0, time.UTC)

	if !date3.Equal(expectedDate3) {
		t.Errorf("Future(WithRelative, WithYears) = %v, want %v", date3, expectedDate3)
	}
}

// TestPast tests the Past function
func TestPast(t *testing.T) {
	setupTest(t)
	defer teardownTest(t)

	// Fix the relative date for testing
	relativeDate := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)

	// Test multiple calls to ensure reproducibility
	date1 := Past(WithRelative(relativeDate))
	date2 := Past(WithRelative(relativeDate))

	// With a fixed seed, we should get consistent results
	expectedDate1 := time.Date(1974, 12, 2, 19, 22, 44, 0, time.UTC)
	expectedDate2 := time.Date(1924, 3, 12, 19, 34, 39, 0, time.UTC)

	// Test date values
	if !date1.Equal(expectedDate1) {
		t.Errorf("Past() = %v, want %v", date1, expectedDate1)
	}

	if !date2.Equal(expectedDate2) {
		t.Errorf("Past() second call = %v, want %v", date2, expectedDate2)
	}

	// Test with custom years
	date3 := Past(WithRelative(relativeDate), WithYears(5))
	expectedDate3 := time.Date(1998, 9, 14, 9, 19, 46, 0, time.UTC)

	if !date3.Equal(expectedDate3) {
		t.Errorf("Past(WithRelative, WithYears) = %v, want %v", date3, expectedDate3)
	}
}

// TestMonth tests the Month function
func TestMonth(t *testing.T) {
	setupTest(t)
	defer teardownTest(t)

	// Test multiple calls to ensure reproducibility
	month1 := Month()
	month2 := Month()

	// With a fixed seed, we should get consistent results
	expectedMonth1 := "September"
	expectedMonth2 := "March"

	if month1 != expectedMonth1 {
		t.Errorf("Month() = %v, want %v", month1, expectedMonth1)
	}

	if month2 != expectedMonth2 {
		t.Errorf("Month() second call = %v, want %v", month2, expectedMonth2)
	}

	// Test with locale option
	month3 := Month(WithLocale("en"))
	expectedMonth3 := "September"

	if month3 != expectedMonth3 {
		t.Errorf("Month(WithLocale(\"en\")) = %v, want %v", month3, expectedMonth3)
	}

	// Test with non-existent locale (should fall back to "en")
	month4 := Month(WithLocale("non-existent"))
	expectedMonth4 := "November"

	if month4 != expectedMonth4 {
		t.Errorf("Month(WithLocale(\"non-existent\")) = %v, want %v", month4, expectedMonth4)
	}
}

// TestWeekday tests the Weekday function
func TestWeekday(t *testing.T) {
	setupTest(t)
	defer teardownTest(t)

	// Test multiple calls to ensure reproducibility
	weekday1 := Weekday()
	weekday2 := Weekday()

	// With a fixed seed, we should get consistent results
	expectedWeekday1 := "Saturday"
	expectedWeekday2 := "Tuesday"

	if weekday1 != expectedWeekday1 {
		t.Errorf("Weekday() = %v, want %v", weekday1, expectedWeekday1)
	}

	if weekday2 != expectedWeekday2 {
		t.Errorf("Weekday() second call = %v, want %v", weekday2, expectedWeekday2)
	}

	// Test with locale option
	weekday3 := Weekday(WithLocale("en"))
	expectedWeekday3 := "Saturday"

	if weekday3 != expectedWeekday3 {
		t.Errorf("Weekday(WithLocale(\"en\")) = %v, want %v", weekday3, expectedWeekday3)
	}

	// Test with non-existent locale (should fall back to "en")
	weekday4 := Weekday(WithLocale("non-existent"))
	expectedWeekday4 := "Sunday"

	if weekday4 != expectedWeekday4 {
		t.Errorf("Weekday(WithLocale(\"non-existent\")) = %v, want %v", weekday4, expectedWeekday4)
	}
}

// TestTimezone tests the Timezone function
func TestTimezone(t *testing.T) {
	setupTest(t)
	defer teardownTest(t)

	// Test multiple calls to ensure reproducibility
	timezone1 := Timezone()
	timezone2 := Timezone()

	// With a fixed seed, we should get consistent results
	expectedTimezone1 := "Europe/Athens"
	expectedTimezone2 := "America/Fortaleza"

	if timezone1 != expectedTimezone1 {
		t.Errorf("Timezone() = %v, want %v", timezone1, expectedTimezone1)
	}

	if timezone2 != expectedTimezone2 {
		t.Errorf("Timezone() second call = %v, want %v", timezone2, expectedTimezone2)
	}

	// Test with locale option
	timezone3 := Timezone(WithLocale("en"))
	expectedTimezone3 := "Australia/Sydney"

	if timezone3 != expectedTimezone3 {
		t.Errorf("Timezone(WithLocale(\"en\")) = %v, want %v", timezone3, expectedTimezone3)
	}

	// Test with non-existent locale (should fall back to base)
	timezone4 := Timezone(WithLocale("non-existent"))
	expectedTimezone4 := "Europe/Vilnius"

	if timezone4 != expectedTimezone4 {
		t.Errorf("Timezone(WithLocale(\"non-existent\")) = %v, want %v", timezone4, expectedTimezone4)
	}
}

// BenchmarkAny benchmarks the Any function
func BenchmarkAny(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Any()
	}
}

// BenchmarkBetween benchmarks the Between function
func BenchmarkBetween(b *testing.B) {
	from := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	to := time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC)

	for i := 0; i < b.N; i++ {
		Between(from, to)
	}
}

// BenchmarkFuture benchmarks the Future function
func BenchmarkFuture(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Future()
	}
}

// BenchmarkPast benchmarks the Past function
func BenchmarkPast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Past()
	}
}

// BenchmarkMonth benchmarks the Month function
func BenchmarkMonth(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Month()
	}
}

// BenchmarkWeekday benchmarks the Weekday function
func BenchmarkWeekday(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Weekday()
	}
}

// BenchmarkTimezone benchmarks the Timezone function
func BenchmarkTimezone(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Timezone()
	}
}
