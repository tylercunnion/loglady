package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/tylercunnion/loglady/pkg/loglady"
)

func main() {
	fmt.Println("my log has a message for you...")
	loglady.Run(bufio.NewReader(os.Stdin))
}
