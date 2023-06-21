package main

import (
	"fmt"
	"os"

	"github.com/ssh-kit/hsh/internal/cmd"
)

func main() {
	if err := cmd.NewRootCmd().Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
