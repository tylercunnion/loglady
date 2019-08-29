package parser

import (
	"encoding/json"
)

type Parser struct{}

// Parse takes a raw line from the log and returns a parsed map of fields. If
// the input line is not valid JSON, it returns a map with one field,
// "message", containing the entire line.
func (p *Parser) Parse(line []byte) (*Obj, error) {
	var obj map[string]interface{}
	if !(json.Valid(line)) {
		obj = parseRaw(line)
	} else {
		obj = parseJSON(line)
	}
	return &Obj{obj}, nil
}

func parseJSON(line []byte) map[string]interface{} {
	var m map[string]interface{}
	json.Unmarshal(line, &m)

	return m
}

func parseRaw(line []byte) map[string]interface{} {
	var m = make(map[string]interface{})
	m["message"] = string(line)
	return m
}
