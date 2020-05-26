package goavrobuf

import (
	"encoding/json"
	"fmt"
)

func NewSchema(schemaSpecification string) (TreeNoder, error) {
	var schema interface{}
	if err := json.Unmarshal([]byte(schemaSpecification), &schema); err != nil {
		return nil, fmt.Errorf("cannot unmarshal schema JSON: %s", err)
	}
	schemaMap, ok := schema.(map[string]interface{})
	if !ok {
		panic("cannot convert json to map")
	}
	root := NewTreeNode(schemaMap)
	return root, nil
}


