package main

import (
	"fmt"
	"os"

	"github.com/AIRE-labs/go-airenet/cmd/aire/launcher"
)

func main() {
	if err := launcher.Launch(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
