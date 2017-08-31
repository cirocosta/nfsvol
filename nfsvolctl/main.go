package main

import (
	"os"

	"github.com/cirocosta/nfsvol/nfsvolctl/commands"
	"gopkg.in/urfave/cli.v1"
)

var version string = "master-dev"

func main() {
	app := cli.NewApp()
	app.Name = "nfsvolctl"
	app.Version = version
	app.Usage = "Controls the 'nfsvol' volume plugin"
	app.Commands = []cli.Command{
		commands.Ls,
	}
	app.Run(os.Args)
}
