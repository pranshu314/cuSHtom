package main

import (
	"os"
	"os/exec"
)

func Execute(args []string) int {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	Must(cmd.Run())

	return 1
}
