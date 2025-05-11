package main

import (
	"fmt"
	"github.com/khchehab/muzayaf/date"
	"time"
)

func mainDate() {
	fmt.Println("Date Data Generation Examples")
	fmt.Println("============================")

	// Generate any random date
	anyDate := date.Any()
	fmt.Printf("Random Date: %s\n", anyDate.Format("2006-01-02"))

	// Generate any random date with a relative date
	relativeDate := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	anyRelativeDate := date.Any(date.WithRelative(relativeDate))
	fmt.Printf("Random Date (relative to 2000-01-01): %s\n\n", anyRelativeDate.Format("2006-01-02"))

	// Generate a random date between two dates
	from := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	to := time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC)
	betweenDate := date.Between(from, to)
	fmt.Printf("Random Date Between 2020-01-01 and 2023-12-31: %s\n\n", betweenDate.Format("2006-01-02"))

	// Generate a random date in the future
	futureDate := date.Future()
	fmt.Printf("Random Future Date: %s\n", futureDate.Format("2006-01-02"))

	// Generate a random date in the future with a relative date
	futureDateRelative := date.Future(date.WithRelative(relativeDate))
	fmt.Printf("Random Future Date (relative to 2000-01-01): %s\n\n", futureDateRelative.Format("2006-01-02"))

	// Generate a random date in the past
	pastDate := date.Past()
	fmt.Printf("Random Past Date: %s\n", pastDate.Format("2006-01-02"))

	// Generate a random date in the past with a relative date
	pastDateRelative := date.Past(date.WithRelative(relativeDate))
	fmt.Printf("Random Past Date (relative to 2000-01-01): %s\n\n", pastDateRelative.Format("2006-01-02"))

	// Generate a random month
	month := date.Month()
	fmt.Printf("Random Month: %s\n", month)

	// Generate a random month with a specific locale
	monthWithLocale := date.Month(date.WithLocale("en"))
	fmt.Printf("Random Month (en locale): %s\n\n", monthWithLocale)

	// Generate a random timezone
	timezone := date.Timezone()
	fmt.Printf("Random Timezone: %s\n\n", timezone)

	// Generate a random weekday
	weekday := date.Weekday()
	fmt.Printf("Random Weekday: %s\n", weekday)

	// Generate a random weekday with a specific locale
	weekdayWithLocale := date.Weekday(date.WithLocale("en"))
	fmt.Printf("Random Weekday (en locale): %s\n", weekdayWithLocale)
}
