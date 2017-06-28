package main

import (
	"fmt"

	"github.com/cirocosta/nfsvol/manager"
	"github.com/pkg/errors"

	v "github.com/docker/go-plugins-helpers/volume"
	log "github.com/sirupsen/logrus"
)

const (
	HostMountPoint = "/mnt"
)

type nfsVolDriver struct {
	logger  *log.Entry
	manager *manager.Manager
}

func newNfsVolDriver() (d nfsVolDriver, err error) {
	m, err := manager.New(manager.Config{
		Root: HostMountPoint,
	})
	if err != nil {
		err = errors.Wrapf(err,
			"Couldn't initiate fs manager mounting at %s",
			HostMountPoint)
		return
	}

	d.logger = log.WithField("from", "driver")
	d.manager = &m
	return
}

func (d nfsVolDriver) Create(req v.Request) (resp v.Response) {
	logger := d.logger.
		WithField("name", req.Name).
		WithField("opts", req.Options)

	logger.Debug("received request to create volume")

	abs, err := d.manager.Create(req.Name)
	if err != nil {
		logger.WithError(err).Error("couldn't create volume")
		resp.Err = err.Error()
		return
	}

	logger.WithField("abs", abs).Debug("volume created")
	return
}

func (d nfsVolDriver) List(req v.Request) (resp v.Response) {
	logger := d.logger.
		WithField("name", req.Name).
		WithField("opts", req.Options)

	logger.Debug("received request to list volumes")

	dirs, err := d.manager.List()
	if err != nil {
		logger.WithError(err).Error("couldn't list volumes")
		resp.Err = err.Error()
		return
	}

	resp.Volumes = make([]*v.Volume, len(dirs))
	for idx, dir := range dirs {
		resp.Volumes[idx] = &v.Volume{
			Name: dir,
		}
	}

	logger.
		WithField("number-of-volumes", len(dirs)).
		Debug("volumes listed")

	return
}

func (d nfsVolDriver) Get(req v.Request) (resp v.Response) {
	logger := d.logger.
		WithField("name", req.Name).
		WithField("opts", req.Options)

	logger.Debug("received request to get volume")

	mp, found, err := d.manager.Get(req.Name)
	if err != nil {
		logger.WithError(err).Error("errored retrieving path for volume")
		resp.Err = err.Error()
		return
	}

	if !found {
		logger.WithError(err).Info("volume not found")
		resp.Err = fmt.Sprintf("volume %s not found", req.Name)
		return
	}

	resp.Volume = &v.Volume{
		Name:       req.Name,
		Mountpoint: mp,
	}

	logger.WithField("mountpoint", mp).Debug("volume found")

	return
}

func (d nfsVolDriver) Remove(req v.Request) (resp v.Response) {
	logger := d.logger.
		WithField("name", req.Name).
		WithField("opts", req.Options)

	logger.Debug("received request to remove volume")

	err := d.manager.Delete(req.Name)
	if err != nil {
		logger.WithError(err).Error("errored trying to delete volume")
		resp.Err = err.Error()
		return
	}

	logger.Debug("volume deleted")

	return v.Response{}
}

func (d nfsVolDriver) Path(req v.Request) v.Response {
	d.logger.
		WithField("name", req.Name).
		Debug("received request to get path of volume")
	return v.Response{}
}

func (d nfsVolDriver) Mount(req v.MountRequest) v.Response {
	d.logger.
		WithField("name", req.Name).
		WithField("id", req.ID).
		Debug("received request to mount volume")
	return v.Response{}
}

func (d nfsVolDriver) Unmount(req v.UnmountRequest) v.Response {
	d.logger.
		WithField("name", req.Name).
		WithField("id", req.ID).
		Debug("received request to unmount volume")
	return v.Response{}
}

func (d nfsVolDriver) Capabilities(v.Request) v.Response {
	return v.Response{
		Capabilities: v.Capability{
			Scope: "global",
		},
	}
}
