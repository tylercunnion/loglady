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
	rd    *bufio.Reader
	p     *parser.Parser
	bytes []byte
	err   error
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{
		rd: bufio.NewReader(r),
		p:  &parser.Parser{},
	}
}

func (s *Scanner) Scan() bool {
	s.bytes, s.err = s.rd.ReadBytes('\n')
	if s.err != nil {
		return false
	}
	return true
}
func (s *Scanner) Err() error {
	if s.err == io.EOF {
		return nil
	}
	return s.err
}

func (s *Scanner) Fields() (map[string]interface{}, error) {
	return s.p.Parse(s.bytes)
}
