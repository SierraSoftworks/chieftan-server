package tools

import (
	"fmt"

	"github.com/SierraSoftworks/chieftan-server/models"
	"github.com/SierraSoftworks/chieftan-server/tasks"

	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

var permissionRemove cli.Command = cli.Command{
	Name:        "remove",
	Description: "Remove permissions from a specific user",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "admin",
			Usage: "Remove administrative privileges",
		},
	},
	Usage:     "EMAIL PERMISSIONS...",
	UsageText: "Update the permissions of a user",
	Action: func(c *cli.Context) error {
		if c.NArg() < 1 {
			cli.ShowCommandHelp(c, "remove")

			return fmt.Errorf("expected you to provide a user ID or email address")
		}

		req := &tasks.RemovePermissionsRequest{
			UserID:      c.Args().Get(0),
			Permissions: []string{},
		}

		if !models.IsValidUserID(req.UserID) {
			req.UserID = models.DeriveID(req.UserID)
		}

		if c.IsSet("admin") {
			req.Permissions = append(req.Permissions, "admin", "admin/users", "project/:project", "project/:project/admin")
		}

		req.Permissions = append(req.Permissions, c.Args()[1:]...)

		user, _, err := tasks.RemovePermissions(req)

		log.WithFields(log.Fields{
			"userID":      req.UserID,
			"name":        user.Name,
			"email":       user.Email,
			"permissions": user.Permissions,
		}).Infof("Updated permissions for '%s' (%s)", user.Name, user.Email)

		return err
	},
}
