package formatter

import (
	"time"
)

type tsConfig struct {
	Parse  string `yaml:"parse"`
	Format string `yaml:"format"`
}

func TimestampFormat(value string, config tsConfig) (string, error) {
	var parsedTime, err = time.Parse(config.Parse, value)
	if err != nil {
		return "", err
	}
	if config.Format != "" {
		return parsedTime.Format(config.Format), nil
	}
	return parsedTime.String(), nil
}
