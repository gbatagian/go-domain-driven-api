package utils

import (
	"bytes"
	"encoding/json"
	"io"
)

func ToJsonString(m map[string]interface{}) string {

	str, _ := json.Marshal(m)
	return string(str)

}

func ToJsonBytesStream(m map[string]interface{}) io.Reader {

	bytes_buffer := new(bytes.Buffer)
	json.NewEncoder(bytes_buffer).Encode(m)

	return bytes_buffer

}
