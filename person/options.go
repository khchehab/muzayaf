package person

// Option struct holds configuration for person data generation
type Option struct {
	locale     string
	gender     string
	suffixType string
}

// OptionFunc is a function that modifies an Option
type OptionFunc func(*Option)

// defaultOption returns the default configuration
func defaultOption() Option {
	return Option{
		locale:     "en",
		gender:     "",
		suffixType: "",
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

// WithLocale sets the locale for person data generation
func WithLocale(locale string) OptionFunc {
	return func(o *Option) {
		o.locale = locale
	}
}

// WithGender sets the gender for person data generation
func WithGender(gender string) OptionFunc {
	return func(o *Option) {
		o.gender = gender
	}
}

// WithSuffixType sets the suffix type for person data generation
func WithSuffixType(suffixType string) OptionFunc {
	return func(o *Option) {
		o.suffixType = suffixType
	}
}
