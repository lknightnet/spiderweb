package log

import "fmt"

type Message struct {
	Level            string                 `json:"level"`
	Text             string                 `json:"text"`
	MicroServiceName string                 `json:"ms_name"`
	Other            map[string]interface{} `json:"other"`
}

func NewOther(input ...interface{}) (map[string]interface{}, error) {
	res := make(map[string]interface{})
	if len(input)%2 != 0 {
		return nil, fmt.Errorf("missing pair")
	}

	for i := 0; i < len(input); i += 2 {
		key, ok := input[i].(string)
		if !ok {
			return nil, fmt.Errorf("unable to convert key to string: %v", input[i])
		}

		value := input[i+1]
		res[key] = value
	}
	return res, nil
}

func NewMessage(lvl, text, msName string, o map[string]interface{}) *Message {
	return &Message{
		Level:            lvl,
		Text:             text,
		MicroServiceName: msName,
		Other:            o,
	}
}
