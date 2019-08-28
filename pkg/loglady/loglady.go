package loglady

import (
	"io"

	"github.com/tylercunnion/loglady/pkg/formatter"
	"github.com/tylercunnion/loglady/pkg/scanner"
)

var data = `
timestamp:
  parse: "2006-01-02T15:04:05.999Z07:00"
  format: "Jan _2 15:04:05.0Z"
levels:
  error:
   - err
fields: 
  - key: "@timestamp"
    type: timestamp
  - key: level
    format: "[%s]"
    type: level
  - key: mdc.application
    format: "(%s)"
  - key: message
`

func Run(r io.Reader) error {
	var logFmt, err = formatter.GetLineFormat([]byte(data))
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
