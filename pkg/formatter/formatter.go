package formatter

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/tylercunnion/loglady/pkg/values"
	"gopkg.in/yaml.v2"
)

type formatPair struct {
	Key    string `yaml:"key"`
	Format string `yaml:"format"`
}

type levelNames struct {
	Debug []string
	Info  []string
	Warn  []string
	Error []string
}

type displayPair struct {
	value   interface{}
	display string
}

type tsConfig struct {
	Field  string `yaml:"field"`
	Parse  string `yaml:"parse"`
	Format string `yaml:"format"`
}

type Formatter struct {
	Fields     []formatPair `yaml:"fields"`
	LevelNames levelNames   `yaml:"levels"`
	Timestamp  *tsConfig    `yaml:"timestamp,omitempty"`
}

const (
	RESET  = "\033[0m"
	RED    = "\033[0;31m"
	YELLOW = "\033[0;33m"
)

func GetFormatter(config []byte) (*Formatter, error) {
	var f = &Formatter{}
	var err = yaml.Unmarshal(config, f)
	//fmt.Printf("%v\n", f)

	if f.Timestamp != nil {
		if f.Timestamp.Field == "" || f.Timestamp.Parse == "" {
			return nil, errors.New("timestamp not fully specified")
		}
	}

	return f, err
}

func listContains(list []string, find string) bool {
	for _, s := range list {
		if s == find {
			return true
		}
	}
	return false
}

func (f *Formatter) setupTime(line map[string]interface{}) (map[string]interface{}, error) {
	if f.Timestamp != nil {
		var ts, ok = values.Get(line, f.Timestamp.Field).(string)
		if ok {
			var parsedTime, err = time.Parse(f.Timestamp.Parse, ts)
			if err != nil {
				return nil, err
			}
			if f.Timestamp.Format != "" {
				line["#timestamp"] = parsedTime.Format(f.Timestamp.Format)
			} else {
				line["#timestamp"] = parsedTime
			}
		}
	}
	return line, nil
}

func (f *Formatter) Format(line map[string]interface{}) (string, error) {
	var displays = make([]displayPair, 0)

	line, err := f.setupTime(line)
	if err != nil {
		return "", err
	}

	level, ok := values.Get(line, "level").(string)
	if ok {
		var color string
		if listContains(f.LevelNames.Error, level) {
			color = RED
		} else if level == "WARN" {
			color = YELLOW
		}
		displays = append(displays, displayPair{color, "%s"})
	}

	for _, fp := range f.Fields {
		displays = append(displays, displayPair{values.Get(line, fp.Key), fp.Format})
	}
	displays = append(displays, displayPair{RESET, "%s"})

	return buildOutputString(displays), nil
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
