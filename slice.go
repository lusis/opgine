package opgine

// WithStringSlice stores the provided string slice value at the provided key
// slices must contain at least one element
func WithStringSlice(key string, val []string) Option {
	return withVal[[]string](key, val)
}

// WithUintSlice stores the provided slice value at the provided key
// slices must contain at least one element
func WithUintSlice(key string, val []uint) Option {
	return withVal[[]uint](key, val)
}

// WithUint8Slice stores the provided slice value at the provided key
// slices must contain at least one element
func WithUint8Slice(key string, val []uint8) Option {
	return withVal[[]uint8](key, val)
}

// WithUint16Slice stores the provided slice value at the provided key
// slices must contain at least one element
func WithUint16Slice(key string, val []uint16) Option {
	return withVal[[]uint16](key, val)
}

// WithUint32Slice stores the provided slice value at the provided key
// slices must contain at least one element
func WithUint32Slice(key string, val []uint32) Option {
	return withVal[[]uint32](key, val)
}

// WithUint64Slice stores the provided slice value at the provided key
// slices must contain at least one element
func WithUint64Slice(key string, val []uint64) Option {
	return withVal[[]uint64](key, val)
}

// WithIntSlice stores the provided slice value at the provided key
// slices must contain at least one element
func WithIntSlice(key string, val []int) Option {
	return withVal[[]int](key, val)
}

// WithInt8Slice stores the provided slice value at the provided key
// slices must contain at least one element
func WithInt8Slice(key string, val []int8) Option {
	return withVal[[]int8](key, val)
}

// WithInt16Slice stores the provided slice value at the provided key
// slices must contain at least one element
func WithInt16Slice(key string, val []int16) Option {
	return withVal[[]int16](key, val)
}

// WithInt32Slice stores the provided slice value at the provided key
// slices must contain at least one element
func WithInt32Slice(key string, val []int32) Option {
	return withVal[[]int32](key, val)
}

// WithInt64Slice stores the provided slice value at the provided key
// slices must contain at least one element
func WithInt64Slice(key string, val []int64) Option {
	return withVal[[]int64](key, val)
}

// WithFloat32Slice stores the provided slice value at the provided key
// slices must contain at least one element
func WithFloat32Slice(key string, val []float32) Option {
	return withVal[[]float32](key, val)
}

// WithFloat64Slice stores the provided slice value at the provided key
// slices must contain at least one element
func WithFloat64Slice(key string, val []float64) Option {
	return withVal[[]float64](key, val)
}

// GetStringSlice gets the value at the provided key
func (o *Opgine) GetStringSlice(key string) ([]string, error) {
	return getVal[[]string](o, key)
}

// MustGetStringSlice gets the value at the provided key panicking if not set
func (o *Opgine) MustGetStringSlice(key string) []string {
	return mustGetVal[[]string](o, key)
}

// GetIntSlice gets the value at the provided key
func (o *Opgine) GetIntSlice(key string) ([]int, error) {
	return getVal[[]int](o, key)
}

// MustGetIntSlice gets the value at the provided key panicking if not set
func (o *Opgine) MustGetIntSlice(key string) []int {
	return mustGetVal[[]int](o, key)
}

// GetInt8Slice gets the value at the provided key
func (o *Opgine) GetInt8Slice(key string) ([]int8, error) {
	return getVal[[]int8](o, key)
}

// MustGetInt8Slice gets the value at the provided key panicking if not set
func (o *Opgine) MustGetInt8Slice(key string) []int8 {
	return mustGetVal[[]int8](o, key)
}

// GetInt16Slice gets the value at the provided key
func (o *Opgine) GetInt16Slice(key string) ([]int16, error) {
	return getVal[[]int16](o, key)
}

// MustGetInt16Slice gets the value at the provided key panicking if not set
func (o *Opgine) MustGetInt16Slice(key string) []int16 {
	return mustGetVal[[]int16](o, key)
}

// GetInt32Slice gets the value at the provided key
func (o *Opgine) GetInt32Slice(key string) ([]int32, error) {
	return getVal[[]int32](o, key)
}

// MustGetInt32Slice gets the value at the provided key panicking if not set
func (o *Opgine) MustGetInt32Slice(key string) []int32 {
	return mustGetVal[[]int32](o, key)
}

// GetInt64Slice gets the value at the provided key
func (o *Opgine) GetInt64Slice(key string) ([]int64, error) {
	return getVal[[]int64](o, key)
}

// MustGetInt64Slice gets the value at the provided key panicking if not set
func (o *Opgine) MustGetInt64Slice(key string) []int64 {
	return mustGetVal[[]int64](o, key)
}

// GetUintSlice gets the value at the provided key
func (o *Opgine) GetUintSlice(key string) ([]uint, error) {
	return getVal[[]uint](o, key)
}

// MustGetUintSlice gets the value at the provided key panicking if not set
func (o *Opgine) MustGetUintSlice(key string) []uint {
	return mustGetVal[[]uint](o, key)
}

// GetUint8Slice gets the value at the provided key
func (o *Opgine) GetUint8Slice(key string) ([]uint8, error) {
	return getVal[[]uint8](o, key)
}

// MustGetUint8Slice gets the value at the provided key panicking if not set
func (o *Opgine) MustGetUint8Slice(key string) []uint8 {
	return mustGetVal[[]uint8](o, key)
}

// GetUint16Slice gets the value at the provided key
func (o *Opgine) GetUint16Slice(key string) ([]uint16, error) {
	return getVal[[]uint16](o, key)
}

// MustGetUint16Slice gets the value at the provided key panicking if not set
func (o *Opgine) MustGetUint16Slice(key string) []uint16 {
	return mustGetVal[[]uint16](o, key)
}

// GetUint32Slice gets the value at the provided key
func (o *Opgine) GetUint32Slice(key string) ([]uint32, error) {
	return getVal[[]uint32](o, key)
}

// MustGetUint32Slice gets the value at the provided key panicking if not set
func (o *Opgine) MustGetUint32Slice(key string) []uint32 {
	return mustGetVal[[]uint32](o, key)
}

// GetUint64Slice gets the value at the provided key
func (o *Opgine) GetUint64Slice(key string) ([]uint64, error) {
	return getVal[[]uint64](o, key)
}

// MustGetUint64Slice gets the value at the provided key panicking if not set
func (o *Opgine) MustGetUint64Slice(key string) []uint64 {
	return mustGetVal[[]uint64](o, key)
}

// GetFloat32Slice gets the value at the provided key
func (o *Opgine) GetFloat32Slice(key string) ([]float32, error) {
	return getVal[[]float32](o, key)
}

// MustGetFloat32Slice gets the value at the provided key panicking if not set
func (o *Opgine) MustGetFloat32Slice(key string) []float32 {
	return mustGetVal[[]float32](o, key)
}

// GetFloat64Slice gets the value at the provided key
func (o *Opgine) GetFloat64Slice(key string) ([]float64, error) {
	return getVal[[]float64](o, key)
}

// MustGetFloat64Slice gets the value at the provided key panicking if not set
func (o *Opgine) MustGetFloat64Slice(key string) []float64 {
	return mustGetVal[[]float64](o, key)
}
