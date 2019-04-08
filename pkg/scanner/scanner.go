package scanner

import (
	"bufio"
	"fmt"
	"io"

	"github.com/tylercunnion/loglady/pkg/parser"
)

type formatter interface {
	Format(map[string]interface{}) (string, error)
}

func Scan(r io.Reader, f formatter) error {
	var s = bufio.NewScanner(r)
	var p = &parser.Parser{}

	for s.Scan() {
		var obj, err = p.Parse(s.Bytes())
		if err != nil {
			return err
		}

		str, err := f.Format(obj)
		if err != nil {
			return err
		}
		fmt.Println(str)
	}

	return s.Err()
}
