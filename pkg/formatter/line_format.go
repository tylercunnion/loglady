package formatter

import (
	"fmt"
	"strings"

	"github.com/logrusorgru/aurora"
	"github.com/tylercunnion/loglady/pkg/parser"

	"gopkg.in/yaml.v2"
)

type fieldDef struct {
	Key    string `yaml:"key"`
	Format string `yaml:"format"`
	Type   string `yaml:"type"`
	Color  string `yaml:"color"`
}

type displayPair struct {
	value   interface{}
	display string
}

// LineFormat describes the fields to include in an output line and how to
// format them.
type LineFormat struct {
	Fields     []fieldDef `yaml:"fields"`
	LevelNames levelNames `yaml:"levels"`
	Timestamp  *tsConfig  `yaml:"timestamp,omitempty"`
	au         aurora.Aurora
}

// NewLineFormat parses the config yaml given by `config`and returns a
// LineFormat which includes the default level names and format strings.
func NewLineFormat(config []byte) (*LineFormat, error) {
	var f = &LineFormat{}
	var err = yaml.Unmarshal(config, f)
	for fi, fp := range f.Fields {
		if fp.Format == "" {
			f.Fields[fi].Format = "%s"
		}
	}
	f.LevelNames = mergeDefaults(f.LevelNames)
	f.au = aurora.NewAurora(true)

	return f, err
}

// FormatLine takes a parsed line and returns a formatted string ready for
// output.
func (f *LineFormat) FormatLine(line *parser.Obj) (string, error) {
	var displays = make([]displayPair, 0)
	var lineLevel logLevel

	for _, fp := range f.Fields {
		var err error
		var rawValue = line.Get(fp.Key)
		if rawValue == nil {
			continue
		}

		var color = colorValue(fp.Color)
		var displayValue interface{}

		switch fp.Type {
		case "timestamp":
			displayValue, err = timestampFormat(rawValue.(string), *f.Timestamp)
			if err != nil {
				return "", err
			}
		case "level":
			lineLevel = getLevel(line.GetAsString(fp.Key), f.LevelNames)
			fallthrough
		default:
			displayValue = rawValue
		}

		displayValue = f.au.Colorize(rawValue, color)
		displays = append(displays, displayPair{displayValue, fp.Format})
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
