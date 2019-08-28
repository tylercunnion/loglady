package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/tylercunnion/loglady/pkg/loglady"
)

func main() {
	fmt.Println("my log has a message for you...")
	err := loglady.Run(bufio.NewReader(os.Stdin), os.Stdout)
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
}
