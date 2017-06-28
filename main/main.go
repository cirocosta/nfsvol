package main

import (
	"io"
	"os"

	"github.com/pkg/errors"

	v "github.com/docker/go-plugins-helpers/volume"
	log "github.com/sirupsen/logrus"
)

const (
	socketAddress      = "/run/docker/plugins/nfsvol.sock"
	logFileDestination = "/var/log/nfsvol/plugin.log"
)

func main() {

	if os.Getenv("DEBUG") != "1" {
		log.SetLevel(log.DebugLevel)
	}

	f, err := os.Create(logFileDestination)
	if err != nil {
		err = errors.Wrapf(err,
			"Failed to create log file at %s", logFileDestination)
		log.Fatal(err)
	}
	defer f.Close()

	stdoutAndFile := io.MultiWriter(os.Stdout, f)

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(stdoutAndFile)

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
