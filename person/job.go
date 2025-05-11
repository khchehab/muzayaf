// Package person provides functionality for generating random person-related data
// including job information such as job areas, descriptors, types, and titles.
package person

import (
	"fmt"
	"github.com/khchehab/muzayaf/internal"
	"github.com/khchehab/muzayaf/random"
)

// JobArea generates a random job area
func JobArea(opts ...OptionFunc) string {
	o := applyOptions(opts)

	data, err := internal.LoadJsonFile("person", o.locale, "job_areas.json")
	if err != nil {
		return fallbackValues[o.locale]["job_area"]
	}

	pool := internal.GetStringSlice(data, "areas")
	if len(pool) == 0 {
		return fallbackValues[o.locale]["job_area"]
	}

	return pool[random.IntN(len(pool))]
}

// JobDescriptor generates a random job descriptor
func JobDescriptor(opts ...OptionFunc) string {
	o := applyOptions(opts)

	data, err := internal.LoadJsonFile("person", o.locale, "job_descriptors.json")
	if err != nil {
		return fallbackValues[o.locale]["job_descriptor"]
	}

	pool := internal.GetStringSlice(data, "descriptors")
	if len(pool) == 0 {
		return fallbackValues[o.locale]["job_descriptor"]
	}

	return pool[random.IntN(len(pool))]
}

// JobType generates a random job type
func JobType(opts ...OptionFunc) string {
	o := applyOptions(opts)

	data, err := internal.LoadJsonFile("person", o.locale, "job_types.json")
	if err != nil {
		return fallbackValues[o.locale]["job_type"]
	}

	pool := internal.GetStringSlice(data, "types")
	if len(pool) == 0 {
		return fallbackValues[o.locale]["job_type"]
	}

	return pool[random.IntN(len(pool))]
}

// JobTitle generates a random job title by concatenating a job descriptor, job area, and job type
func JobTitle(opts ...OptionFunc) string {
	descriptor := JobDescriptor(opts...)
	area := JobArea(opts...)
	jobType := JobType(opts...)

	return fmt.Sprintf("%s %s %s", descriptor, area, jobType)
}
