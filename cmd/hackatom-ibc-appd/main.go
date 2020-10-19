package main

import (
	"os"

	"github.com/mconcat/hackatom-ibc-app/cmd/hackatom-ibc-appd/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := cmd.Execute(rootCmd); err != nil {
		os.Exit(1)
	}
}
