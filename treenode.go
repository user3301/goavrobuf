package goavrobuf

import "fmt"

type TreeNoder interface{
	Name() string
	NodeType() interface{}
	Fields() ([]TreeNoder, error)
}

type treeNode struct {
	name string
	nodeType interface{}
	rawFields interface{}
}

func (t treeNode) NodeType() interface{} {
	return t.nodeType
}

func newTreeNode(name string, nodeType, rawFields interface{}) TreeNoder {
	return &treeNode{
		name:      name,
		rawFields: rawFields,
		nodeType: nodeType,
	}
}

func (t treeNode) Name() string {
	return t.name
}

func (t treeNode) Fields() ([]TreeNoder, error) {
	if t.rawFields == nil {
		return nil, nil
	}
	if s, ok := t.rawFields.([]interface{}); !ok {
		return nil, fmt.Errorf("fields type ought to be []interface{}; received %T", t.rawFields)
	} else {
		fields := make([]TreeNoder, len(s))
		for i, v := range s {
			obj, ok := v.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("fields type ought to be map[string]interface{}; received %T", v)
			}
			nameRaw, ok := obj["name"]
			if !ok {
				return nil, fmt.Errorf("missing name %v", obj)
			}
			name, ok := nameRaw.(string)
			if !ok {
				return nil, fmt.Errorf("name type ought to be string; received %T", nameRaw)
			}
			rawType, ok := obj["type"]
			if !ok {
				return nil, fmt.Errorf("missing type %v", obj)
			}
			var f interface{}
			if ff, ok := obj["fields"]; ok {
				f = ff
			}
			fields[i] = newTreeNode(name, rawType, f)
		}
		return fields, nil
	}
}
