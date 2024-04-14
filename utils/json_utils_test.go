package utils

import (
	"bytes"
	"strings"
	"testing"
)

func TestToJsonString(t *testing.T) {
	// arrange
	sample_map := map[string]interface{}{
		"key1": "value1",
		"key2": 1,
		"key3": 1.5999,
		"key4": false,
		"key5": nil,
	}
	expectedJsonString := `{"key1":"value1","key2":1,"key3":1.5999,"key4":false,"key5":null}`

	// act
	jsonString := ToJsonString(sample_map)

	// assert
	if jsonString != expectedJsonString {
		t.Errorf("Unexpected json string: %v", jsonString)
	}

}

func TestToJsonBytesStream(t *testing.T) {
	//arrange
	sample_map := map[string]interface{}{
		"key1": "value1",
		"key2": 1,
		"key3": 1.5999,
		"key4": false,
		"key5": nil,
	}

	expectedJsonString := `{"key1":"value1","key2":1,"key3":1.5999,"key4":false,"key5":null}`

	// act
	jsonBytesStream := ToJsonBytesStream(sample_map)
	jsonBytesString := jsonBytesStream.(*bytes.Buffer).String()

	if strings.TrimSuffix(jsonBytesString, "\n") != expectedJsonString {
		t.Errorf("Unexpected json string: %v", jsonBytesString)
	}
}
