package formatter

import (
	"fmt"
	"strings"

	"github.com/tylercunnion/loglady/pkg/values"
)

type Formatter struct {
}

func (f *Formatter) Format(line map[string]interface{}) (string, error) {
	var level = values.Get(line, "level")
	var color = "\033[0m"
	if level == "ERROR" {
		color = "\033[0;31m"
	} else if level == "WARN" {
		color = "\033[0;33m"
	}

	return buildOutputString([]displayPair{
		displayPair{color, "%s"},
		displayPair{values.Get(line, "level"), "[%s] "},
		displayPair{values.Get(line, "mdc.application"), "(%s) "},
		displayPair{values.Get(line, "message"), "%s"},
	}), nil
}

type displayPair struct {
	value   interface{}
	display string
}

func buildOutputString(displayValues []displayPair) string {
	var formatStrings = make([]string, 0, len(displayValues))
	var outputVals = make([]interface{}, 0, len(displayValues))

	for _, p := range displayValues {
		if p.value != nil {
			formatStrings = append(formatStrings, p.display)
			outputVals = append(outputVals, p.value)
		}
	}

	return fmt.Sprintf(strings.Join(formatStrings, ""), outputVals...)
}
