package main

import (
	"fmt"
	"os"
	"os/exec"
)

func Execute(args []string) int {
	if args[0] == "cd" {
		if len(args) == 1 {
			Must(os.Chdir("/"))
			return 1
		} else {
			Must(os.Chdir(args[1]))
			return 1
		}
	} else if args[0] == "exit" {
		fmt.Printf("\n\nBye, See you again\n\n")
		return 0
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	Must(cmd.Run())

	return 1
}
