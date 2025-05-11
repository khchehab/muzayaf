package date

import "time"

// Option struct holds configuration for date data generation
type Option struct {
	locale   string
	relative time.Time
	years    int
	months   int
	days     int
}

// OptionFunc is a function that modifies an Option
type OptionFunc func(*Option)

// defaultOption returns the default configuration
func defaultOption() Option {
	return Option{
		locale:   "en",
		relative: time.Now(),
		years:    100,
		months:   0,
		days:     0,
	}
}

// applyOptions applies the provided option functions to the default options
func applyOptions(opts []OptionFunc) Option {
	opt := defaultOption()
	for _, o := range opts {
		o(&opt)
	}
	return opt
}

// WithLocale sets the locale for date data generation
func WithLocale(locale string) OptionFunc {
	return func(o *Option) {
		o.locale = locale
	}
}

// WithRelative sets the relative date for date data generation
func WithRelative(relative time.Time) OptionFunc {
	return func(o *Option) {
		o.relative = relative
	}
}

// WithYears sets the number of years to go back or forward
func WithYears(years int) OptionFunc {
	return func(o *Option) {
		o.years = years
	}
}

// WithMonths sets the number of months to go back or forward
func WithMonths(months int) OptionFunc {
	return func(o *Option) {
		o.months = months
	}
}

// WithDays sets the number of days to go back or forward
func WithDays(days int) OptionFunc {
	return func(o *Option) {
		o.days = days
	}
}
