package mltq

import (
	"encoding/json"
	"fmt"
)

type LogLine struct {
	C    string `json:"c"`
	Attr struct {
		Type    string `json:"type"`
		Command struct {
			Find      string          `json:"find"`
			Filter    json.RawMessage `json:"filter"`
			Sort      json.RawMessage `json:"sort,omitempty"`
			Limit     int             `json:"limit"`
			MaxTimeMS int             `json:"maxTimeMS"`
			Db        string          `json:"$db"`
		} `json:"command"`
	} `json:"attr"`
}

func LogToQuery(log string) (string, error) {
	var line LogLine
	err := json.Unmarshal([]byte(log), &line)
	if err != nil {
		return "", err
	}
	if !(line.C == "COMMAND" && line.Attr.Type == "command") {
		fmt.Println(line.C, line.Attr.Type)
		return "", nil
	}
	command := line.Attr.Command
	if command.Find == "" {
		return "", nil
	}
	result := fmt.Sprintf(`db.%s.find(%s)`, command.Find, command.Filter)
	if command.Sort != nil {
		result = fmt.Sprintf(`%s.sort(%s)`, result, command.Sort)
	}
	if command.Limit != 0 {
		result = fmt.Sprintf(`%s.limit(%d)`, result, command.Limit)
	}
	if command.MaxTimeMS != 0 {
		result = fmt.Sprintf(`%s.maxTimeMS(%d)`, result, command.MaxTimeMS)
	}
	return result, nil
}
