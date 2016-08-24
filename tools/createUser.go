package tools

import (
	"fmt"

	"github.com/SierraSoftworks/chieftan-server/tasks"

	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

var CreateUser cli.Command = cli.Command{
	Name:        "create:user",
	Description: "Create a new user and return their unique ID",
	Category:    "administration",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:   "admin",
			EnvVar: "CHIEFTAN_CREATE_USER_ADMIN",
			Usage:  "Create a user with administrative privileges",
		},
	},
	Usage:     "NAME EMAIL",
	UsageText: "Create a new user and return their unique ID",
	Action: func(c *cli.Context) error {
		if c.NArg() < 2 {
			cli.ShowCommandHelp(c, "create:user")

			return fmt.Errorf("expected you to provide both a name and email for the user")
		}

		req := &tasks.CreateUserRequest{
			Name:        c.Args().Get(0),
			Email:       c.Args().Get(1),
			Permissions: []string{},
		}

		log.WithFields(log.Fields{
			"Name":  req.Name,
			"Email": req.Email,
		}).Infof("Creating user '%s' with email '%s'", req.Name, req.Email)

		if c.IsSet("admin") {
			req.Permissions = []string{
				"admin",
				"admin/users",
				"project/:project",
				"project/:project/admin",
			}
		}

		user, _, err := tasks.CreateUser(req)

		if err != nil {
			return err
		}

		log.WithFields(log.Fields{
			"UserID": user.ID,
		}).Infof("User created with ID: %s", user.ID)

		log.WithFields(log.Fields{
			"UserID": user.ID,
		}).Infof("Run 'chieftan create:token %s' to generate a new access token for this user", user.ID)

		return nil
	},
}
