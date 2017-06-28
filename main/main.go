package main

import (
	"os"

	"github.com/pkg/errors"

	v "github.com/docker/go-plugins-helpers/volume"
	log "github.com/sirupsen/logrus"
)

const (
	socketAddress = "/run/docker/plugins/nfsvol.sock"
)

func main() {
	if os.Getenv("DEBUG") != "" {
		log.SetLevel(log.DebugLevel)
	}

	d, err := newNfsVolDriver()
	if err != nil {
		err = errors.Wrapf(err,
			"Failed to initialize NFS volume driver")
		log.Fatal(err)
	}

	h := v.NewHandler(d)
	log.Infof("Listening on %s", socketAddress)
	log.Error(h.ServeUnix(socketAddress, 0))
}
