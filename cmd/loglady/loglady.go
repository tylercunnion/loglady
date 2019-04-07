package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/tylercunnion/loglady/pkg/formatter"
	"github.com/tylercunnion/loglady/pkg/scanner"
)

func main() {
	var logFmt = &formatter.Formatter{}
	scanner.Scan(bufio.NewReader(os.Stdin), logFmt)
	fmt.Println("done")
}
