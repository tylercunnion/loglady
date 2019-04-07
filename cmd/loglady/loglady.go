package main

import (
	"bufio"
	"fmt"
	"local/logparse/pkg/scanner"
	"os"
)

func main() {
	scanner.Scan(bufio.NewReader(os.Stdin))
	fmt.Println("done")
}
