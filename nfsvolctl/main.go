package main

import (
	"os"

	"github.com/cirocosta/nfsvol/nfsvolctl/commands"
	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "nfsvolctl"
	app.Usage = "Controls the 'nfsvol' volume plugin"
	app.Commands = []cli.Command{
		commands.Ls,
	}
	app.Run(os.Args)
}
