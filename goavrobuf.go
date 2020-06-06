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
		return nil, fmt.Errorf("schema is not an json object %v", schema)
	}
	_, ok = schemaMap["name"]
	if !ok {
		return nil, fmt.Errorf("missing name %v", schemaMap)
	}
	//TODO

	//return newTreeNode(n.(string), schema), nil
	return nil, nil
}

//func ToProto3(root TreeNoder) {
//	seen := make(RecordSet)
//	queue := make([]TreeNoder, 0)
//	queue = append(queue, root)
//	for len(queue) != 0 {
//		size := len(queue)
//		for i := 0; i < size; i++ {
//			current := queue[0]
//			queue = queue[1:]
//			switch schemaType := current.GetProtoType().(type) {
//			case map[string]interface{}:
//				handleMapType(queue, seen)
//			}
//		}
//	}
//}

