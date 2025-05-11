package main

import (
	"fmt"
	"github.com/khchehab/muzayaf/person"
)

func mainPerson() {
	fmt.Println("Person Data Generation Examples")
	fmt.Println("==============================")

	// Generate random gender
	gender := person.Gender()
	fmt.Printf("Random Gender: %s\n\n", gender)

	// Generate random first name
	firstName := person.FirstName()
	fmt.Printf("Random First Name: %s\n", firstName)

	// Generate gender-specific first name
	maleFirstName := person.FirstName(person.WithGender(person.GenderMale))
	femaleFirstName := person.FirstName(person.WithGender(person.GenderFemale))
	fmt.Printf("Male First Name: %s\n", maleFirstName)
	fmt.Printf("Female First Name: %s\n\n", femaleFirstName)

	// Generate random last name
	lastName := person.LastName()
	fmt.Printf("Random Last Name: %s\n\n", lastName)

	// Generate random middle name
	middleName := person.MiddleName()
	fmt.Printf("Random Middle Name: %s\n", middleName)

	// Generate gender-specific middle name
	maleMiddleName := person.MiddleName(person.WithGender(person.GenderMale))
	femaleMiddleName := person.MiddleName(person.WithGender(person.GenderFemale))
	fmt.Printf("Male Middle Name: %s\n", maleMiddleName)
	fmt.Printf("Female Middle Name: %s\n\n", femaleMiddleName)

	// Generate random prefix
	prefix := person.Prefix()
	fmt.Printf("Random Prefix: %s\n", prefix)

	// Generate gender-specific prefix
	malePrefix := person.Prefix(person.WithGender(person.GenderMale))
	femalePrefix := person.Prefix(person.WithGender(person.GenderFemale))
	fmt.Printf("Male Prefix: %s\n", malePrefix)
	fmt.Printf("Female Prefix: %s\n\n", femalePrefix)

	// Generate random suffix
	suffix := person.Suffix()
	fmt.Printf("Random Suffix: %s\n", suffix)

	// Generate specific suffix types
	academicSuffix := person.Suffix(person.WithSuffixType(person.SuffixTypeAcademic))
	generationalSuffix := person.Suffix(person.WithSuffixType(person.SuffixTypeGenerational))
	fmt.Printf("Academic Suffix: %s\n", academicSuffix)
	fmt.Printf("Generational Suffix: %s\n\n", generationalSuffix)

	// Generate a complete person
	fmt.Println("Complete Person Examples:")
	fmt.Println("------------------------")

	// Example 1: Random person
	randomGender := person.Gender()
	fmt.Printf("Person 1: %s %s %s %s %s\n",
		person.Prefix(person.WithGender(randomGender)),
		person.FirstName(person.WithGender(randomGender)),
		person.MiddleName(person.WithGender(randomGender)),
		person.LastName(),
		person.Suffix(),
	)

	// Example 2: Male person with academic suffix
	fmt.Printf("Person 2: %s %s %s %s %s\n",
		person.Prefix(person.WithGender(person.GenderMale)),
		person.FirstName(person.WithGender(person.GenderMale)),
		person.MiddleName(person.WithGender(person.GenderMale)),
		person.LastName(),
		person.Suffix(person.WithSuffixType(person.SuffixTypeAcademic)),
	)

	// Example 3: Female person with generational suffix
	fmt.Printf("Person 3: %s %s %s %s %s\n",
		person.Prefix(person.WithGender(person.GenderFemale)),
		person.FirstName(person.WithGender(person.GenderFemale)),
		person.MiddleName(person.WithGender(person.GenderFemale)),
		person.LastName(),
		person.Suffix(person.WithSuffixType(person.SuffixTypeGenerational)),
	)

	// Job Examples
	fmt.Println("\nJob Examples:")
	fmt.Println("------------")

	// Generate random job area
	jobArea := person.JobArea()
	fmt.Printf("Random Job Area: %s\n", jobArea)

	// Generate random job descriptor
	jobDescriptor := person.JobDescriptor()
	fmt.Printf("Random Job Descriptor: %s\n", jobDescriptor)

	// Generate random job type
	jobType := person.JobType()
	fmt.Printf("Random Job Type: %s\n", jobType)

	// Generate random job title
	jobTitle := person.JobTitle()
	fmt.Printf("Random Job Title: %s\n", jobTitle)

	// Generate job information with locale
	fmt.Printf("Job Title with Locale: %s\n", person.JobTitle(person.WithLocale("en")))

	// Example: Person with job title
	personGender := person.Gender()
	fmt.Printf("\nPerson with Job: %s %s %s %s - %s\n",
		person.Prefix(person.WithGender(personGender)),
		person.FirstName(person.WithGender(personGender)),
		person.LastName(),
		person.Suffix(),
		person.JobTitle(),
	)
}
