# Opgine - the functional option engine
Opgine (pronounced like `engine`) is a small engine for functional options similar to command line argument parsers

## Quickstart
The general idea is that you would use opgine for any function options that are TRULY optional. A good example is a db query where you might want all results (no predicate) or filtered results (n predicates). Take the following example:

```go
package main
import (
    "github.com/lusis/opgine"
)

type foo struct {}
type storer interface {
    Find(context.Context, opts ...opgine.Option) (*foo, error)
}

type storeimpl struct {}
func (s *storeimpl) Find(ctx context.Context, opts ...opgine.Option) {
    o, oerr := opgine.New(opts...)
    if oerr != nil {
        return nil, fmt.Errorf("error reading arguments: %w", oerr)
    }
    if o.Check("myboolkey") { // not required since GetBool will return an appropriate error if not set
        val, valerr := o.GetBool("myboolkey")
        if valerr != nil {
            return nil, fmt.Errorf("error getting value: %w", valerr)
        }
        // add the value to your db predicates
    }
    stringval := o.MustGetString("mystringkey") // can panic
    // add the value to your db predicates
    return &foo{},nil
}

func main() {
    impl := &storeimpl{}
    res, err := impl.Find(context.TODO(), opgine.WithBool("myboolkey",true),opgine.WithString("mystringkey","mystringval"))
    if err != nil {
        log.Fatal(err)
    }
    // something with res
}
```

## Supported options
All options generally support the following pattern:
`opgine.WithType(key, val)`
i.e.
`opgine.WithFloat64("myfloat",0.000000000003)`

To parse any options, you should call:

`og, err := opgine.New(opts...)`

As with other uses of functional option approaches, the only safe way to parse/populate is at construction time. Opgine does not offer a way to set a value after the fact and opgine uses a `sync.RWMutex` to ensure thread safety.

To get values that are passed in there are three mechanisms:
- `opgine.GetType(key)` (i.e. `opgine.GetString(key)`): this can return a nuanced error that you can check against
- `opgine.MustGetType(key)` (i.e. `opgine.MustGetString(key)`): this will always return the set value or panic if it cannot.
- `opgine.Get(key)`: this is a type-indifferent getter that can return an error. This is useful when passing an option via `opgine.WithInterface(key, val)`

Additionally, you can also set two non-value options which can be useful for understanding the relationship between the options:

- `opgine.WithAny()`: retrieved via `opgine.RequireAny()`
- `opgine.WithAll()`: retrieved via `opgine.RequireAll()`

These two options are mutually exclusive and require at least one other option to be set as well as `RequireAny()` doesn't really make sense standing alone. 

## Supported Types
- `int`,`int8`,`int16`,`int32`,`int64`
- `uint`,`uint8`,`uint16`,`uint32`,`uint64`
- `float32`,`float64`
- `string`
- `bool`
- `interface{}`
- `chan`,`func` via `opgine.WithInterface()` for now
- slices of any `int/uint`,`intX/uintX`, `string`, `float32`,`float64`

## Creating your own options
This is optional (hah!) but you MIGHT want to wrap your `opgine.Option` in your own type for flexibility and slightly more clean code on your side. You can do something like:

```go
// Package main ...
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/lusis/opgine"
)
type myOption = opgine.Option

// WithUsername ...
func WithUsername(s string) myOption {
	return opgine.WithString("username", s)
}

type foo struct{}

func (f *foo) Find(_ context.Context, opts ...myOption) {
	o, oerr := opgine.New(opts...)
	if oerr != nil {
		log.Fatal(oerr)
	}
	username := o.MustGet("username")
	fmt.Println(username)
}
func main() {
	f := &foo{}
	f.Find(context.TODO(), WithUsername("bob"))
}
```

I'm still working on a clean way to wrap hydrating the option.

## Background
I really love functional options in Go as a pattern. Once you grok how they work, it's a pattern that can make maintaining code easier over time as you don't have to constantly modify function signatures. If something is optional, you can either call:

```go
type foo struct {}
func (f *foo) Do(ctx context.Context, requiredstring string, optionalint int, optionalbool bool){
    // blah blah
}
res, err := foo.Do(ctx, "stringval", nil, false) // oh wait do I use false here? Does false mean actually a false value or unset value?)
```

or you can do

```go
type foo struct {}
func (f *foo) Do(ctx context.Context, requiredstring string, opts ...opgine.Option){
    // check if the bool is actually set and use value
    // check if the int is actually set and use value
}
res, err := foo.Do(ctx, "stringval")
// or
res, err := foo.Do(ctx, "stringval",opgine.WithBool("thing-that-can-be-true-false-or-not-set",false))
```

I use this pattern the most when writing interfaces for storers as it helps me keep the interface definition cleaner.
In the past in several projects and codebases, I've done more specific implemtations of this like so:

```go
foo.Do(ctx, options.WithUsername("username"),options.WithCreatedBefore(time.Now()))
```

and decided to try and build a reusable version. It also plays nice with go generics as well.

Take the following:

```go
type GenericStorer[T any] interface {
        // Get by id
        Get(ctx context.Context, id string) (*T, error)
        // Insert the thing
        Insert(ctx context.Context, thing *T) (*T, error)
        // Delete deletes the thing
        Delete(ctx context.Context, id string) error
        // Update updates a thing 
        Update(ctx context.Context, id string, opts ...opgine.Option) (*T, error)
        // Find find the thing by criteria
        Find(ctx context.Context, opts ...opgine.Option) ([]*T, error)
}
```

Yes, it might violate a least-surprise principal if different implemenations need different options so I tend to avoid that.

### Why not `map[string]interface{}`?
Sure you could do this but you'd also have to implement all the presence, type check and value check options yourself inline with the implementation. If you write multiple implementations, shaving off even that logic can really help readability.

And in the end, if you don't NEED to pass options, your function call is still going to look "ugly" with `nil` littered all over the place:

```go
res, err := foo.Find(ctx,nil) // give me all
// vs
res, err :=foo.Find(ctx, map[string]interface{}{"username":"susan"})
```

## TODO
- [ ] Add struct tag support for flexibility
- [ ] Add protobuf implementation (i.e. `repeated opgine.v1.Option options = 1;`)