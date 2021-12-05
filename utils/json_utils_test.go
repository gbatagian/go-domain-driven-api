package utils

import (
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToJsonString(t *testing.T) {

	sample_map := map[string]interface{}{
		"key1": "value1",
		"key2": 1,
		"key3": 1.5999,
		"key4": false,
		"key5": nil,
	}

	jsonString := ToJsonString(sample_map)
	expectedJsonString := `{"key1":"value1","key2":1,"key3":1.5999,"key4":false,"key5":null}`
	assert.Equal(t, jsonString, expectedJsonString)

}

func TestToJsonBytesStream(t *testing.T) {

	sample_map := map[string]interface{}{
		"key1": "value1",
		"key2": 1,
		"key3": 1.5999,
		"key4": false,
		"key5": nil,
	}

	jsonBytesStream := ToJsonBytesStream(sample_map)
	fmt.Print(jsonBytesStream)
	buffer := new(strings.Builder)
	io.Copy(buffer, jsonBytesStream)
	expectedJsonString := `{"key1":"value1","key2":1,"key3":1.5999,"key4":false,"key5":null}`
	assert.Equal(t, strings.Trim(buffer.String(), "\n"), expectedJsonString)

}
