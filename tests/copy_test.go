package tests

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	ts := newTester(t)
	defer ts.teardown()

	t.Run("copy w/ empty store", func(t *testing.T) {
		_, err := ts.run("copy")
		require.Error(t, err)
	})

	ts.initStore()

	t.Run("copy usage", func(t *testing.T) {
		out, err := ts.run("copy")
		require.Error(t, err)
		assert.Equal(t, "\nError: Usage: "+filepath.Base(ts.Binary)+" cp <FROM> <TO>\n", out)
	})

	t.Run("copy w/o destination", func(t *testing.T) {
		out, err := ts.run("copy foo")
		require.Error(t, err)
		assert.Equal(t, "\nError: Usage: "+filepath.Base(ts.Binary)+" cp <FROM> <TO>\n", out)
	})

	t.Run("copy non existing source", func(t *testing.T) {
		out, err := ts.run("copy foo bar")
		require.Error(t, err)
		assert.Equal(t, "\nError: foo does not exist\n", out)
	})

	ts.initSecrets("")

	t.Run("recursive copy", func(t *testing.T) {
		_, err := ts.run("copy foo/ bar")
		require.NoError(t, err)
	})

	t.Run("copy existing secret to non-existing destination", func(t *testing.T) {
		out, err := ts.run("copy foo/bar foo/baz")
		require.NoError(t, err)
		assert.Empty(t, out)

		orig, err := ts.run("show -f foo/bar")
		require.NoError(t, err)

		cp, err := ts.run("show -f foo/baz")
		require.NoError(t, err)

		assert.Equal(t, orig, cp)
	})
}
