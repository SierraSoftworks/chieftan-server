package tools

import (
	"fmt"

	"github.com/SierraSoftworks/chieftan-server/tasks"

	"github.com/urfave/cli"
)

var RemoveToken cli.Command = cli.Command{
	Name:        "remove:token",
	Description: "Remove a specific access token",
	Category:    "administration",
	Flags:       []cli.Flag{},
	Usage:       "TOKEN",
	UsageText:   "Provide the access token to be removed",
	Action: func(c *cli.Context) error {
		if c.NArg() < 1 {
			cli.ShowCommandHelp(c, "remove:token")

			return fmt.Errorf("expected you to provide the token to be removed")
		}

		req := &tasks.RemoveTokenRequest{
			Token: c.Args().Get(0),
		}

		_, err := tasks.RemoveToken(req)
		if err != nil {
			return err
		}

		return nil
	},
}
