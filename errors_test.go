package opgine

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestErrors(t *testing.T) {
	t.Run("notset", func(t *testing.T) {
		o, err := New(withVal("mybool", false))
		require.NoError(t, err)
		require.NotNil(t, o)
		res, reserr := getVal[[]string](o, "invalidkey")
		require.ErrorIs(t, reserr, ErrNotSet)
		require.Nil(t, res)
	})
	type testCase struct {
		opts []Option
		err  error
	}
	testCases := map[string]testCase{
		"alreadyset": {
			opts: []Option{withVal("myboolval", false), withVal("myboolval", true)},
			err:  ErrAlreadySet,
		},
		"emptystring": {
			opts: []Option{WithString("mystringval", "")},
			err:  ErrEmptyValue,
		},
		"emptykey": {
			opts: []Option{withVal[string]("", "val")},
			err:  ErrInvalidKey,
		},
		"emptyslice": {
			opts: []Option{withVal[[]string]("myslice", []string{})},
			err:  ErrAtLeastOne,
		},
		"emptymap": {
			opts: []Option{withVal[map[string]interface{}]("mymap", map[string]interface{}{})},
			err:  ErrAtLeastOne,
		},
		"nilmap": {
			opts: []Option{withVal[map[string]interface{}]("mymap", nil)},
			err:  ErrEmptyValue,
		},
		"atleastone": {
			opts: []Option{},
			err:  ErrAtLeastOne,
		},
		"nilval": {
			opts: []Option{withVal[interface{}]("myval", nil)},
			err:  ErrEmptyValue,
		},
		"nilchan": {
			opts: []Option{withVal[chan (error)]("mychan", nil)},
			err:  ErrEmptyValue,
		},
		"niltime": {
			opts: []Option{withVal[time.Time]("mytime", time.Time{})},
			err:  ErrEmptyValue,
		},
		"onlyany": {
			opts: []Option{WithAny()},
			err:  ErrAtLeastOne,
		},
		"onlyall": {
			opts: []Option{WithAll()},
			err:  ErrAtLeastOne,
		},
		"conflicting-all-any": {
			opts: []Option{WithAll(), WithAny()},
			err:  ErrConflict,
		},
		"conflicting-any-all": {
			opts: []Option{WithAny(), WithAll()},
			err:  ErrConflict,
		},
	}
	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			o, err := New(tc.opts...)
			require.Equal(t, tc.err, err)
			require.ErrorIs(t, err, tc.err)
			require.Nil(t, o)

		})
	}
}
