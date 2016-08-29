package tools

import (
	"fmt"

	"github.com/SierraSoftworks/chieftan-server/models"
	"github.com/SierraSoftworks/chieftan-server/tasks"

	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

var userSetPermissions cli.Command = cli.Command{
	Name:        "set-permissions",
	Description: "Update the permissions belonging to a specific user",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "admin",
			Usage: "Grant administrative privileges",
		},
		cli.BoolFlag{
			Name:  "append",
			Usage: "Append the new permissions instead of replacing",
		},
	},
	Usage:     "EMAIL PERMISSIONS...",
	UsageText: "Update the permissions of a user",
	Action: func(c *cli.Context) error {
		if c.NArg() < 1 {
			cli.ShowCommandHelp(c, "set-permissions")

			return fmt.Errorf("expected you to provide a user ID or email address")
		}

		req := &tasks.SetPermissionsRequest{
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

		log.WithFields(log.Fields{
			"UserID":      req.UserID,
			"Permissions": req.Permissions,
		}).Infof("Updating permissions for '%s'", req.UserID)

		_, err := tasks.SetPermissions(req)

		return err
	},
}