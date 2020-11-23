package mltq

import (
	"encoding/json"
)

func LogToQuery(log string) (string, error) {
	line := make(map[string]string)
	err := json.Unmarshal([]byte(log), &line)
	if err != nil {
		return "", err
	}
	cmd, ok := line["c"]
	if !ok || cmd != "COMMAND" {
		return "", nil
	}
	return line["attr"], nil
}
