package scanner

import (
	"bufio"
	"fmt"
	"io"

	"github.com/tylercunnion/loglady/pkg/parser"
)

type lineFormatter interface {
	FormatLine(map[string]interface{}) (string, error)
}

type Scanner struct {
	r io.Reader
	f lineFormatter
}

func NewScanner(r io.Reader, f lineFormatter) *Scanner {
	return &Scanner{
		r: r,
		f: f,
	}
}

func (s *Scanner) Scan() error {
	return scan(s.r, s.f)
}

func scan(r io.Reader, f lineFormatter) error {
	var s = bufio.NewScanner(r)
	var p = &parser.Parser{}

	for s.Scan() {
		var obj, err = p.Parse(s.Bytes())
		if err != nil {
			return err
		}

		str, err := f.FormatLine(obj)
		if err != nil {
			return err
		}
		fmt.Println(str)
	}

	return s.Err()
}
