package parser

import (
	"encoding/json"
)

type Parser struct {
}

func (p *Parser) Parse(line []byte) (map[string]interface{}, error) {
	if !(json.Valid(line)) {
		return parseRaw(line)
	}
	return parseJson(line)
}

func parseJson(line []byte) (map[string]interface{}, error) {
	var m map[string]interface{}
	json.Unmarshal(line, &m)

	return m, nil
}

func parseRaw(line []byte) (map[string]interface{}, error) {
	var m = make(map[string]interface{})
	m["message"] = string(line)
	return m, nil
}
