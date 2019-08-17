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

type Scanner struct {
	r io.Reader
	f formatter
}

func NewScanner(r io.Reader, f formatter) *Scanner {
	return &Scanner{
		r: r,
		f: f,
	}
}

func (s *Scanner) Scan() error {
	return scan(s.r, s.f)
}

func scan(r io.Reader, f formatter) error {
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
