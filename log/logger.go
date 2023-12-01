package log

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Logger struct {
	addr string
}

func NewLogger(addr string) *Logger {
	return &Logger{
		addr: addr,
	}
}

func (l *Logger) Print(msg *Message) error {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(msg)
	if err != nil {
		return fmt.Errorf("info message is not encoding with error: %w", err)
	}

	resp, err := http.Post(l.addr, "application/json", &buf)
	if err != nil {
		return fmt.Errorf("info message is not sent with an error: %w", err)
	}

	validError(resp.StatusCode)
	return nil
}

func validError(status int) {
	if status != http.StatusOK {
		panic(errors.New("you message not delivered"))
	}
}
