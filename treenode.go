package goavrobuf

type JsonTreeNoder interface {
	GetName() string
	GetNodeType() NodeType
	GetFields() []JsonTreeNoder
	GetTypeData() interface{}
}

type JsonTreeNode struct {
	name     string
	dataType interface{}
	fields   interface{}
}

type JsonTreeRooter interface {
	JsonTreeNoder
	GetNamespace() string
}

type JsonTreeRoot struct {
	name      string
	namespace string
	dataType  interface{}
	fields    interface{}
}

func (node JsonTreeRoot) GetName() string {
	return node.name
}

func (node JsonTreeRoot) GetNodeType() NodeType {
	if node.dataType.(string) != "record" {
		panic("root is not record type")
	}
	return Record
}

func NewJsonTreeNode(n string, data, fields interface{}) JsonTreeNoder {
	return &JsonTreeNode{
		name:     n,
		dataType: data,
		fields:   fields,
	}
}

func NewJsonTreeRoot(n, ns string, data, fields interface{}) JsonTreeRooter {
	return &JsonTreeRoot{
		name:      n,
		namespace: ns,
		dataType:  data,
		fields:    fields,
	}
}

func (node JsonTreeNode) GetName() string {
	return node.name
}

func (node JsonTreeRoot) GetFields() []JsonTreeNoder {
	var fields []JsonTreeNoder
	if node.GetNodeType() == Primitive {
		return fields
	}
	sm, ok := node.fields.([]interface{})
	if !ok {
		panic("cannot parse fields")
	}
	for _, cc := range sm {
		n, ok := cc.(map[string]interface{})
		if !ok {
			panic("cannot get field")
		}

		v, ok := n["type"].(map[string]interface{})
		var f interface{}
		if ok {
			f = v["fields"]
		}
		jsonTreeNode := NewJsonTreeNode(n["name"].(string), n["type"], f)
		fields = append(fields, jsonTreeNode)
	}
	return fields
}

func (node JsonTreeRoot) GetTypeData() interface{} {
	return node.dataType
}

func (node JsonTreeRoot) GetNamespace() string {
	return node.namespace
}

func (node JsonTreeNode) GetNodeType() NodeType {
	switch node.dataType.(type) {
	case map[string]interface{}:
		v, ok := node.dataType.(map[string]interface{})["type"]
		if !ok {
			panic("cannot determine node type")
		}
		switch v.(string) {
		case "record":
			return Record
		case "enum":
			return Enum
		default:
			return Unknown
		}
	case []interface{}:
		return Repeated
	case string:
		return Primitive
	default:
		return Unknown
	}
}

func (node JsonTreeNode) GetFields() []JsonTreeNoder {
	var fields []JsonTreeNoder

	if seen.Contain(node.GetName()) {
		return fields
	}
	if node.GetNodeType() == Primitive {
		return fields
	}
	sm, ok := node.fields.([]interface{})
	if !ok {
		return fields
	}
	for _, cc := range sm {
		node, ok := cc.(map[string]interface{})
		if !ok {
			panic("cannot get field")
		}
		f := node["fields"]
		jsonTreeNode := NewJsonTreeNode(node["name"].(string), node["type"], f)
		fields = append(fields, jsonTreeNode)
	}
	return fields
}

func (node JsonTreeNode) GetTypeData() interface{} {
	return node.dataType
}
