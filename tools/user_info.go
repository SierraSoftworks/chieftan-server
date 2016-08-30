package tools

import (
	"fmt"

	"github.com/SierraSoftworks/chieftan-server/models"
	"github.com/SierraSoftworks/chieftan-server/tasks"

	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

var userInfo cli.Command = cli.Command{
	Name:        "info",
	Description: "Get details of a specific user",
	Flags:       []cli.Flag{},
	Usage:       "EMAIL",
	UsageText:   "Provide the ID or email address of the user",
	Action: func(c *cli.Context) error {
		if c.NArg() < 1 {
			cli.ShowCommandHelp(c, "info")

			return fmt.Errorf("expected you to provide a user ID or email address")
		}

		req := &tasks.GetUserRequest{
			ID: c.Args().Get(0),
		}

		if !models.IsValidUserID(req.ID) {
			req.ID = models.DeriveID(req.ID)
		}

		log.WithFields(log.Fields{
			"UserID": req.ID,
		}).Infof("Looking for user '%s'", req.ID)

		user, err := tasks.GetUser(req)

		if err != nil {
			return err
		}

		log.WithFields(log.Fields{
			"userID":      user.ID,
			"name":        user.Name,
			"email":       user.Email,
			"permissions": user.Permissions,
		}).Infof("%s (%s)", user.Name, user.Email)

		return nil
	},
}
