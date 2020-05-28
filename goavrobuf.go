package goavrobuf

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

func NewSchema(schemaSpecification string) (JsonTreeNoder, error) {
	var schema interface{}
	if err := json.Unmarshal([]byte(schemaSpecification), &schema); err != nil {
		return nil, fmt.Errorf("cannot unmarshal schema JSON: %s", err)
	}
	schemaMap, ok := schema.(map[string]interface{})
	if !ok || schemaMap["type"].(string) != "record" {
		panic("cannot convert json to map")
	}
	return NewJsonTreeRoot(schemaMap["name"].(string), schemaMap["namespace"].(string), schemaMap["type"], schemaMap["fields"]), nil
}

func GenerateProto3(root JsonTreeNoder) string {
	var buffer bytes.Buffer
	buffer.WriteString("syntax = \"proto3\";\n")
	queue := make([]JsonTreeNoder, 0)
	//breadth first iterate
	queue = append(queue, root)
	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			current := queue[0]
			queue = queue[1:]
			switch current.GetNodeType() {
			case Record:
				b := handleRecord(current)
				buffer.Write(b.Bytes())
				for _, n := range current.GetFields() {
					if n.GetNodeType() == Record || n.GetNodeType() == Enum {
						queue = append(queue, n)
					}
				}
			case Enum:
				b := handleEnum(current)
				buffer.Write(b.Bytes())
			}
		}
	}
	fmt.Print("done!")
	return buffer.String()
}

func handleEnum(node JsonTreeNoder) bytes.Buffer {
	var buffer bytes.Buffer
	counter := 0
	buffer.WriteString(fmt.Sprintf("enum %s {\n\t", node.GetName()))
	_, eSymbols := GetEnumSymbols(node.GetTypeData())
	for _, v := range eSymbols {
		buffer.WriteString(fmt.Sprintf("%s_%s = %d;\n\t", node.GetName(), v, counter))
		counter++
	}
	buffer.WriteString("}\n")
	return buffer
}

func handleRecord(node JsonTreeNoder) bytes.Buffer {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("message %s {\n\t", node.GetName()))
	counter := 1
	for _, v := range node.GetFields() {
		switch v.GetNodeType() {
		case Record:
			buffer.WriteString(fmt.Sprintf("%s %s = %d;\n\t", v.GetName(), strings.ToLower(v.GetName()), counter))
			counter++
		case Primitive:
			buffer.WriteString(fmt.Sprintf("%s %s = %d;\n\t", BuildinTypeToString(v.GetTypeData()), strings.ToLower(v.GetName()), counter))
			counter++
		case Repeated:
			rType, rName := GetRepeatedTypeName(v.GetTypeData())
			if rName == "" {
				rName = v.GetName()
			}
			buffer.WriteString(fmt.Sprintf("repeated %s %s = %d;\n\t", rType, rName, counter))
			counter++
		case Enum:
			eName, _ := GetEnumSymbols(v.GetTypeData())
			buffer.WriteString(fmt.Sprintf("%s %s = %d;\n\t", eName, strings.ToLower(eName), counter))
			counter++
		case Unknown:
			continue
		}
	}
	buffer.WriteString("}\n")
	return buffer
}
