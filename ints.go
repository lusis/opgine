package opgine

// WithInt64 stores the provided int64 value at the provided key
func WithInt64(key string, val int64) Option {
	return withVal[int64](key, val)
}

// WithInt32 stores the provided int32 value at the provided key
func WithInt32(key string, val int32) Option {
	return withVal[int32](key, val)
}

// WithInt16 stores the provided int16 value at the provided key
func WithInt16(key string, val int16) Option {
	return withVal[int16](key, val)
}

// WithInt8 stores the provided int8 value at the provided key
func WithInt8(key string, val int8) Option {
	return withVal[int8](key, val)
}

// WithInt stores the provided int value at the provided key
func WithInt(key string, val int) Option {
	return withVal[int](key, val)
}

// WithUint64 stores the provided uint64 value at the provided key
func WithUint64(key string, val uint64) Option {
	return withVal[uint64](key, val)
}

// WithUint32 stores the provided uint32 value at the provided key
func WithUint32(key string, val uint32) Option {
	return withVal[uint32](key, val)
}

// WithUint16 stores the provided uint16 value at the provided key
func WithUint16(key string, val uint16) Option {
	return withVal[uint16](key, val)
}

// WithUint8 stores the provided uint8 value at the provided key
func WithUint8(key string, val uint8) Option {
	return withVal[uint8](key, val)
}

// WithUint stores the provided uint value at the provided key
func WithUint(key string, val uint) Option {
	return withVal[uint](key, val)
}

// GetInt gets the value at the provided key
func (o *Opgine) GetInt(key string) (int, error) {
	return getVal[int](o, key)
}

// MustGetInt gets the value at the provided key
func (o *Opgine) MustGetInt(key string) int {
	return mustGetVal[int](o, key)
}

// GetInt8 gets the value at the provided key
func (o *Opgine) GetInt8(key string) (int8, error) {
	return getVal[int8](o, key)
}

// MustGetInt8 gets the value at the provided key
func (o *Opgine) MustGetInt8(key string) int8 {
	return mustGetVal[int8](o, key)
}

// GetInt16 gets the value at the provided key
func (o *Opgine) GetInt16(key string) (int16, error) {
	return getVal[int16](o, key)
}

// MustGetInt16 gets the value at the provided key
func (o *Opgine) MustGetInt16(key string) int16 {
	return mustGetVal[int16](o, key)
}

// GetInt32 gets the value at the provided key
func (o *Opgine) GetInt32(key string) (int32, error) {
	return getVal[int32](o, key)
}

// MustGetInt32 gets the value at the provided key
func (o *Opgine) MustGetInt32(key string) int32 {
	return mustGetVal[int32](o, key)
}

// GetInt64 gets the value at the provided key
func (o *Opgine) GetInt64(key string) (int64, error) {
	return getVal[int64](o, key)
}

// MustGetInt64 gets the value at the provided key
func (o *Opgine) MustGetInt64(key string) int64 {
	return mustGetVal[int64](o, key)
}

// GetUint gets the value at the provided key
func (o *Opgine) GetUint(key string) (uint, error) {
	return getVal[uint](o, key)
}

// MustGetUint gets the value at the provided key
func (o *Opgine) MustGetUint(key string) uint {
	return mustGetVal[uint](o, key)
}

// GetUint8 gets the value at the provided key
func (o *Opgine) GetUint8(key string) (uint8, error) {
	return getVal[uint8](o, key)
}

// MustGetUint8 gets the value at the provided key
func (o *Opgine) MustGetUint8(key string) uint8 {
	return mustGetVal[uint8](o, key)
}

// GetUint16 gets the value at the provided key
func (o *Opgine) GetUint16(key string) (uint16, error) {
	return getVal[uint16](o, key)
}

// MustGetUint16 gets the value at the provided key
func (o *Opgine) MustGetUint16(key string) uint16 {
	return mustGetVal[uint16](o, key)
}

// GetUint32 gets the value at the provided key
func (o *Opgine) GetUint32(key string) (uint32, error) {
	return getVal[uint32](o, key)
}

// MustGetUint32 gets the value at the provided key
func (o *Opgine) MustGetUint32(key string) uint32 {
	return mustGetVal[uint32](o, key)
}

// GetUint64 gets the value at the provided key
func (o *Opgine) GetUint64(key string) (uint64, error) {
	return getVal[uint64](o, key)
}

// MustGetUint64 gets the value at the provided key
func (o *Opgine) MustGetUint64(key string) uint64 {
	return mustGetVal[uint64](o, key)
}
