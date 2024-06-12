package main

import (
	"dar/internal/cmd/root"
	"os"
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

func main() {
	code := mainRun()
	os.Exit(int(code))
}

func mainRun() exitCode {
	rootCommand, err := root.NewRootCommand()
	if err != nil {
		return exitError
	}

	if err = rootCommand.Execute(); err != nil {
		return exitError
	}

	return exitOK
}
