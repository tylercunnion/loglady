package values

import (
	"strings"
)

// Get returns an item from the given map. Since the map is generated from JSON,
// and therefore can have arbitrarily nested data structures, you can address
// these by using dotted notation - "name.last" for instance.
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
