package parser

import (
	"testing"

	gc "gopkg.in/check.v1"
)

func Test(t *testing.T) { gc.TestingT(t) }

type ParserSuite struct {
	p *Parser
}

var _ = gc.Suite(&ParserSuite{})

func (s *ParserSuite) SetUpSuite(c *gc.C) {
	s.p = &Parser{}
}

func (s *ParserSuite) TestRawParser(c *gc.C) {
	var logline = "hello world"
	var results, err = s.p.Parse([]byte(logline))
	c.Assert(err, gc.IsNil)
	c.Check(results["message"], gc.Equals, logline)
}

func (s *ParserSuite) TestJsonParser(c *gc.C) {
	var logline = `{"level": "INFO", "message": "hello world", "id": 3}`
	var results, err = s.p.Parse([]byte(logline))
	c.Assert(err, gc.IsNil)
	c.Check(results["level"], gc.Equals, "INFO")
	c.Check(results["message"], gc.Equals, "hello world")
	c.Check(results["id"], gc.Equals, float64(3))
}
