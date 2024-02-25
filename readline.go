package main

import (
	"bufio"
	"os"
)

func Read_line() string {
	reader := bufio.NewReader(os.Stdin)

	line, err := reader.ReadString('\n')
	Must(err)

	// if len(strings.TrimSpace(line)) == 0 {
	// 	return ''
	// }

	return line
}
