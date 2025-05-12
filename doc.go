// Package muzayaf is a comprehensive data generation library for Go that provides
// functionality for generating random test data. It consists of several sub-packages,
// each focused on a specific type of data generation.
//
// # Packages
//
// The muzayaf library is organized into the following packages:
//
//   - color: Generate random colors in various formats (RGB, HSL, CMYK, named colors)
//   - date: Generate random dates, times, months, weekdays, and timezones
//   - number: Generate random numbers in different formats (integers, floats, binary, hex, roman)
//   - person: Generate random person data (names, genders, job titles, prefixes, suffixes)
//
// # Installation
//
// To install muzayaf, use:
//
//	go get github.com/khchehab/muzayaf
//
// # Usage
//
// Import the specific packages you need:
//
//	import (
//	    "github.com/khchehab/muzayaf/color"
//	    "github.com/khchehab/muzayaf/date"
//	    "github.com/khchehab/muzayaf/number"
//	    "github.com/khchehab/muzayaf/person"
//	)
//
// Then use the package functions to generate data:
//
//	// Generate a random color
//	rgb := color.RGB()
//	fmt.Println("RGB color:", rgb.String())
//
//	// Generate a random date
//	randomDate := date.Any()
//	fmt.Println("Date:", randomDate.Format("2006-01-02"))
//
//	// Generate a random number
//	randomInt := number.Int(number.WithIntMin(10), number.WithIntMax(100))
//	fmt.Println("Integer:", randomInt)
//
//	// Generate a random person
//	firstName := person.FirstName()
//	lastName := person.LastName()
//	fmt.Printf("Person: %s %s\n", firstName, lastName)
//
// For more detailed information about each package, see their individual documentation.
package muzayaf
