package opgine

import "fmt"

// ErrAlreadySet is the error when a value has already been set
var ErrAlreadySet = fmt.Errorf("value already set")

// ErrInvalidKey is the error when the provided key is invalid
var ErrInvalidKey = fmt.Errorf("invalid key name")

// ErrEmptyValue is the error when the provided value is empty
var ErrEmptyValue = fmt.Errorf("empty value not allowed")

// ErrNotSet is the error when attempting to get a key that is not set
var ErrNotSet = fmt.Errorf("value is not set")

// ErrMismatchedTypes is the error when a value is not the requested type
var ErrMismatchedTypes = fmt.Errorf("value is not of type")

// ErrConflict is the error when a value conflicts with another value
var ErrConflict = fmt.Errorf("options conflict")

// ErrAtLeastOne is the error returned when no slice values are provided
var ErrAtLeastOne = fmt.Errorf("at least one value must be provided")
