package formatter

import (
	"fmt"
	"strings"

	"github.com/tylercunnion/loglady/pkg/values"
	"gopkg.in/yaml.v2"
)

type formatPair struct {
	Key    string `yaml:"key"`
	Format string `yaml:"format"`
	Type   string `yaml:"type"`
}

type displayPair struct {
	value   interface{}
	display string
}

type LineFormat struct {
	Fields     []formatPair `yaml:"fields"`
	LevelNames levelNames   `yaml:"levels"`
	Timestamp  *tsConfig    `yaml:"timestamp,omitempty"`
}

func GetLineFormat(config []byte) (*LineFormat, error) {
	var f = &LineFormat{}
	var err = yaml.Unmarshal(config, f)
	for fi, fp := range f.Fields {
		if fp.Format == "" {
			f.Fields[fi].Format = "%s"
		}
	}
	f.LevelNames = mergeDefaults(f.LevelNames)

	return f, err
}

func (f *LineFormat) FormatLine(line map[string]interface{}) (string, error) {
	var displays = make([]displayPair, 0)
	var lineLevel logLevel

	for _, fp := range f.Fields {
		switch fp.Type {
		case "timestamp":
			val, err := TimestampFormat(values.Get(line, fp.Key).(string), *f.Timestamp)
			if err != nil {
				return "", err
			}
			displays = append(displays, displayPair{val, fp.Format})
		case "level":
			lineLevel = getLevel(values.Get(line, fp.Key).(string), f.LevelNames)
			fallthrough
		default:
			displays = append(displays, displayPair{values.Get(line, fp.Key), fp.Format})
		}
	}

	return buildOutputString(displays, lineLevel), nil
}

func buildOutputString(displayValues []displayPair, lineLevel logLevel) string {
	var formatStrings = make([]string, 0, len(displayValues))
	var outputVals = make([]interface{}, 0, len(displayValues))

	for _, p := range displayValues {
		if p.value != nil {
			formatStrings = append(formatStrings, p.display)
			outputVals = append(outputVals, p.value)
		}
	}

	return fmt.Sprintf(strings.Join(formatStrings, " "), outputVals...)
}
