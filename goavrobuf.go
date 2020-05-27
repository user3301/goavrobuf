package goavrobuf

import (
	"bytes"
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

func ToProto(root TreeNoder, destination string) {

}

// iterate the tree with breadth first search
// why bfs? as bfs gets nodes level by level
// so the proto message expends in correct order
func bfs(root TreeNoder) string {
	var buffer bytes.Buffer
	q := make([]TreeNoder, 0)
	q = append(q, root)

	for len(q) != 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			currentNode := q[0]
			q = q[1:]
			switch currentNode.GetType() {
			case Root:
				handleRootNode(&buffer, currentNode)
				for _, n := range currentNode.GetChildren() {
					q = append(q, NewTreeNode(n))
				}
			case Primitive:
				handlePrimitive(&buffer, currentNode)
			}
		}
	}
}

func handlePrimitive(buffer *bytes.Buffer, node TreeNoder) {
	defer buffer.WriteString("\n")

}

func handleRootNode(buffer *bytes.Buffer, node TreeNoder) {
	defer buffer.WriteString("\n")
	buffer.WriteString("syntax = \"proto3\"\n")
	buffer.WriteString(fmt.Sprintf("message %s {\n\t"))
	counter := 1
	for _, field := range node.GetChildren() {
		node := NewTreeNode(field)
		buffer.WriteString(fmt.Sprintf("%s %s = %d;", node.GetOriginalType().(string), node.GetName(), counter))
		counter++
	}
}
