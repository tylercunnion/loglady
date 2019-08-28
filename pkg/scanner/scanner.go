package scanner

import (
	"bufio"
	"io"

	"github.com/tylercunnion/loglady/pkg/parser"
)

type lineFormatter interface {
	FormatLine(map[string]interface{}) (string, error)
}

type Scanner struct {
	scan *bufio.Scanner
	p    *parser.Parser
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{
		scan: bufio.NewScanner(r),
		p:    &parser.Parser{},
	}
}

func (s *Scanner) Scan() bool {
	return s.scan.Scan()
}
func (s *Scanner) Err() error {
	return s.scan.Err()
}

func (s *Scanner) Fields() (map[string]interface{}, error) {
	return s.p.Parse(s.scan.Bytes())
}
