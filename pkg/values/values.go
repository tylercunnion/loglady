package values

import (
	"strings"
)

func Get(obj map[string]interface{}, name string) interface{} {
	var splitName = strings.SplitN(name, ".", 2)
	if len(splitName) == 1 {
		return obj[name]
	}

	var subObject, ok = obj[splitName[0]].(map[string]interface{})
	if !ok {
		return nil
	}
	return Get(subObject, splitName[1])
}
