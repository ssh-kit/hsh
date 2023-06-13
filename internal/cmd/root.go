package cmd

import (
	cli "github.com/urfave/cli/v2"
)

func NewRootCmd() *cli.App {
	return &cli.App{
		Name:  "batcmd",
		Usage: "Manage BatCmd service",
		Commands: []*cli.Command{
			NewServeCommand(),
		},
	}
}
