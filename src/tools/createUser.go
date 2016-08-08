package tools

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/SierraSoftworks/chieftan-server/src/tasks"

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

		var stdErr io.Writer

		if c.App.ErrWriter != nil {
			stdErr = c.App.ErrWriter
		} else {
			stdErr = os.Stderr
		}

		infoLogger := log.New(stdErr, "[INFO] ", 0)

		infoLogger.Printf("Creating user '%s' with email '%s'", req.Name, req.Email)

		if c.IsSet("admin") {
			req.Permissions = []string{
				"admin",
				"admin/users",
				"project/:project",
				"project/:project/admin",
			}
		}

		user, err := tasks.CreateUser(req)

		if err != nil {
			return err
		}

		infoLogger.Printf("User created with ID:")
		fmt.Println(user.ID)
		infoLogger.Printf("Run 'chieftan create:token %s' to generate a new access token for this user", user.ID)

		return nil
	},
}
