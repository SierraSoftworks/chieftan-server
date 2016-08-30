package tools

import "github.com/urfave/cli"

var User cli.Command = cli.Command{
	Name:        "user",
	Description: "Commands used to manage user accounts",
	Flags:       []cli.Flag{},
	Subcommands: []cli.Command{
		userCreate,
		userRemove,
	},
}
