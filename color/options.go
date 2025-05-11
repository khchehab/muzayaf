package color

// Option struct holds configuration for color data generation
type Option struct {
	locale string
}

// OptionFunc is a function that modifies an Option
type OptionFunc func(*Option)

// defaultOption returns the default configuration
func defaultOption() Option {
	return Option{
		locale: "en",
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

// WithLocale sets the locale for color data generation
func WithLocale(locale string) OptionFunc {
	return func(o *Option) {
		o.locale = locale
	}
}
