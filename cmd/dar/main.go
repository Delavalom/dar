package main

import (
	"os"

	"github.com/delavalom/dar/internal/cmd/root"
)

type exitCode int

const (
	exitOK      exitCode = 0
	exitError   exitCode = 1
	exitCancel  exitCode = 2
	exitAuth    exitCode = 4
	exitPending exitCode = 8
)

func init() {
	os.Getenv("DAR_PATH")
}

func run() exitCode {
	rootCommand, err := root.NewRootCommand()
	if err != nil {
		return exitError
	}

	if err = rootCommand.Execute(); err != nil {
		return exitError
	}

	return exitOK
}

func main() {
	code := run()
	os.Exit(int(code))
}
