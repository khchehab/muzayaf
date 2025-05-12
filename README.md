# Muzayaf

Muzayaf is a feature-rich data generation library for Go. It provides a comprehensive set of tools for generating random test data such as colors, dates, numbers, and person-related information, with support for localization and customization.

[![Go Report Card](https://goreportcard.com/badge/github.com/khchehab/muzayaf)](https://goreportcard.com/report/github.com/khchehab/muzayaf)
[![Go Reference](https://pkg.go.dev/badge/github.com/khchehab/muzayaf.svg)](https://pkg.go.dev/github.com/khchehab/muzayaf)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Table of Contents

- [Installation](#installation)
- [Quick Start](#quick-start)
- [Available Packages](#available-packages)
- [Package Documentation](#package-documentation)
  - [Color](#color)
  - [Date](#date)
  - [Number](#number)
  - [Person](#person)
- [Customization](#customization)
  - [Localization](#localization)
  - [Random Source](#random-source)
- [Examples](#examples)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

## Installation

To install Muzayaf, use `go get`:

```bash
go get github.com/khchehab/muzayaf
```

## Quick Start

```go
import (
    "github.com/khchehab/muzayaf/color"
    "github.com/khchehab/muzayaf/date"
    "github.com/khchehab/muzayaf/number"
    "github.com/khchehab/muzayaf/person"
)

func main() {
    // Generate random data
    randomColor := color.RGB()
    randomDate := date.Any()
    randomNumber := number.Int()
    randomName := person.FirstName()
    
    // Use the generated data
    fmt.Println("Color:", randomColor.String())
    fmt.Println("Date:", randomDate.Format("2006-01-02"))
    fmt.Println("Number:", randomNumber)
    fmt.Println("Name:", randomName)
}
```

## Available Packages

Muzayaf consists of several sub-packages, each focused on specific types of data generation:

| Package | Description | Import Path |
|---------|-------------|-------------|
| **[color](https://pkg.go.dev/github.com/khchehab/muzayaf/color)** | Generate random colors in RGB, HSL, CMYK formats and named colors | `github.com/khchehab/muzayaf/color` |
| **[date](https://pkg.go.dev/github.com/khchehab/muzayaf/date)** | Generate random dates, times, months, weekdays, and timezones | `github.com/khchehab/muzayaf/date` |
| **[number](https://pkg.go.dev/github.com/khchehab/muzayaf/number)** | Generate random numbers in various formats (int, float, binary, hex, roman) | `github.com/khchehab/muzayaf/number` |
| **[person](https://pkg.go.dev/github.com/khchehab/muzayaf/person)** | Generate random person data (names, genders, job titles) | `github.com/khchehab/muzayaf/person` |

## Package Documentation

### Color

The `color` package provides functionality for generating random colors in different color spaces.

#### Features:

- RGB/RGBA colors
- HSL/HSLA colors
- CMYK colors
- Named colors (with localization support)

#### Example:

```go
// Generate a random RGB color
rgb := color.RGB()
fmt.Println("RGB color:", rgb.String()) // e.g., "rgb(120, 30, 215)"

// Generate a random RGBA color with transparency
rgba := color.RGBA()
fmt.Println("RGBA color:", rgba.String()) // e.g., "rgba(120, 30, 215, 0.75)"

// Generate a random HSL color
hsl := color.HSL()
fmt.Println("HSL color:", hsl.String()) // e.g., "hsl(270°, 75%, 48%)"

// Generate a random CMYK color
cmyk := color.CMYK()
fmt.Println("CMYK color:", cmyk.String()) // e.g., "cmyk(45%, 87%, 10%, 20%)"

// Generate a random color name
colorName := color.ColorName()
fmt.Println("Color name:", colorName) // e.g., "darkseagreen"

// Generate a color name with specific locale
enColorName := color.ColorName(color.WithLocale("en"))
fmt.Println("English color name:", enColorName)
```

### Date

The `date` package provides functionality for generating random dates, times, and date-related strings.

#### Features:

- Random dates (past, future, or in a specific range)
- Month names
- Weekday names
- Timezones

#### Example:

```go
// Generate any random date
anyDate := date.Any()
fmt.Println("Random date:", anyDate.Format("2006-01-02"))

// Generate a date between two specific dates
from := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
to := time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC)
betweenDate := date.Between(from, to)
fmt.Println("Date between 2020-2023:", betweenDate.Format("2006-01-02"))

// Generate a date in the future (with options)
futureDate := date.Future(date.WithYears(5)) // Maximum 5 years in the future
fmt.Println("Future date:", futureDate.Format("2006-01-02"))

// Generate a date in the past
pastDate := date.Past()
fmt.Println("Past date:", pastDate.Format("2006-01-02"))

// Generate a random month name
month := date.Month()
fmt.Println("Random month:", month) // e.g., "September"

// Generate a random weekday
weekday := date.Weekday()
fmt.Println("Random weekday:", weekday) // e.g., "Wednesday"

// Generate a random timezone
timezone := date.Timezone()
fmt.Println("Random timezone:", timezone) // e.g., "America/New_York"
```

### Number

The `number` package provides functionality for generating random numbers in various formats.

#### Features:

- Integers (with options for range and multiples)
- Floating-point numbers (with precision control)
- Binary numbers
- Octal numbers
- Hexadecimal numbers
- Roman numerals

#### Example:

```go
// Generate a random integer
randomInt := number.Int()
fmt.Println("Random integer:", randomInt)

// Generate an integer in a specific range
boundedInt := number.Int(number.WithIntMin(10), number.WithIntMax(100))
fmt.Println("Integer between 10-100:", boundedInt)

// Generate an integer that's a multiple of 5
multipleInt := number.Int(number.WithIntMin(10), number.WithIntMax(100), number.WithIntMultiple(5))
fmt.Println("Multiple of 5 between 10-100:", multipleInt) // e.g., 15, 20, 25...

// Generate a random float
randomFloat := number.Float()
fmt.Println("Random float:", randomFloat) // Between 0.0 and 1.0

// Generate a float with specific precision
preciseFloat := number.Float(number.WithFloatFractionDigits(4))
fmt.Println("Float with 4 decimal places:", preciseFloat)

// Generate a random binary number
binaryNumber := number.Binary(number.WithBinaryPrefix(true))
fmt.Println("Binary number:", binaryNumber) // e.g., "0b10110"

// Generate a random hexadecimal number
hexNumber := number.Hex(number.WithHexPrefix(true))
fmt.Println("Hex number:", hexNumber) // e.g., "0x1a3f"

// Generate a random roman numeral
romanNumber := number.Roman(number.WithRomanMin(1), number.WithRomanMax(100))
fmt.Println("Roman numeral:", romanNumber) // e.g., "XLII" (42)
```

### Person

The `person` package provides functionality for generating random person-related data.

#### Features:

- First names (with gender filtering)
- Middle names (with gender filtering)
- Last names
- Genders
- Prefixes (with gender filtering)
- Suffixes (academic or generational)
- Job titles and related information

#### Example:

```go
// Generate a random gender
gender := person.Gender()
fmt.Println("Gender:", gender) // e.g., "male" or "female"

// Generate random first name
firstName := person.FirstName()
fmt.Println("First name:", firstName)

// Generate gender-specific first name
maleFirstName := person.FirstName(person.WithGender(person.GenderMale))
femaleFirstName := person.FirstName(person.WithGender(person.GenderFemale))
fmt.Println("Male first name:", maleFirstName)
fmt.Println("Female first name:", femaleFirstName)

// Generate a random last name
lastName := person.LastName()
fmt.Println("Last name:", lastName)

// Generate a random prefix
prefix := person.Prefix()
fmt.Println("Prefix:", prefix) // e.g., "Mr.", "Mrs.", "Dr."

// Generate a random suffix (academic or generational)
academicSuffix := person.Suffix(person.WithSuffixType(person.SuffixTypeAcademic))
fmt.Println("Academic suffix:", academicSuffix) // e.g., "PhD", "MD"

generationalSuffix := person.Suffix(person.WithSuffixType(person.SuffixTypeGenerational))
fmt.Println("Generational suffix:", generationalSuffix) // e.g., "Jr.", "Sr.", "III"

// Generate a job title
jobTitle := person.JobTitle()
fmt.Println("Job title:", jobTitle) // e.g., "Senior Marketing Manager"

// Creating a complete person example
fmt.Printf("Full name: %s %s %s %s, %s\n", 
    person.Prefix(), 
    person.FirstName(), 
    person.LastName(), 
    person.Suffix(),
    person.JobTitle())
// e.g., "Dr. Jane Smith PhD, Lead Marketing Director"
```

## Customization

### Localization

Most data generation functions support localization through the `WithLocale` option. Currently, the library primarily supports English ("en") locale with the framework in place to add more locales.

```go
// Generate a month name in English
englishMonth := date.Month(date.WithLocale("en"))
fmt.Println("English month:", englishMonth)

// Generate a first name using English names
englishName := person.FirstName(person.WithLocale("en"))
fmt.Println("English name:", englishName)
```

### Random Source

You can customize the random source used by the library for deterministic output (useful for testing):

```go
import (
    "math/rand/v2"
    "github.com/khchehab/muzayaf/random"
)

// Set a custom random source with a fixed seed
src := rand.NewPCG(42, 43)
random.SetRandomSource(src)

// Now all random generation will use this source
color1 := color.RGB() // Will always produce the same color with the same seed

// Reset to the default time-based random source
random.ResetRandomSource()
```

## Examples

Complete examples of using each package can be found in the `example` directory:

- `example/color.go` - Examples of color generation
- `example/date.go` - Examples of date generation
- `example/number.go` - Examples of number generation
- `example/person.go` - Examples of person data generation

To run the examples:

```bash
go run example/main.go
```

## Testing

The library includes comprehensive tests for all packages. To run the tests:

```bash
go test -v ./...
```

Each package also includes benchmarks for performance testing:

```bash
go test -bench=. ./...
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.