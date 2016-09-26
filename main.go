package main

import (
	"os"

	"github.com/SierraSoftworks/chieftan-server/tools"

	"strings"

	"github.com/SierraSoftworks/chieftan-server/models"
	log "github.com/Sirupsen/logrus"
	"github.com/getsentry/raven-go"
	"github.com/urfave/cli"
)

func main() {

	if envDSN := os.Getenv("SENTRY_DSN"); envDSN != "" {
		raven.SetDSN(envDSN)
	} else if sentry_dsn != "" {
		raven.SetDSN(sentry_dsn)
	}

	raven.SetRelease(version)

	app := cli.NewApp()
	app.Name = "Chieftan"
	app.Usage = "Manage your Chieftan instance"

	app.Author = "Benjamin Pannell"
	app.Email = "admin@sierrasoftworks.com"
	app.Copyright = "Sierra Softworks © 2016"
	app.Version = version

	app.Commands = cli.Commands{
		RunServer,
		tools.User,
		tools.Token,
		tools.Permission,
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "log-level",
			Usage: "DEBUG|INFO|WARN|ERROR",
			Value: "INFO",
		},
		cli.StringFlag{
			Name:   "mongodb",
			EnvVar: "MONGODB_URL",
			Usage:  "MongoDB server URL",
			Value:  "mongodb://localhost:27017/chieftan",
		},
	}

	app.Before = func(c *cli.Context) error {
		log.WithFields(log.Fields{
			"log-level": c.GlobalString("log-leve"),
			"mongodb":   c.GlobalString("mongodb"),
		}).Info("Starting")

		logLevel := c.GlobalString("log-level")
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

		err := models.Connect(c.GlobalString("mongodb"))
		if err != nil {
			return err
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		os.Exit(1)
	}

	os.Exit(0)
}
