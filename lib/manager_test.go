package lib_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	manager "github.com/cirocosta/nfsvol/lib"
)

func TestNew_failsWithoutRootSpecified(t *testing.T) {
	_, err := manager.New(manager.Config{})
	assert.Error(t, err)
}

func TestNew_failsWithInexistentRoot(t *testing.T) {
	_, err := manager.New(manager.Config{
		Root: "/a/b/c/d/e/f/g/h/i",
	})
	assert.Error(t, err)
}

func TestNew_failsWithNonAbsolutePath(t *testing.T) {
	_, err := manager.New(manager.Config{
		Root: "var/log",
	})
	assert.Error(t, err)
}

func TestNew_succeedsWithWriteableAbsolutePath(t *testing.T) {
	dir, err := ioutil.TempDir("", "")
	assert.NoError(t, err)
	defer os.RemoveAll(dir)

	_, err = manager.New(manager.Config{
		Root: dir,
	})
	assert.NoError(t, err)
}

func TestCreate_failsIfEmptyPath(t *testing.T) {
	dir, err := ioutil.TempDir("", "")
	assert.NoError(t, err)
	defer os.RemoveAll(dir)

	m, err := manager.New(manager.Config{
		Root: dir,
	})
	assert.NoError(t, err)

	_, err = m.Create("")
	assert.Error(t, err)
}

func TestCreate_failsIfNotAbsolutePath(t *testing.T) {
	dir, err := ioutil.TempDir("", "")
	assert.NoError(t, err)
	defer os.RemoveAll(dir)

	m, err := manager.New(manager.Config{
		Root: dir,
	})
	assert.NoError(t, err)

	_, err = m.Create("abc")
	assert.Error(t, err)
}

func TestCreate_succeedsWithAbsolutePath(t *testing.T) {
	dir, err := ioutil.TempDir("", "")
	assert.NoError(t, err)
	defer os.RemoveAll(dir)

	m, err := manager.New(manager.Config{
		Root: dir,
	})
	assert.NoError(t, err)

	_, err = m.Create("/abc")
	assert.NoError(t, err)
}
