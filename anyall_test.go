package opgine_test

import (
	"testing"

	"github.com/lusis/opgine"
	"github.com/stretchr/testify/require"
)

func TestAnyAll(t *testing.T) {
	t.Run("any", func(t *testing.T) {
		o, oerr := opgine.New(opgine.WithAll(), opgine.WithBool("mybool", true), opgine.WithInt("myint", 1))
		require.NoError(t, oerr)
		require.NotNil(t, o)

		require.True(t, o.RequiresAll())
		require.False(t, o.RequiresAny())
		require.Equal(t, 1, o.MustGetInt("myint"))
		require.Equal(t, true, o.MustGetBool("mybool"))
	})
	t.Run("all", func(t *testing.T) {
		o, oerr := opgine.New(opgine.WithAny(), opgine.WithBool("mybool", true), opgine.WithInt("myint", 1))
		require.NoError(t, oerr)
		require.NotNil(t, o)

		require.True(t, o.RequiresAny())
		require.False(t, o.RequiresAll())
		require.Equal(t, 1, o.MustGetInt("myint"))
		require.Equal(t, true, o.MustGetBool("mybool"))
	})
}
