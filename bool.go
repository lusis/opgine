package opgine

// WithBool stores the provided bool value at the provided key
func WithBool(key string, val bool) Option {
	return withVal[bool](key, val)
}

// GetBool gets the bool value at the named key
func (o *Opgine) GetBool(key string) (bool, error) {
	return getVal[bool](o, key)
}

// MustGetBool gets the value at the provided key panicking if not set
func (o *Opgine) MustGetBool(key string) bool {
	return mustGetVal[bool](o, key)
}
