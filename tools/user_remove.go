package tools

import (
	"fmt"

	"github.com/SierraSoftworks/chieftan-server/models"
	"github.com/SierraSoftworks/chieftan-server/tasks"

	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

var userRemove cli.Command = cli.Command{
	Name:        "remove",
	Description: "Remove a specific user from the database",
	Flags:       []cli.Flag{},
	Usage:       "EMAIL",
	UsageText:   "Provide the ID or email address of the user to remove",
	Action: func(c *cli.Context) error {
		if c.NArg() < 1 {
			cli.ShowCommandHelp(c, "remove")

			return fmt.Errorf("expected you to provide a user ID or email address")
		}

		req := &tasks.RemoveUserRequest{
			UserID: c.Args().Get(0),
		}

		if !models.IsValidUserID(req.UserID) {
			req.UserID = models.DeriveID(req.UserID)
		}

		log.WithFields(log.Fields{
			"UserID": req.UserID,
		}).Infof("Removing user '%s'", req.UserID)

		user, _, err := tasks.RemoveUser(req)

		if err != nil {
			return err
		}

		log.WithFields(log.Fields{
			"userID": user.ID,
			"name":   user.Name,
			"email":  user.Email,
		}).Infof("Removed '%s' (%s)", user.Name, user.Email)

		return nil
	},
}
