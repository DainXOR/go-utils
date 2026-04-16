package errorx

import "strings"

type ComposedError struct {
	Errors []error `json:"errors"`
}

func (m *ComposedError) Error() string {
	var msg strings.Builder
	msg.WriteString("Errors: ")
	for _, err := range m.Errors {
		msg.WriteString(err.Error() + " > ")
	}

	return msg.String()
}

func (m *ComposedError) Add(err error) {
	m.Errors = append(m.Errors, err)
}
