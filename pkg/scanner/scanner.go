package scanner

import (
	"bufio"
	"fmt"
	"io"

	"local/logparse/pkg/formatter"
	"local/logparse/pkg/parser"
)

func Scan(r io.Reader) error {
	var s = bufio.NewScanner(r)
	var p = &parser.Parser{}
	var f = &formatter.Formatter{}

	for s.Scan() {
		var obj, err = p.Parse(s.Bytes())
		if err != nil {
			fmt.Println(s.Text())
			continue
		}

		str, err := f.Format(obj)
		fmt.Println(str)
	}

	return s.Err()
}
