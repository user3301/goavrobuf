package goavrobuf

type TreeNoder interface {
	GetName() string
	GetChildren() []map[string]interface{}
}

type TreeNode struct {
	Name     string
	Type     NodeType
	Children []map[string]interface{}
}

func NewTreeNode(node map[string]interface{}) TreeNoder {
	t, ok := node["type"]
	if !ok {
		panic("cannot determine type of node")
	}
	n, ok := node["name"]
	if !ok {
		panic("cannot determine name of node")
	}
	nodeType := GetType(t)
	var children []map[string]interface{}
	if c, ok := node["fields"]; ok {
		children = getChildren(c)
	}
	return &TreeNode{
		Name:     n.(string),
		Type:     nodeType,
		Children: children,
	}
}

func (t *TreeNode) GetName() string {
	return t.Name
}

func (t TreeNode) GetChildren() []map[string]interface{} {
	return t.Children
}

func GetType(t interface{}) NodeType {
	switch t.(type) {
	case string:
		if t == "record" {
			return Root
		}
		return Primitive
	case map[string]interface{}:
		return Record
	default:
		panic("unknown schema type")
	}
}

func getChildren(c interface{}) []map[string]interface{} {
	sm , ok := c.([]interface{})
	if !ok {
		panic("cannot parse fields")
	}
	var children []map[string]interface{}
	for _, cc := range sm {
		node, ok := cc.(map[string]interface{})
		if !ok {
			panic("cannot get child")
		}
		children = append(children, node)
	}
	return children
}