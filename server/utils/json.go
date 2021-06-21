package utils

import (
	"bytes"
	"encoding/json"
	"strings"
)

func FormatJson(data string) string {
	var out bytes.Buffer
	json.Indent(&out, []byte(data), "", "    ")
	return out.String()
}

func SerializeToJson(data interface{}) string {
	b, _ := json.Marshal(data)
	return string(b)
}

func UnserializeFromJson(data string, result interface{}) error {
	d := json.NewDecoder(strings.NewReader(data))
	d.UseNumber()
	return d.Decode(result)
}

func FormatStruct(data interface{}) string {
	result := SerializeToJson(data)
	return FormatJson(result)
}
