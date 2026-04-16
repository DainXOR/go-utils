package httpx

import (
	"errors"
	"fmt"
	"strings"
)

type HttpError struct {
	Code   HttpCode `json:"code"`
	Err    error    `json:"error"`
	Detail string   `json:"detail"`
}

func Error(code HttpCode, information ...string) HttpError {
	message := ""
	var detail strings.Builder

	if len(information) > 0 {
		message = information[0]

		for _, info := range information[1:] {
			detail.WriteString(info + " ")
		}
	}

	return HttpError{
		Code:   code,
		Err:    errors.New(message),
		Detail: detail.String(),
	}
}

func (m *HttpError) Error() string {
	return fmt.Sprintf("%s (%d): %s - %s", Http.Name(m.Code), m.Code, m.Err, m.Detail)
}

func (m *HttpError) Msg() string {
	return m.Err.Error()
}

func (m *HttpError) Details() string {
	return m.Detail
}

func ErrorNotFound(information ...string) HttpError {
	return Error(Http.C400().NotFound(), information...)
}

func ErrorInternal(information ...string) HttpError {
	return Error(Http.C500().InternalServerError(), information...)
}
