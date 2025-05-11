package person

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

// TestFirstName tests the FirstName function
func TestFirstName(t *testing.T) {
	setupTest(t)
	defer teardownTest(t)

	// Test multiple calls to ensure reproducibility
	firstName1 := FirstName()
	firstName2 := FirstName()

	// With a fixed seed, we should get consistent results
	expectedFirstName1 := "George"
	expectedFirstName2 := "Georgia"

	if firstName1 != expectedFirstName1 {
		t.Errorf("FirstName() = %v, want %v", firstName1, expectedFirstName1)
	}

	if firstName2 != expectedFirstName2 {
		t.Errorf("FirstName() second call = %v, want %v", firstName2, expectedFirstName2)
	}

	// Test with gender option - female
	femaleFirstName := FirstName(WithGender(GenderFemale))
	expectedFemaleFirstName := "Lucy"

	if femaleFirstName != expectedFemaleFirstName {
		t.Errorf("FirstName(WithGender(GenderFemale)) = %v, want %v", femaleFirstName, expectedFemaleFirstName)
	}

	// Test with gender option - male
	maleFirstName := FirstName(WithGender(GenderMale))
	expectedMaleFirstName := "Nathan"

	if maleFirstName != expectedMaleFirstName {
		t.Errorf("FirstName(WithGender(GenderMale)) = %v, want %v", maleFirstName, expectedMaleFirstName)
	}

	// Test with locale option
	localeFirstName := FirstName(WithLocale("en"))
	expectedLocaleFirstName := "Gerald"

	if localeFirstName != expectedLocaleFirstName {
		t.Errorf("FirstName(WithLocale(\"en\")) = %v, want %v", localeFirstName, expectedLocaleFirstName)
	}
}

// TestMiddleName tests the MiddleName function
func TestMiddleName(t *testing.T) {
	setupTest(t)
	defer teardownTest(t)

	// Test multiple calls to ensure reproducibility
	middleName1 := MiddleName()
	middleName2 := MiddleName()

	// With a fixed seed, we should get consistent results
	expectedMiddleName1 := "Graham"
	expectedMiddleName2 := "Frances"

	if middleName1 != expectedMiddleName1 {
		t.Errorf("MiddleName() = %v, want %v", middleName1, expectedMiddleName1)
	}

	if middleName2 != expectedMiddleName2 {
		t.Errorf("MiddleName() second call = %v, want %v", middleName2, expectedMiddleName2)
	}

	// Test with gender option - female
	femaleMiddleName := MiddleName(WithGender(GenderFemale))
	expectedFemaleMiddleName := "Laura"

	if femaleMiddleName != expectedFemaleMiddleName {
		t.Errorf("MiddleName(WithGender(GenderFemale)) = %v, want %v", femaleMiddleName, expectedFemaleMiddleName)
	}

	// Test with gender option - male
	maleMiddleName := MiddleName(WithGender(GenderMale))
	expectedMaleMiddleName := "Nicholas"

	if maleMiddleName != expectedMaleMiddleName {
		t.Errorf("MiddleName(WithGender(GenderMale)) = %v, want %v", maleMiddleName, expectedMaleMiddleName)
	}

	// Test with locale option
	localeMiddleName := MiddleName(WithLocale("en"))
	expectedLocaleMiddleName := "Grant"

	if localeMiddleName != expectedLocaleMiddleName {
		t.Errorf("MiddleName(WithLocale(\"en\")) = %v, want %v", localeMiddleName, expectedLocaleMiddleName)
	}
}

// TestLastName tests the LastName function
func TestLastName(t *testing.T) {
	setupTest(t)
	defer teardownTest(t)

	// Test multiple calls to ensure reproducibility
	lastName1 := LastName()
	lastName2 := LastName()

	// With a fixed seed, we should get consistent results
	expectedLastName1 := "Parker"
	expectedLastName2 := "Edwards"

	if lastName1 != expectedLastName1 {
		t.Errorf("LastName() = %v, want %v", lastName1, expectedLastName1)
	}

	if lastName2 != expectedLastName2 {
		t.Errorf("LastName() second call = %v, want %v", lastName2, expectedLastName2)
	}

	// Test with locale option
	localeLastName := LastName(WithLocale("en"))
	expectedLocaleLastName := "Parker"

	if localeLastName != expectedLocaleLastName {
		t.Errorf("LastName(WithLocale(\"en\")) = %v, want %v", localeLastName, expectedLocaleLastName)
	}
}

