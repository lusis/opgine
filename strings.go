package opgine

// WithString stores the provided string value at the provided key
func WithString(key, val string) Option {
	return withVal[string](key, val)
}

// GetString gets the string value at the named key
func (o *Opgine) GetString(key string) (string, error) {
	return getVal[string](o, key)
}

// MustGetString gets the string value at the named key panicking if not set
func (o *Opgine) MustGetString(key string) string {
	return mustGetVal[string](o, key)
}
