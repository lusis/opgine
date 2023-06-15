package opgine_test

import (
	"testing"

	"github.com/lusis/opgine"
	"github.com/stretchr/testify/require"
)

func TestInts(t *testing.T) {
	keyname := t.Name()
	testCases := map[string]testCase{
		"int": {
			val: -1,
			opt: opgine.WithInt(keyname, -1),
			getfunc: func(o *opgine.Opgine) (any, error) {
				return o.GetInt(keyname)
			},
			mustfunc: func(o *opgine.Opgine) any {
				return o.MustGetInt(keyname)
			},
		},
		"int8": {
			val: int8(-1),
			opt: opgine.WithInt8(keyname, -1),
			getfunc: func(o *opgine.Opgine) (any, error) {
				return o.GetInt8(keyname)
			},
			mustfunc: func(o *opgine.Opgine) any {
				return o.MustGetInt8(keyname)
			},
		},
		"int16": {
			val: int16(-1),
			opt: opgine.WithInt16(keyname, -1),
			getfunc: func(o *opgine.Opgine) (any, error) {
				return o.GetInt16(keyname)
			},
			mustfunc: func(o *opgine.Opgine) any {
				return o.MustGetInt16(keyname)
			},
		},
		"int32": {
			val: int32(-1),
			opt: opgine.WithInt32(keyname, -1),
			getfunc: func(o *opgine.Opgine) (any, error) {
				return o.GetInt32(keyname)
			},
			mustfunc: func(o *opgine.Opgine) any {
				return o.MustGetInt32(keyname)
			},
		},
		"int64": {
			val: int64(-1),
			opt: opgine.WithInt64(keyname, -1),
			getfunc: func(o *opgine.Opgine) (any, error) {
				return o.GetInt64(keyname)
			},
			mustfunc: func(o *opgine.Opgine) any {
				return o.MustGetInt64(keyname)
			},
		},
		"uint": {
			val: uint(1),
			opt: opgine.WithUint(keyname, 1),
			getfunc: func(o *opgine.Opgine) (any, error) {
				return o.GetUint(keyname)
			},
			mustfunc: func(o *opgine.Opgine) any {
				return o.MustGetUint(keyname)
			},
		},
		"uint8": {
			val: uint8(1),
			opt: opgine.WithUint8(keyname, 1),
			getfunc: func(o *opgine.Opgine) (any, error) {
				return o.GetUint8(keyname)
			},
			mustfunc: func(o *opgine.Opgine) any {
				return o.MustGetUint8(keyname)
			},
		},
		"uint16": {
			val: uint16(1),
			opt: opgine.WithUint16(keyname, 1),
			getfunc: func(o *opgine.Opgine) (any, error) {
				return o.GetUint16(keyname)
			},
			mustfunc: func(o *opgine.Opgine) any {
				return o.MustGetUint16(keyname)
			},
		},
		"uint32": {
			val: uint32(1),
			opt: opgine.WithUint32(keyname, 1),
			getfunc: func(o *opgine.Opgine) (any, error) {
				return o.GetUint32(keyname)
			},
			mustfunc: func(o *opgine.Opgine) any {
				return o.MustGetUint32(keyname)
			},
		},
		"uint64": {
			val: uint64(1),
			opt: opgine.WithUint64(keyname, 1),
			getfunc: func(o *opgine.Opgine) (any, error) {
				return o.GetUint64(keyname)
			},
			mustfunc: func(o *opgine.Opgine) any {
				return o.MustGetUint64(keyname)
			},
		},
		"float32": {
			val: float32(0.12345679),
			opt: opgine.WithFloat32(keyname, 0.12345679),
			getfunc: func(o *opgine.Opgine) (any, error) {
				return o.GetFloat32(keyname)
			},
			mustfunc: func(o *opgine.Opgine) any {
				return o.MustGetFloat32(keyname)
			},
		},
		"float64": {
			val: 0.123456791111111112,
			opt: opgine.WithFloat64(keyname, 0.1234567911111111111111112),
			getfunc: func(o *opgine.Opgine) (any, error) {
				return o.GetFloat64(keyname)
			},
			mustfunc: func(o *opgine.Opgine) any {
				return o.MustGetFloat64(keyname)
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
