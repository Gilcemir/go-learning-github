package validator

import "bytes"

type ValidationError struct {
	Messages []string
}

func newValidationError() *ValidationError {
	return &ValidationError{
		Messages: make([]string, 5),
	}
}

func (e *ValidationError) Error() string {
	var buffer bytes.Buffer
	for _, err := range e.Messages {
		if err != "" {
			buffer.WriteString(err)
			buffer.WriteString("\n")
		}
	}
	return buffer.String()
}
