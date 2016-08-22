package main

import (
	"fmt"
	"net/http"

	"github.com/SierraSoftworks/chieftan-server/src/api"
	log "github.com/Sirupsen/logrus"
	"github.com/rs/cors"
	"github.com/urfave/cli"
)

var RunServer = cli.Command{
	Name:        "server",
	Description: "Run the Chieftan API server",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:   "ip",
			EnvVar: "IP",
			Value:  "",
		},
		cli.Int64Flag{
			Name:   "port",
			EnvVar: "PORT",
			Value:  3000,
		},
	},
	Action: func(c *cli.Context) error {
		log.Info("Starting API server")

		log.Debug("Registering API routes")
		mux := http.NewServeMux()
		mux.Handle("/api/", http.StripPrefix("/api", api.Router()))
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(404)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"code":404, "error": "Not Found", "message": "The method you attempted to make use of could not be found on our system."}`))
		})

		listenOn := fmt.Sprintf("%s:%d", c.String("ip"), c.Int64("port"))
		log.WithFields(log.Fields{
			"server": c.String("ip"),
			"port":   c.Int64("port"),
		}).Infof("Listening on %s", listenOn)
		return http.ListenAndServe(listenOn, cors.New(cors.Options{
			AllowCredentials: true,
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		}).Handler(mux))
	},
}
