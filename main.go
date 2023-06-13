package main

import (
	"fmt"
	"os"

	"github.com/batx-dev/batcmd/internal/cmd"
)

func main() {
	if err := cmd.NewRootCmd().Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
