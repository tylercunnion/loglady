package parser

import (
	"strings"
)

type Obj struct {
	parseMap map[string]interface{}
}

// Get returns an item from the given map. Since the map is generated from JSON,
// and therefore can have arbitrarily nested data structures, you can address
// these by using dotted notation - "name.last" for instance.
func (o *Obj) Get(name string) interface{} {
	return get(o.parseMap, name)
}

func (o *Obj) GetAsBool(name string) bool {
	return o.Get(name).(bool)
}

func (o *Obj) GetAsString(name string) string {
	return o.Get(name).(string)
}

func (o *Obj) GetAsFloat(name string) float64 {
	return o.Get(name).(float64)
}

func (o *Obj) GetAsSlice(name string) []interface{} {
	return o.Get(name).([]interface{})
}

func get(obj map[string]interface{}, name string) interface{} {
	var splitName = strings.SplitN(name, ".", 2)
	if len(splitName) == 1 {
		return obj[name]
	}

	var subObject, ok = obj[splitName[0]].(map[string]interface{})
	if !ok {
		return nil
	}
	return get(subObject, splitName[1])
}
