package main

import (
	v "github.com/docker/go-plugins-helpers/volume"
)

type nfsVolDriver struct{}

func newNfsVolDriver () (d nfsVolDriver, err error) {
  return
}

func (d nfsVolDriver) Create(v.Request) v.Response {
  return v.Response{}
}

func (d nfsVolDriver) List(v.Request) v.Response {
  return v.Response{}
}

func (d nfsVolDriver) Get(v.Request) v.Response {
  return v.Response{}
}

func (d nfsVolDriver) Remove(v.Request) v.Response {
  return v.Response{}
}

func (d nfsVolDriver) Path(v.Request) v.Response {
  return v.Response{}
}

func (d nfsVolDriver) Mount(v.MountRequest) v.Response {
  return v.Response{}
}

func (d nfsVolDriver) Unmount(v.UnmountRequest) v.Response {
  return v.Response{}
}

func (d nfsVolDriver) Capabilities(v.Request) v.Response {
  return v.Response{}
}
