package opgine_test

import (
	"testing"

	"github.com/lusis/opgine"
	"github.com/stretchr/testify/require"
)

func TestBool(t *testing.T) {
	keyname := t.Name()
	testCases := map[string]testCase{
		"true": {
			val: true,
			opt: opgine.WithBool(keyname, true),
			getfunc: func(o *opgine.Opgine) (any, error) {
				return o.GetBool(keyname)
			},
			mustfunc: func(o *opgine.Opgine) any {
				return o.MustGetBool(keyname)
			},
		},
		"false": {
			val: false,
			opt: opgine.WithBool(keyname, false),
			getfunc: func(o *opgine.Opgine) (any, error) {
				return o.GetBool(keyname)
			},
			mustfunc: func(o *opgine.Opgine) any {
				return o.MustGetBool(keyname)
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
