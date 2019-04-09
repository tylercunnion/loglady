package loglady

import (
	"bufio"
	"fmt"
	"os"

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

func Run() {
	var logFmt, err = formatter.GetFormatter([]byte(data))
	if err != nil {
		fmt.Println(err)
	}
	err = scanner.Scan(bufio.NewReader(os.Stdin), logFmt)
	if err != nil {
		fmt.Println(err)
	}
}
