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
