package opgine_test

import (
	"testing"

	"github.com/lusis/opgine"
	"github.com/stretchr/testify/require"
)

func TestString(t *testing.T) {
	keyname := t.Name()
	val := keyname
	res, err := opgine.New(opgine.WithString(keyname, val))
	require.NoError(t, err)
	require.NotNil(t, res)
	require.True(t, res.Check(keyname))

	t.Run("getstring", func(t *testing.T) {
		v, verr := res.GetString(keyname)
		require.NoError(t, verr)
		require.Equal(t, val, v)
	})
	t.Run("mustgetstring", func(t *testing.T) {
		require.Panics(t, func() { res.MustGetString("invalid") })
		require.Panics(t, func() { res.MustGetBool(keyname) })
		require.NotPanics(t, func() { res.MustGetString(keyname) })
	})

	t.Run("mismatch", func(t *testing.T) {
		b, berr := res.GetBool(keyname)
		require.ErrorIs(t, berr, opgine.ErrMismatchedTypes)
		require.False(t, b)
	})
}
