package formatter

import "strings"

type logLevel int

// Constants that represent the various logging levels.
const (
	DEBUG logLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

type levelNames struct {
	Debug []string
	Info  []string
	Warn  []string
	Error []string
	Fatal []string
}

func mergeDefaults(names levelNames) levelNames {
	names.Debug = mergeDefaultList([]string{"debug"}, names.Debug)
	names.Info = mergeDefaultList([]string{"info"}, names.Info)
	names.Warn = mergeDefaultList([]string{"warn", "warning"}, names.Warn)
	names.Error = mergeDefaultList([]string{"error"}, names.Error)
	names.Fatal = mergeDefaultList([]string{"fatal"}, names.Fatal)
	return names
}

func mergeDefaultList(defaults []string, custom []string) []string {
	if len(custom) == 0 {
		return defaults
	}
	return append(defaults, custom...)
}

func listContains(list []string, find string) bool {
	find = strings.ToLower(find)
	for _, s := range list {
		if strings.ToLower(s) == find {
			return true
		}
	}
	return false
}

func getLevel(levelValue string, names levelNames) logLevel {
	var l logLevel
	if listContains(names.Fatal, levelValue) {
		l = FATAL
	} else if listContains(names.Error, levelValue) {
		l = ERROR
	} else if listContains(names.Warn, levelValue) {
		l = WARN
	} else if listContains(names.Info, levelValue) {
		l = INFO
	} else if listContains(names.Debug, levelValue) {
		l = DEBUG
	}

	return l
}
