package scanner

import (
	"bufio"
	"io"

	"github.com/tylercunnion/loglady/pkg/parser"
)

type lineFormatter interface {
	FormatLine(map[string]interface{}) (string, error)
}

// Scanner reads over the input data line-by-line and returns parsed data.
type Scanner struct {
	rd    *bufio.Reader
	p     *parser.Parser
	bytes []byte
	err   error
}

// NewScanner prepares a new scanner from the given input Reader.
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{
		rd: bufio.NewReader(r),
		p:  &parser.Parser{},
	}
}

// Scan advances the scanner to the next line, returning true if a new line
// was found and false otherwise. This function does not return an error,
// instead check Err().
func (s *Scanner) Scan() bool {
	if s.err != nil {
		return false
	}

	s.bytes, s.err = s.rd.ReadBytes('\n')
	return s.err == nil
}

// Err returns any error encountered on the last scan.
func (s *Scanner) Err() error {
	if s.err == io.EOF {
		return nil
	}
	return s.err
}

// Fields returns parsed fields from the line read in the last Scan().
func (s *Scanner) Fields() (*parser.Obj, error) {
	return s.p.Parse(s.bytes)
}
