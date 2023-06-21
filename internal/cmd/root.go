package cmd

import (
	cli "github.com/urfave/cli/v2"
)

func NewRootCmd() *cli.App {
	return &cli.App{
		Name:  "hsh",
		Usage: "Manage HTTPShell Service",
		Commands: []*cli.Command{
			NewServeCommand(),
		},
	}
}
