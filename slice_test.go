package opgine_test

import (
	"testing"

	"github.com/lusis/opgine"
	"github.com/stretchr/testify/require"
)

func TestSlices(t *testing.T) {
	keyname := t.Name()
	testCases := map[string]testCase{
		"stringslice": {
			val: []string{"a", "b", "c", t.Name()},
			opt: opgine.WithStringSlice(keyname, []string{"a", "b", "c", t.Name()}),
			getfunc: func(o *opgine.Opgine) (any, error) {
				return o.GetStringSlice(keyname)
			},
			mustfunc: func(o *opgine.Opgine) any {
				return o.MustGetStringSlice(keyname)
			},
		},
		"intslice": {
			val: []int{-1, -2, 1, 2},
			opt: opgine.WithIntSlice(keyname, []int{-1, -2, 1, 2}),
			getfunc: func(o *opgine.Opgine) (any, error) {
				return o.GetIntSlice(keyname)
			},
			mustfunc: func(o *opgine.Opgine) any {
				return o.MustGetIntSlice(keyname)
			},
		},
		"uintslice": {
			val: []uint{1, 2, 3, 4},
			opt: opgine.WithUintSlice(keyname, []uint{1, 2, 3, 4}),
			getfunc: func(o *opgine.Opgine) (any, error) {
				return o.GetUintSlice(keyname)
			},
			mustfunc: func(o *opgine.Opgine) any {
				return o.MustGetUintSlice(keyname)
			},
		},
		"int8slice": {
			val: []int8{-1, -2, 1, 2},
			opt: opgine.WithInt8Slice(keyname, []int8{-1, -2, 1, 2}),
			getfunc: func(o *opgine.Opgine) (any, error) {
				return o.GetInt8Slice(keyname)
			},
			mustfunc: func(o *opgine.Opgine) any {
				return o.MustGetInt8Slice(keyname)
			},
		},
		"uint8slice": {
			val: []uint8{1, 2, 3, 4},
			opt: opgine.WithUint8Slice(keyname, []uint8{1, 2, 3, 4}),
			getfunc: func(o *opgine.Opgine) (any, error) {
				return o.GetUint8Slice(keyname)
			},
			mustfunc: func(o *opgine.Opgine) any {
				return o.MustGetUint8Slice(keyname)
			},
		},
		"int16slice": {
			val: []int16{-1, -2, 1, 2},
			opt: opgine.WithInt16Slice(keyname, []int16{-1, -2, 1, 2}),
			getfunc: func(o *opgine.Opgine) (any, error) {
				return o.GetInt16Slice(keyname)
			},
			mustfunc: func(o *opgine.Opgine) any {
				return o.MustGetInt16Slice(keyname)
			},
		},
		"uint16slice": {
			val: []uint16{1, 2, 3, 4},
			opt: opgine.WithUint16Slice(keyname, []uint16{1, 2, 3, 4}),
			getfunc: func(o *opgine.Opgine) (any, error) {
				return o.GetUint16Slice(keyname)
			},
			mustfunc: func(o *opgine.Opgine) any {
				return o.MustGetUint16Slice(keyname)
			},
		},
		"int32slice": {
			val: []int32{-1, -2, 1, 2},
			opt: opgine.WithInt32Slice(keyname, []int32{-1, -2, 1, 2}),
			getfunc: func(o *opgine.Opgine) (any, error) {
				return o.GetInt32Slice(keyname)
			},
			mustfunc: func(o *opgine.Opgine) any {
				return o.MustGetInt32Slice(keyname)
			},
		},
		"uint32slice": {
			val: []uint32{1, 2, 3, 4},
			opt: opgine.WithUint32Slice(keyname, []uint32{1, 2, 3, 4}),
			getfunc: func(o *opgine.Opgine) (any, error) {
				return o.GetUint32Slice(keyname)
			},
			mustfunc: func(o *opgine.Opgine) any {
				return o.MustGetUint32Slice(keyname)
			},
		},
		"int64slice": {
			val: []int64{-1, -2, 1, 2},
			opt: opgine.WithInt64Slice(keyname, []int64{-1, -2, 1, 2}),
			getfunc: func(o *opgine.Opgine) (any, error) {
				return o.GetInt64Slice(keyname)
			},
			mustfunc: func(o *opgine.Opgine) any {
				return o.MustGetInt64Slice(keyname)
			},
		},
		"uint64slice": {
			val: []uint64{1, 2, 3, 4},
			opt: opgine.WithUint64Slice(keyname, []uint64{1, 2, 3, 4}),
			getfunc: func(o *opgine.Opgine) (any, error) {
				return o.GetUint64Slice(keyname)
			},
			mustfunc: func(o *opgine.Opgine) any {
				return o.MustGetUint64Slice(keyname)
			},
		},
		"float32slice": {
			val: []float32{0.1, -0.2, 0.01, 0.02},
			opt: opgine.WithFloat32Slice(keyname, []float32{0.1, -0.2, 0.01, 0.02}),
			getfunc: func(o *opgine.Opgine) (any, error) {
				return o.GetFloat32Slice(keyname)
			},
			mustfunc: func(o *opgine.Opgine) any {
				return o.MustGetFloat32Slice(keyname)
			},
		},
		"float64slice": {
			val: []float64{0.1000000001, -0.2000000000002, 0.00000000000001, 0.000000000000000002},
			opt: opgine.WithFloat64Slice(keyname, []float64{0.1000000001, -0.2000000000002, 0.00000000000001, 0.000000000000000002}),
			getfunc: func(o *opgine.Opgine) (any, error) {
				return o.GetFloat64Slice(keyname)
			},
			mustfunc: func(o *opgine.Opgine) any {
				return o.MustGetFloat64Slice(keyname)
			},
		},
	}
	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			o, oerr := opgine.New(tc.opt)
			require.NoError(t, oerr)
			require.NotNil(t, o)
			res, reserr := tc.getfunc(o)
			require.NoError(t, reserr)
			require.Equal(t, tc.val, res)
			require.NotPanics(t, func() { tc.mustfunc(o) })
			require.Panics(t, func() { o.MustGetString(keyname) })
		})
	}
}
