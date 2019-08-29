package parser

import (
	"encoding/json"
)

type Parser struct {
}

// Parse takes a raw line from the log and returns a parsed map of fields. If
// the input line is not valid JSON, it returns a map with one field,
// "message", containing the entire line.
func (p *Parser) Parse(line []byte) (map[string]interface{}, error) {
	if !(json.Valid(line)) {
		return parseRaw(line)
	}
	return parseJSON(line)
}

func parseJSON(line []byte) (map[string]interface{}, error) {
	var m map[string]interface{}
	json.Unmarshal(line, &m)

	return m, nil
}

func parseRaw(line []byte) (map[string]interface{}, error) {
	var m = make(map[string]interface{})
	m["message"] = string(line)
	return m, nil
}
