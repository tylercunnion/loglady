package loglady

import (
	"io"
	"io/ioutil"

	"github.com/tylercunnion/loglady/pkg/formatter"
	"github.com/tylercunnion/loglady/pkg/scanner"
)

func Run(r io.Reader) error {
	data, err := ioutil.ReadFile("loglady.yaml")
	if err != nil {
		return err
	}

	logFmt, err := formatter.GetLineFormat(data)
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
