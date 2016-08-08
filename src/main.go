package main

import (
	"os"

	"github.com/SierraSoftworks/chieftan-server/src/tools"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Chieftan"
	app.Usage = "Manage your Chieftan instance"

	app.Author = "Benjamin Pannell"
	app.Email = "admin@sierrasoftworks.com"
	app.Copyright = "Sierra Softworks © 2016"
	app.Version = "1.0.0-dev"

	app.Commands = cli.Commands{
		tools.CreateUser,
		tools.CreateToken,
		tools.RemoveToken,
	}

	app.Run(os.Args)
}
