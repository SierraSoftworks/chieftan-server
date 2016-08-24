package tools

import (
	"fmt"

	"github.com/SierraSoftworks/chieftan-server/models"
	"github.com/SierraSoftworks/chieftan-server/tasks"
	log "github.com/Sirupsen/logrus"

	"github.com/urfave/cli"
)

var CreateToken cli.Command = cli.Command{
	Name:        "create:token",
	Description: "Create a new access token for the given user",
	Category:    "administration",
	Flags:       []cli.Flag{},
	Usage:       "USER",
	UsageText:   "Provide the user's unique ID or email address",
	Action: func(c *cli.Context) error {
		if c.NArg() < 1 {
			cli.ShowCommandHelp(c, "create:token")

			return fmt.Errorf("expected you to provide either the user's ID or email address")
		}

		req := &tasks.CreateTokenRequest{
			UserID: c.Args().Get(0),
		}

		if !models.IsValidUserID(req.UserID) {
			req.UserID = models.DeriveID(req.UserID)
		}

		log.WithFields(log.Fields{
			"UserID": req.UserID,
		}).Infof("Creating new token for user with ID: %s", req.UserID)

		token, _, err := tasks.CreateToken(req)
		if err != nil {
			return err
		}

		log.WithFields(log.Fields{"token": token}).Infof("Access token created: %s", token)

		return nil
	},
}
