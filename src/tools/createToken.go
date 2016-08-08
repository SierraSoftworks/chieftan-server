package tools

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/SierraSoftworks/chieftan-server/src/models"
	"github.com/SierraSoftworks/chieftan-server/src/tasks"

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

		var stdErr io.Writer

		if c.App.ErrWriter != nil {
			stdErr = c.App.ErrWriter
		} else {
			stdErr = os.Stderr
		}

		if !models.IsValidUserID(req.UserID) {
			req.UserID = models.DeriveID(req.UserID)
		}

		infoLogger := log.New(stdErr, "[INFO] ", 0)

		infoLogger.Printf("Creating new token for user with ID: %s", req.UserID)

		token, err := tasks.CreateToken(req)
		if err != nil {
			return err
		}

		infoLogger.Printf("Access token created:")
		fmt.Println(token)

		return nil
	},
}
