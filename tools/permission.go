package tools

import "github.com/urfave/cli"

var Permission cli.Command = cli.Command{
	Name:        "permissions",
	Description: "Commands used to manage user permissions",
	Flags:       []cli.Flag{},
	Subcommands: []cli.Command{
		permissionSet,
		permissionAdd,
		permissionRemove,
	},
}