// TestGender tests the Gender function
func TestGender(t *testing.T) {
	setupTest(t)
	defer teardownTest(t)

	// Test multiple calls to ensure reproducibility
	gender1 := Gender()
	gender2 := Gender()

	// With a fixed seed, we should get consistent results
	expectedGender1 := "male"
	expectedGender2 := "male"

	if gender1 != expectedGender1 {
		t.Errorf("Gender() = %v, want %v", gender1, expectedGender1)
	}

	if gender2 != expectedGender2 {
		t.Errorf("Gender() second call = %v, want %v", gender2, expectedGender2)
	}

	// Test with locale option
	localeGender := Gender(WithLocale("en"))
	expectedLocaleGender := "male"

	if localeGender != expectedLocaleGender {
		t.Errorf("Gender(WithLocale(\"en\")) = %v, want %v", localeGender, expectedLocaleGender)
	}
}

// TestPrefix tests the Prefix function
func TestPrefix(t *testing.T) {
	setupTest(t)
	defer teardownTest(t)

	// Test multiple calls to ensure reproducibility
	prefix1 := Prefix()
	prefix2 := Prefix()

	// With a fixed seed, we should get consistent results
	expectedPrefix1 := "Sir"
	expectedPrefix2 := "Ms."

	if prefix1 != expectedPrefix1 {
		t.Errorf("Prefix() = %v, want %v", prefix1, expectedPrefix1)
	}

	if prefix2 != expectedPrefix2 {
		t.Errorf("Prefix() second call = %v, want %v", prefix2, expectedPrefix2)
	}

	// Test with gender option - female
	femalePrefix := Prefix(WithGender(GenderFemale))
	expectedFemalePrefix := "Dr."

	if femalePrefix != expectedFemalePrefix {
		t.Errorf("Prefix(WithGender(GenderFemale)) = %v, want %v", femalePrefix, expectedFemalePrefix)
	}

	// Test with gender option - male
	malePrefix := Prefix(WithGender(GenderMale))
	expectedMalePrefix := "Professor"

	if malePrefix != expectedMalePrefix {
		t.Errorf("Prefix(WithGender(GenderMale)) = %v, want %v", malePrefix, expectedMalePrefix)
	}

	// Test with locale option
	localePrefix := Prefix(WithLocale("en"))
	expectedLocalePrefix := "Lord"

	if localePrefix != expectedLocalePrefix {
		t.Errorf("Prefix(WithLocale(\"en\")) = %v, want %v", localePrefix, expectedLocalePrefix)
	}
}

// TestSuffix tests the Suffix function
func TestSuffix(t *testing.T) {
	setupTest(t)
	defer teardownTest(t)

	// Test multiple calls to ensure reproducibility
	suffix1 := Suffix()
	suffix2 := Suffix()

	// With a fixed seed, we should get consistent results
	expectedSuffix1 := "Sr."
	expectedSuffix2 := "MA"

	if suffix1 != expectedSuffix1 {
		t.Errorf("Suffix() = %v, want %v", suffix1, expectedSuffix1)
	}

	if suffix2 != expectedSuffix2 {
		t.Errorf("Suffix() second call = %v, want %v", suffix2, expectedSuffix2)
	}

	// Test with suffix type option - academic
	academicSuffix := Suffix(WithSuffixType(SuffixTypeAcademic))
	expectedAcademicSuffix := "DO"

	if academicSuffix != expectedAcademicSuffix {
		t.Errorf("Suffix(WithSuffixType(SuffixTypeAcademic)) = %v, want %v", academicSuffix, expectedAcademicSuffix)
	}

	// Test with suffix type option - generational
	generationalSuffix := Suffix(WithSuffixType(SuffixTypeGenerational))
	expectedGenerationalSuffix := "VIII"

	if generationalSuffix != expectedGenerationalSuffix {
		t.Errorf("Suffix(WithSuffixType(SuffixTypeGenerational)) = %v, want %v", generationalSuffix, expectedGenerationalSuffix)
	}

	// Test with locale option
	localeSuffix := Suffix(WithLocale("en"))
	expectedLocaleSuffix := "I"

	if localeSuffix != expectedLocaleSuffix {
		t.Errorf("Suffix(WithLocale(\"en\")) = %v, want %v", localeSuffix, expectedLocaleSuffix)
	}
}

// TestJobArea tests the JobArea function
func TestJobArea(t *testing.T) {
	setupTest(t)
	defer teardownTest(t)

	// Test multiple calls to ensure reproducibility
	jobArea1 := JobArea()
	jobArea2 := JobArea()

	// With a fixed seed, we should get consistent results
	expectedJobArea1 := "Education"
	expectedJobArea2 := "Human Resources"

	if jobArea1 != expectedJobArea1 {
		t.Errorf("JobArea() = %v, want %v", jobArea1, expectedJobArea1)
	}

	if jobArea2 != expectedJobArea2 {
		t.Errorf("JobArea() second call = %v, want %v", jobArea2, expectedJobArea2)
	}

	// Test with locale option
	localeJobArea := JobArea(WithLocale("en"))
	expectedLocaleJobArea := "Education"

	if localeJobArea != expectedLocaleJobArea {
		t.Errorf("JobArea(WithLocale(\"en\")) = %v, want %v", localeJobArea, expectedLocaleJobArea)
	}
}

