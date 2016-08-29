package tools

import "github.com/urfave/cli"

var Token cli.Command = cli.Command{
	Name:        "token",
	Description: "Commands used to manage user access tokens",
	Flags:       []cli.Flag{},
	Subcommands: []cli.Command{
		tokenCreate,
		tokenList,
		tokenRemove,
	},
}
