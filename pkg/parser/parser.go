package parser

import (
	"encoding/json"
	"errors"
)

type Parser struct {
}

func (p *Parser) Parse(line []byte) (map[string]interface{}, error) {
	if !(json.Valid(line)) {
		return nil, errors.New("line is not json")
	}

	var m map[string]interface{}
	json.Unmarshal(line, &m)

	return m, nil
}
