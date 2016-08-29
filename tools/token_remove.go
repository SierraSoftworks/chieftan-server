package tools

import (
	"fmt"

	"github.com/SierraSoftworks/chieftan-server/models"
	"github.com/SierraSoftworks/chieftan-server/tasks"

	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

var tokenRemove cli.Command = cli.Command{
	Name:        "remove",
	Description: "Remove a specific access token",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "user",
			Usage: "The email or ID of a user to remove tokens for",
		},
		cli.BoolFlag{
			Name:  "global",
			Usage: "Remove all tokens on this server",
		},
	},
	Usage:     "TOKEN",
	UsageText: "Provide the access token to be removed",
	Action: func(c *cli.Context) error {
		hasToken := c.NArg() > 0
		hasGlobal := c.IsSet("global")
		hasUser := c.IsSet("user")

		if hasToken && hasGlobal || hasToken && hasUser || hasUser && hasGlobal {
			cli.ShowCommandHelp(c, "remove")

			return fmt.Errorf("expected you to provide only the token to be removed, a user, or the global flag")
		}

		if hasToken {
			req := &tasks.RemoveTokenRequest{
				Token: c.Args().Get(0),
			}
			log.WithField("token", req.Token).Infof("Removing token '%s'", req.Token)
			_, err := tasks.RemoveToken(req)
			if err != nil {
				return err
			}
		}

		if hasGlobal {
			log.Info("Removing all tokens on this server")
			_, err := tasks.RemoveAllTokens(&tasks.RemoveAllTokensRequest{})
			if err != nil {
				return err
			}
		}

		if hasUser {
			req := &tasks.RemoveAllTokensRequest{
				UserID: c.String("user"),
			}

			if !models.IsValidUserID(req.UserID) {
				req.UserID = models.DeriveID(req.UserID)
			}

			log.WithField("userID", req.UserID).Infof("Removing all tokens belonging to '%s'", req.UserID)
			_, err := tasks.RemoveAllTokens(req)
			if err != nil {
				return err
			}
		}

		return nil
	},
}
