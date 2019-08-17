package loglady

import (
	"io"

	"github.com/tylercunnion/loglady/pkg/formatter"
	"github.com/tylercunnion/loglady/pkg/scanner"
)

var data = `
timestamp:
  field: "@timestamp"
  parse: "2006-01-02T15:04:05.999Z07:00"
  format: "Jan _2 15:04:05.0Z"
levels:
  info:
   - INFO
  warn: 
   - WARN
  error:
   - ERROR
   - err
fields: 
  - key: "#timestamp"
    format: "%s "
  - key: level
    format: "[%s] "
  - key: mdc.application
    format: "(%s) "
  - key: message
    format: "%s"
`

func Run(r io.Reader) error {
	var logFmt, err = formatter.GetFormatter([]byte(data))
	if err != nil {
		return err
	}
	var scanner = scanner.NewScanner(r, logFmt)
	err = scanner.Scan()
	if err != nil {
		return err
	}
	return nil
}
