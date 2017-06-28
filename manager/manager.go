package manager

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"golang.org/x/sys/unix"
)

type Manager struct {
	root string
}

type Config struct {
	Root string
}

func New(cfg Config) (manager Manager, err error) {
	if cfg.Root == "" {
		err = errors.Errorf("Root not specified.")
		return
	}

	if !filepath.IsAbs(cfg.Root) {
		err = errors.Errorf(
			"Root (%s) must be an absolute path",
			cfg.Root)
		return
	}

	err = unix.Access(cfg.Root, unix.W_OK)
	if err != nil {
		err = errors.Wrapf(err,
			"Root (%s) must be writable.",
			cfg.Root)
		return
	}

	manager.root = cfg.Root
	return
}

func (m Manager) List() (directories []string, err error) {
	files, err := ioutil.ReadDir(m.root)
	if err != nil {
		err = errors.Wrapf(err,
			"Couldn't list files/directories from %s", m.root)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			directories = append(directories, file.Name())
		}
	}

	return
}

func (m Manager) Create(path string) (absPath string, err error) {
	if path == "" {
		err = errors.Errorf(
			"Can't create with empty path")
		return
	}

	if !filepath.IsAbs(path) {
		err = errors.Errorf(
			"path (%s) must be absolute",
			path)
		return
	}

	absPath = filepath.Join(m.root, path)
	err = os.MkdirAll(absPath, 0755)
	if err != nil {
		err = errors.Wrapf(err,
			"Couldn't create directory %s", absPath)
		return
	}

	return
}
