package utils

import (
	"bytes"
	"encoding/json"
	"io"
)

func ToJsonString[E any](m map[string]E) string {
	str, _ := json.Marshal(m)

	return string(str)
}

func ToJsonBytesStream[E any](m map[string]E) io.Reader {
	bytes_buffer := new(bytes.Buffer)
	json.NewEncoder(bytes_buffer).Encode(m)

	return bytes_buffer
}
