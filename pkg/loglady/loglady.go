package loglady

import (
	"io"
	"io/ioutil"

	"github.com/tylercunnion/loglady/pkg/formatter"
	"github.com/tylercunnion/loglady/pkg/scanner"
)

// Run is the main entrypoint for the loglady app, accepting a Reader to get
// input from and a Writer to send the formatted output to.
func Run(r io.Reader, w io.Writer) error {
	// TODO: Customize the config location
	data, err := ioutil.ReadFile("loglady.yaml")
	if err != nil {
		return err
	}

	logFmt, err := formatter.NewLineFormat(data)
	if err != nil {
		return err
	}

	var scanner = scanner.NewScanner(r)
	for scanner.Scan() {
		fields, err := scanner.Fields()
		if err != nil {
			return err
		}

		formattedString, err := logFmt.FormatLine(fields)
		_, err = w.Write([]byte(formattedString + "\n"))
		if err != nil {
			return err
		}
	}

	if err = scanner.Err(); err != nil {
		return err
	}
	return nil
}
