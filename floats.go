package opgine

// WithFloat64 stores the provided float64 value at the provided key
func WithFloat64(key string, val float64) Option {
	return withVal[float64](key, val)
}

// WithFloat32 stores the provided float32 value at the provided key
func WithFloat32(key string, val float32) Option {
	return withVal[float32](key, val)
}

// GetFloat32 gets the value at the provided key
func (o *Opgine) GetFloat32(key string) (float32, error) {
	return getVal[float32](o, key)
}

// MustGetFloat32 gets the value at the provided key
func (o *Opgine) MustGetFloat32(key string) float32 {
	return mustGetVal[float32](o, key)
}

// GetFloat64 gets the value at the provided key
func (o *Opgine) GetFloat64(key string) (float64, error) {
	return getVal[float64](o, key)
}

// MustGetFloat64 gets the value at the provided key
func (o *Opgine) MustGetFloat64(key string) float64 {
	return mustGetVal[float64](o, key)
}
