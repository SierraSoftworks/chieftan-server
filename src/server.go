package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/urfave/cli"
)

var RunServer = cli.Command{
	Name:        "server",
	Description: "Run the Chieftan API server",
	Flags: []cli.Flag{
		cli.Int64Flag{
			Name:   "port",
			EnvVar: "PORT",
			Value:  3000,
		},
	},
	Action: func(c *cli.Context) error {
		l := log.New(c.App.Writer, "[INFO] ", 0)

		l.Println("Starting API server")

		mux := http.NewServeMux()
		mux.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello world!"))
		})
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(404)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"code":404, "error": "Not Found", "message": "The method you attempted to make use of could not be found on our system."}`))
		})

		return http.ListenAndServe(fmt.Sprintf("%s:%d", "", c.Int64("port")), mux)
	},
}
