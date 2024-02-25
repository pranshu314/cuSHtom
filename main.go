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
		curr_dir, err := os.Getwd()
		Must(err)
		tmp := strings.Split(curr_dir, "/")
		curr_dir = tmp[len(tmp)-1]

		currUser, err := user.Current()
		Must(err)
		usr_name := currUser.Username

		host_name, err := os.Hostname()
		Must(err)

		clear()
		cshm_loop(usr_name, host_name, curr_dir)

	} else {
		fmt.Println("Is Not a Terminal")
	}
}

func cshm_loop(usr string, host string, pwd string) {
	token_dlim := []rune{' ', '\t', '\r', '\n', '\a'}

	status := 1

	for {

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
		panic(err)
	}
}
