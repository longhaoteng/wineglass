package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func JsonIndent(data []byte) string {
	if len(data) == 0 {
		return ""
	}
	var j bytes.Buffer
	err := json.Indent(&j, data, "", "    ")
	if err != nil {
		return fmt.Sprintf("json format errors: %v", err)
	}
	return j.String()
}
