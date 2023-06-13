package cmd

import (
	"github.com/urfave/cli/v2"

	"github.com/batx-dev/batcmd/internal/http"
)

func NewServeCommand() *cli.Command {
	cmd := &ServeCommand{}
	return &cli.Command{
		Name:    "serve",
		Aliases: []string{"s"},
		Usage:   "Serve BatCmd as a api server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "listen",
				Aliases:     []string{"l"},
				EnvVars:     []string{"BATCMD_LISTEN"},
				Destination: &cmd.listen,
				Value:       ":8080",
				Usage:       "The address to http listen",
			},
			&cli.StringFlag{
				Name:        "backend-url",
				Aliases:     []string{"u"},
				EnvVars:     []string{"BATCMD_BACKEND_URL"},
				Destination: &cmd.backendURL,
				Usage:       "The url to connect backend service such as ssh",
				Required:    true,
			},
		},
		Action: cmd.Run,
	}
}

type ServeCommand struct {
	listen     string
	backendURL string
}

func (c *ServeCommand) Run(cCtx *cli.Context) error {
	return http.ListenAndServe(c.listen, c.backendURL)
}
