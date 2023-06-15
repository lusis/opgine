package opgine

import (
	"sync"
)

// Opgine stores options
type Opgine struct {
	m          sync.RWMutex
	optMap     map[string]any
	requireAll *bool
	requireAny *bool
}

// New returns a new [opgine.Opgine] configured with the provided [opgine.Option]s
// you should never create an [opgine.Opgine] via struct literal for thread safety reasons.
// always use [opgine.New]
func New(opts ...Option) (*Opgine, error) {
	if len(opts) == 0 {
		return nil, ErrAtLeastOne
	}
	og := &Opgine{m: sync.RWMutex{}, optMap: map[string]any{}}
	og.m.Lock()
	for _, opt := range opts {
		if err := opt(og); err != nil {
			og.m.Unlock()
			return nil, err
		}
	}
	og.m.Unlock()
	// we need to check if any options are set
	// but we also need to check if the only option is one of the modifier ones
	// since those two are unique and make no sense outside another option
	if len(opts) == 0 || ((og.requireAll != nil || og.requireAny != nil) && len(opts) == 1) {
		return nil, ErrAtLeastOne
	}

	return og, nil
}

// Check indicates if the value has been set or not
func (o *Opgine) Check(key string) bool {
	o.m.RLock()
	defer o.m.RUnlock()
	return o.isset(key)
}

// Get gets the value at the provided key
func (o *Opgine) Get(key string) (any, error) {
	return get(o, key)
}

// MustGet gets the value at the provided key panicking if not set
func (o *Opgine) MustGet(key string) any {
	v, err := get(o, key)
	if err != nil {
		panic(err)
	}
	return v
}

func get(o *Opgine, key string) (any, error) {
	o.m.RLock()
	defer o.m.RUnlock()
	if !o.isset(key) {
		return nil, ErrNotSet
	}
	return o.optMap[key], nil
}

func getVal[T any](o *Opgine, key string) (T, error) {
	var zero T
	val, err := get(o, key)
	if err != nil {
		return zero, err
	}
	v, ok := val.(T)
	if !ok {
		return zero, ErrMismatchedTypes
	}
	return v, nil
}

func mustGetVal[T any](o *Opgine, key string) T {
	val, err := get(o, key)
	if err != nil {
		panic(err)
	}
	v, ok := val.(T)
	if !ok {
		panic(ErrMismatchedTypes)
	}
	return v
}

func (o *Opgine) isset(key string) bool {
	_, ok := o.optMap[key]
	return ok
}
