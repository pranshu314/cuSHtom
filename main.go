package main

import (
	"fmt"
	"os"
	"os/user"
	"strings"

	"github.com/mattn/go-isatty"
)

func main() {
	if isatty.IsTerminal(os.Stdin.Fd()) {
		currUser, err := user.Current()
		Must(err)
		usr_name := currUser.Username

		host_name, err := os.Hostname()
		Must(err)

		clear()

		fmt.Printf("\n\n---------Welcome to cuSHtom Shell----------\n\n")

		cshm_loop(usr_name, host_name)

	} else {
		fmt.Println("Is Not a Terminal")
	}
}

func cshm_loop(usr string, host string) {
	token_dlim := []rune{' ', '\t', '\r', '\n', '\a'}

	status := 1

	for {
		curr_dir, err := os.Getwd()
		Must(err)
		tmp := strings.Split(curr_dir, "/")
		pwd := tmp[len(tmp)-1]

		fmt.Printf("[%s@%s]:-%s>>> ", usr, host, pwd)
		line := Read_line()
		args := Split_line(line, token_dlim)
		status = Execute(args)

		if status == 0 {
			break
		}
	}
}

func clear() {
	fmt.Printf("\033[H\033[J")
}

func Must(err error) {
	if err != nil {
		fmt.Printf("Error:- %s\n", err)
	}
}
