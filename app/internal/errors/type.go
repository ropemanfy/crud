package errors

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

const noneValue = "None"

var errorsMap = make(map[string]Error)

type Error struct {
	Data         interface{}       `json:"data"`
	IsTrue       bool              `json:"error"`
	ErrorText    string            `json:"errorText"`
	CustomErrors map[string]string `json:"additionalErrors"`
	Cause        interface{}       `json:"-"`
	Status       int               `json:"-"`
}

func new(msg, trKey string) Error {
	errorsMap[trKey] = Error{ErrorText: msg, IsTrue: true, Status: fasthttp.StatusInternalServerError}
	return errorsMap[trKey]
}

func (e Error) WithData(data interface{}) Error {
	e.Data = data
	return e
}

func (e Error) Error() (errStr string) {

	if e.Data != nil {
		errStr = fmt.Sprintf(": %v", e.Data)
	}
	if e.Cause != "" {
		errStr = fmt.Sprintf("%s cause: %s", errStr, e.Cause)
	}
	return e.ErrorText + errStr
}

func (e Error) SetCause(format string, a ...interface{}) Error {
	e.Cause = fmt.Sprintf(format, a...)
	return e
}

func (e Error) AddCause(args ...string) Error {

	result := make(map[string]interface{})
	for i := 0; i < len(args); i += 2 {
		strKey := args[i]
		result[strKey] = noneValue
		if i+1 < len(args) {
			result[strKey] = args[i+1]
		}
	}
	e.Cause = result
	return e
}

func (e Error) SetStatus(code int) Error {
	e.Status = code
	return e
}

func (e Error) Code() int {
	return e.Status
}

func (e Error) Send(ctx *fiber.Ctx) (err error) {
	ctx.Status(e.Status)
	ctx.Response().Header.SetContentType("application/json")
	err = json.NewEncoder(ctx).Encode(e)
	return
}

func Map() (errors map[string]Error) {

	errors = make(map[string]Error)
	for k, v := range errorsMap {
		errors[k] = v
	}
	return
}
