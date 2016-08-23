package main

import (
	"os"

	"github.com/SierraSoftworks/chieftan-server/tools"

	"strings"

	"github.com/SierraSoftworks/chieftan-server/models"
	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Chieftan"
	app.Usage = "Manage your Chieftan instance"

	app.Author = "Benjamin Pannell"
	app.Email = "admin@sierrasoftworks.com"
	app.Copyright = "Sierra Softworks © 2016"
	app.Version = version

	app.Commands = cli.Commands{
		RunServer,
		tools.CreateUser,
		tools.CreateToken,
		tools.RemoveToken,
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "log-level",
			Usage: "DEBUG|INFO|WARN|ERROR",
		},
		cli.StringFlag{
			Name:   "mongodb",
			EnvVar: "MONGODB_URL",
			Usage:  "mongodb://localhost:27017/chieftan",
			Value:  "mongodb://localhost:27017/chieftan",
		},
	}

	app.Before = func(c *cli.Context) error {
		if c.IsSet("log-level") {
			logLevel := c.String("log-level")
			switch strings.ToUpper(logLevel) {
			case "DEBUG":
				log.SetLevel(log.DebugLevel)
			case "INFO":
				log.SetLevel(log.InfoLevel)
			case "WARN":
				log.SetLevel(log.WarnLevel)
			case "ERROR":
				log.SetLevel(log.ErrorLevel)
			default:
				log.SetLevel(log.InfoLevel)
			}
		}

		models.Connect(c.String("mongodb"))

		return nil
	}

	app.Run(os.Args)
}
