package opgine

import (
	"reflect"
	"strings"
	"time"
)

// Option represents an option
type Option func(*Opgine) error

// WithInterface allows setting an unsupported option
func WithInterface(key string, val interface{}) Option {
	return withVal(key, val)
}

// withVal is a generic for setting the provided value at the provided key
func withVal[T any](key string, val T) Option {
	return func(o *Opgine) error {
		if strings.TrimSpace(key) == "" {
			return ErrInvalidKey
		}
		if o.isset(key) {
			return ErrAlreadySet
		}
		// we're going to do two types of checks here
		// if we can avoid reflecting we will
		switch v := any(val).(type) {
		case bool, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
			// there's no safe zero-value check for these so we just pass
			break
		case time.Time:
			if v.IsZero() {
				return ErrEmptyValue
			}
		case string:
			if strings.TrimSpace(v) == "" {
				return ErrEmptyValue
			}
		case nil:
			return ErrEmptyValue
		default:
			k := reflect.ValueOf(val)
			// more checks that require reflect
			switch k.Kind() {
			case reflect.Slice, reflect.Map:
				if k.IsNil() {
					return ErrEmptyValue
				}
				if k.Len() == 0 {
					return ErrAtLeastOne
				}
			case reflect.Chan, reflect.Func:
				if k.IsNil() {
					return ErrEmptyValue
				}
			}
		}
		o.optMap[key] = val

		return nil
	}
}
