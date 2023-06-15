package opgine

// WithAll indicates that all options are required
// can be used to hint db predicates for AND operations
func WithAll() Option {
	return func(o *Opgine) error {
		r := true
		if o.requireAny != nil {
			return ErrConflict
		}
		o.requireAll = &r
		return nil
	}
}

// WithAny indicates that any options are allowed
// can be used to hint db predicates for OR operations
func WithAny() Option {
	return func(o *Opgine) error {
		r := true
		if o.requireAll != nil {
			return ErrConflict
		}
		o.requireAny = &r
		return nil
	}
}

// RequiresAll checks to see if the option for requiring all values was set
func (o *Opgine) RequiresAll() bool {
	o.m.RLock()
	defer o.m.RUnlock()
	if o.requireAll == nil {
		return false
	}
	return *o.requireAll
}

// RequiresAny checks to see if the option for requiring any value is set
func (o *Opgine) RequiresAny() bool {
	o.m.RLock()
	defer o.m.RUnlock()
	if o.requireAny == nil {
		return false
	}
	return *o.requireAny
}
