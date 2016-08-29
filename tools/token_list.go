package tools

import (
	"fmt"

	"github.com/SierraSoftworks/chieftan-server/models"
	"github.com/SierraSoftworks/chieftan-server/tasks"

	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

var tokenList cli.Command = cli.Command{
	Name:        "list",
	Description: "Get a list of a user's access tokens",
	Flags:       []cli.Flag{},
	Usage:       "EMAIL",
	UsageText:   "Provide the email address or unique ID of the user to get tokens for",
	Action: func(c *cli.Context) error {
		if c.NArg() < 1 {
			cli.ShowCommandHelp(c, "list")

			return fmt.Errorf("expected you to provide the email address or unique ID of the user")
		}

		req := &tasks.GetUserTokensRequest{
			ID: c.Args().Get(0),
		}

		if !models.IsValidUserID(req.ID) {
			req.ID = models.DeriveID(req.ID)
		}

		tokens, _, err := tasks.GetUserTokens(req)
		if err != nil {
			return err
		}

		log.WithField("userID", req.ID).WithField("tokens", tokens).Infof("User Tokens for '%s'", req.ID)

		for _, token := range tokens {
			log.Infoln(token)
		}

		return nil
	},
}