// TestJobDescriptor tests the JobDescriptor function
func TestJobDescriptor(t *testing.T) {
	setupTest(t)
	defer teardownTest(t)

	// Test multiple calls to ensure reproducibility
	jobDescriptor1 := JobDescriptor()
	jobDescriptor2 := JobDescriptor()

	// With a fixed seed, we should get consistent results
	expectedJobDescriptor1 := "Analyst"
	expectedJobDescriptor2 := "Executive"

	if jobDescriptor1 != expectedJobDescriptor1 {
		t.Errorf("JobDescriptor() = %v, want %v", jobDescriptor1, expectedJobDescriptor1)
	}

	if jobDescriptor2 != expectedJobDescriptor2 {
		t.Errorf("JobDescriptor() second call = %v, want %v", jobDescriptor2, expectedJobDescriptor2)
	}

	// Test with locale option
	localeJobDescriptor := JobDescriptor(WithLocale("en"))
	expectedLocaleJobDescriptor := "Analyst"

	if localeJobDescriptor != expectedLocaleJobDescriptor {
		t.Errorf("JobDescriptor(WithLocale(\"en\")) = %v, want %v", localeJobDescriptor, expectedLocaleJobDescriptor)
	}
}

// TestJobType tests the JobType function
func TestJobType(t *testing.T) {
	setupTest(t)
	defer teardownTest(t)

	// Test multiple calls to ensure reproducibility
	jobType1 := JobType()
	jobType2 := JobType()

	// With a fixed seed, we should get consistent results
	expectedJobType1 := "Agent"
	expectedJobType2 := "Specialist"

	if jobType1 != expectedJobType1 {
		t.Errorf("JobType() = %v, want %v", jobType1, expectedJobType1)
	}

	if jobType2 != expectedJobType2 {
		t.Errorf("JobType() second call = %v, want %v", jobType2, expectedJobType2)
	}

	// Test with locale option
	localeJobType := JobType(WithLocale("en"))
	expectedLocaleJobType := "Agent"

	if localeJobType != expectedLocaleJobType {
		t.Errorf("JobType(WithLocale(\"en\")) = %v, want %v", localeJobType, expectedLocaleJobType)
	}
}

// TestJobTitle tests the JobTitle function
func TestJobTitle(t *testing.T) {
	setupTest(t)
	defer teardownTest(t)

	// Test multiple calls to ensure reproducibility
	jobTitle1 := JobTitle()
	jobTitle2 := JobTitle()

	// With a fixed seed, we should get consistent results
	expectedJobTitle1 := "Analyst Human Resources Agent"
	expectedJobTitle2 := "Strategist Healthcare Executive"

	if jobTitle1 != expectedJobTitle1 {
		t.Errorf("JobTitle() = %v, want %v", jobTitle1, expectedJobTitle1)
	}

	if jobTitle2 != expectedJobTitle2 {
		t.Errorf("JobTitle() second call = %v, want %v", jobTitle2, expectedJobTitle2)
	}

	// Test with locale option
	localeJobTitle := JobTitle(WithLocale("en"))
	expectedLocaleJobTitle := "Head Consulting Strategist"

	if localeJobTitle != expectedLocaleJobTitle {
		t.Errorf("JobTitle(WithLocale(\"en\")) = %v, want %v", localeJobTitle, expectedLocaleJobTitle)
	}
}

// BenchmarkFirstName benchmarks the FirstName function
func BenchmarkFirstName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FirstName()
	}
}

// BenchmarkMiddleName benchmarks the MiddleName function
func BenchmarkMiddleName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MiddleName()
	}
}

// BenchmarkLastName benchmarks the LastName function
func BenchmarkLastName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LastName()
	}
}

// BenchmarkGender benchmarks the Gender function
func BenchmarkGender(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Gender()
	}
}

// BenchmarkPrefix benchmarks the Prefix function
func BenchmarkPrefix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Prefix()
	}
}

// BenchmarkSuffix benchmarks the Suffix function
func BenchmarkSuffix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Suffix()
	}
}

// BenchmarkJobArea benchmarks the JobArea function
func BenchmarkJobArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JobArea()
	}
}

// BenchmarkJobDescriptor benchmarks the JobDescriptor function
func BenchmarkJobDescriptor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JobDescriptor()
	}
}

// BenchmarkJobType benchmarks the JobType function
func BenchmarkJobType(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JobType()
	}
}

// BenchmarkJobTitle benchmarks the JobTitle function
func BenchmarkJobTitle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JobTitle()
	}
}
