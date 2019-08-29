package parser

import (
	"encoding/json"

	gc "gopkg.in/check.v1"
)

type ObjSuite struct{}

var _ = gc.Suite(&ObjSuite{})

func (s *ObjSuite) TestGetFunctions(c *gc.C) {
	var data = `{
	"string": "hello",
	"number": 3.14,
	"bool": true,
	"slice": [1,2,3]
}
`
	var m map[string]interface{}
	json.Unmarshal([]byte(data), &m)
	var p = &Obj{m}
	c.Check(p.GetAsString("string"), gc.Equals, "hello")
	c.Check(p.GetAsFloat("number"), gc.Equals, 3.14)
	c.Check(p.GetAsBool("bool"), gc.Equals, true)
	c.Check(p.GetAsSlice("slice"), gc.DeepEquals, []interface{}{1.0, 2.0, 3.0})
}
