package manager_test

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/cirocosta/nfsvol/manager"
	"github.com/stretchr/testify/assert"
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

	absPath, err := m.Create("/abc")
	assert.NoError(t, err)
	assert.Equal(t, path.Join(dir, "abc"), absPath)

	finfo, err := os.Stat(absPath)
	assert.NoError(t, err)
	assert.True(t, finfo.IsDir())
}

func TestList_canList0Directorise(t *testing.T) {
	dir, err := ioutil.TempDir("", "")
	assert.NoError(t, err)
	defer os.RemoveAll(dir)

	m, err := manager.New(manager.Config{
		Root: dir,
	})
	assert.NoError(t, err)

	dirs, err := m.List()
	assert.NoError(t, err)
	assert.Len(t, dirs, 0)
}

func TestList_listsDirectories(t *testing.T) {
	dir, err := ioutil.TempDir("", "")
	assert.NoError(t, err)
	defer os.RemoveAll(dir)

	m, err := manager.New(manager.Config{
		Root: dir,
	})
	assert.NoError(t, err)

	_, err = m.Create("/abc")
	assert.NoError(t, err)

	_, err = m.Create("/def")
	assert.NoError(t, err)

	dirs, err := m.List()
	assert.NoError(t, err)
	assert.Len(t, dirs, 2)
	assert.Equal(t, "abc", dirs[0])
	assert.Equal(t, "def", dirs[1])
}

func TestGet_doesntErrorIfNotFound(t *testing.T) {
	dir, err := ioutil.TempDir("", "")
	assert.NoError(t, err)
	defer os.RemoveAll(dir)

	m, err := manager.New(manager.Config{
		Root: dir,
	})
	assert.NoError(t, err)

	_, found, err := m.Get("abc")
	assert.NoError(t, err)
	assert.False(t, found)
}
