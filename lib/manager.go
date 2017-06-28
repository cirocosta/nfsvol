package lib

import (
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

func (m Manager) Create(name string) (err error) {
	return
}
