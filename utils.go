package goavrobuf

func BuildinTypeToString(t interface{}) string {
	s, ok := t.(string)
	if !ok {
		panic("cannot parse primitive to string type")
	}
	return s
}

func GetRepeatedTypeName(t interface{}) (string, string) {
	a, ok := t.([]interface{})
	if !ok {
		panic("cannot parse repeated type to slice")
	}
	for _, v := range a {
		if m, ok := v.(map[string]interface{}); ok {
			return m["type"].(string), m["items"].(string)
		}
	}
	panic("cannot determine type and name")
}

func GetEnumSymbols(t interface{}) (string, []string) {
	m, ok := t.(map[string]interface{})
	if !ok {
		panic("cannot parse enum type to map")
	}
	if m["type"].(string) != "enum" {
		panic("not a enum type")
	}
	name, ok := m["name"]
	if !ok {
		panic("cannot get enum name")
	}
	symbols, ok := m["symbols"]
	if !ok {
		panic("cannot get symbols")
	}
	ss, ok := symbols.([]interface{})
	if !ok {
		panic("symbols is not a slice")
	}
	s := make([]string, len(ss))
	for _, v := range ss {
		s = append(s, v.(string))
	}
	return name.(string), s
}
