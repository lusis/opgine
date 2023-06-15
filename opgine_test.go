package opgine_test

import (
	"testing"

	"github.com/lusis/opgine"
	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	keyname := t.Name()
	res, err := opgine.New(opgine.WithInt(keyname, 1234))
	require.NoError(t, err)
	require.NotNil(t, res)
	require.True(t, res.Check(keyname))
	t.Run("get", func(t *testing.T) {
		v, verr := res.Get(keyname)
		require.NoError(t, verr)
		require.Equal(t, 1234, v.(int))
	})
	t.Run("mustget", func(t *testing.T) {
		require.NotPanics(t, func() { res.MustGet(keyname) })
		require.Panics(t, func() { res.MustGet("invalidkey") })
	})

	t.Run("mismatch", func(t *testing.T) {
		b, berr := res.GetBool(keyname)
		require.ErrorIs(t, berr, opgine.ErrMismatchedTypes)
		require.False(t, b)
	})
}

func TestWithInterface(t *testing.T) {
	o, oerr := opgine.New(opgine.WithInterface("mykey", "abcdefg"))
	require.NoError(t, oerr)
	require.NotNil(t, o)
	gres, gerr := o.Get("mykey")
	require.NoError(t, gerr)
	require.NotEmpty(t, gres)
	require.Equal(t, "abcdefg", gres.(string))
}

type testCase struct {
	opt      opgine.Option
	val      any
	getfunc  func(*opgine.Opgine) (any, error)
	mustfunc func(*opgine.Opgine) any
}
