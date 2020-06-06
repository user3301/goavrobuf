package goavrobuf

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTreeNode_Name(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		expected := "record"
		node := newTreeNode(expected, nil, nil)
		actual := node.Name()
		assert.Equal(t, expected, actual)
	})
}

func TestTreeNode_NodeType(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		expected := []string{"null", "boolean"}
		node := newTreeNode("long", expected, nil)
		actual := node.NodeType()
		assert.IsType(t, []string{}, actual)
	})
}

func TestTreeNode_Fields(t *testing.T) {
	t.Run("NilRawFields", func(t *testing.T) {
		node := newTreeNode("nil", nil, nil)
		fields, err := node.Fields()
		require.NoError(t, err)
		assert.Nil(t, fields)
	})
	t.Run("FieldsNotASlice", func(t *testing.T) {
		jsonStr := "{\"fields\": \"user3301\"}"
		var jsonObj map[string]interface{}
		err := json.Unmarshal([]byte(jsonStr), &jsonObj)
		require.NoError(t, err)
		node := newTreeNode("foo", nil, jsonObj)
		fields, err := node.Fields()
		assert.Nil(t, fields)
		assert.EqualError(
			t,
			err,
			"fields type ought to be []interface{}; received map[string]interface {}",
		)
	})
	t.Run("FieldNotMap", func(t *testing.T) {
		jsonStr := "{\"fields\": [\"null\"]}"
		var jsonObj map[string]interface{}
		err := json.Unmarshal([]byte(jsonStr), &jsonObj)
		require.NoError(t, err)
		node := newTreeNode("foo", nil, jsonObj["fields"])
		fields, err := node.Fields()
		assert.Nil(t, fields)
		assert.EqualError(
			t,
			err,
			"fields type ought to be map[string]interface{}; received string",
		)
	})
	t.Run("FieldMissingNameField", func(t *testing.T) {
		jsonStr := "{\"fields\": [{}]}"
		var jsonObj map[string]interface{}
		err := json.Unmarshal([]byte(jsonStr), &jsonObj)
		require.NoError(t, err)
		node := newTreeNode("foo", nil, jsonObj["fields"])
		fields, err := node.Fields()
		assert.Nil(t, fields)
		assert.EqualError(t, err, "missing name map[]")
	})
	t.Run("FieldNameNotStringType", func(t *testing.T) {
		jsonStr := "{\"fields\": [{\"name\":{}}]}"
		var jsonObj map[string]interface{}
		err := json.Unmarshal([]byte(jsonStr), &jsonObj)
		require.NoError(t, err)
		node := newTreeNode("foo", nil, jsonObj["fields"])
		fields, err := node.Fields()
		assert.Nil(t, fields)
		assert.EqualError(
			t,
			err,
			"name type ought to be string; received map[string]interface {}",
		)
	})
	t.Run("FieldMissingType", func(t *testing.T) {
		jsonStr := "{\"fields\": [{\"name\":\"bar\"}]}"
		var jsonObj map[string]interface{}
		err := json.Unmarshal([]byte(jsonStr), &jsonObj)
		require.NoError(t, err)
		node := newTreeNode("foo", nil, jsonObj["fields"])
		fields, err := node.Fields()
		assert.Nil(t, fields)
		assert.EqualError(t, err, "missing type map[name:bar]")
	})
	t.Run("Success", func(t *testing.T) {
		jsonStr := "{\"fields\": [{\"name\":\"bar\", \"type\": \"record\", \"fields\":[]}]}"
		var jsonObj map[string]interface{}
		err := json.Unmarshal([]byte(jsonStr), &jsonObj)
		require.NoError(t, err)
		node := newTreeNode("foo", nil, jsonObj["fields"])
		fields, err := node.Fields()
		assert.NoError(t,err)
		assert.Equal(t, 1, len(fields))
	})
}
