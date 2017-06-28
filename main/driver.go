package main

import (
	v "github.com/docker/go-plugins-helpers/volume"
	log "github.com/sirupsen/logrus"
)

type nfsVolDriver struct {
	logger *log.Entry
}

func newNfsVolDriver() (d nfsVolDriver, err error) {
	d.logger = log.WithField("from", "driver")
	return
}

func (d nfsVolDriver) Create(req v.Request) v.Response {
	d.logger.
		WithField("name", req.Name).
		WithField("opts", req.Options).
		Debug("received request to create volume")
	return v.Response{}
}

func (d nfsVolDriver) List(req v.Request) v.Response {
	d.logger.Debug("received request to list volumes")
	return v.Response{}
}

func (d nfsVolDriver) Get(req v.Request) v.Response {
	d.logger.
		WithField("name", req.Name).
		Debug("received request to get volume")
	return v.Response{}
}

func (d nfsVolDriver) Remove(req v.Request) v.Response {
	d.logger.
		WithField("name", req.Name).
		Debug("received request to remove volume")
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
