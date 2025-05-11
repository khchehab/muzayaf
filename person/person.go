// Package person provides functionality for generating random person-related data
// such as names, genders, prefixes, and suffixes.
package person

const (
	// GenderFemale represents the female gender identifier
	GenderFemale = "female"
	// GenderMale represents the male gender identifier
	GenderMale = "male"

	// SuffixTypeAcademic represents academic suffixes (e.g., Ph.D., M.D.)
	SuffixTypeAcademic = "academic"
	// SuffixTypeGenerational represents generational suffixes (e.g., Jr., Sr., III)
	SuffixTypeGenerational = "generational"
)

var (
	fallbackValues = map[string]map[string]string{
		"en": {
			"first_name":     "John",
			"middle_name":    "Smith",
			"last_name":      "Doe",
			"gender":         GenderFemale,
			"prefix":         "Dr.",
			"suffix":         "Sr.",
			"job_area":       "Sales",
			"job_descriptor": "Lead",
			"job_type":       "Officer",
		},
	}
)
