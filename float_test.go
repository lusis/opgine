package opgine_test

import (
	"testing"

	"github.com/lusis/opgine"
	"github.com/stretchr/testify/require"
)

func TestFloats(t *testing.T) {
	keyname := t.Name()
	testCases := map[string]testCase{
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
