package responsepkg

import (
	"mohhefni/go-online-shop/infra/errorpkg"

	"github.com/labstack/echo/v4"
)

type Response struct {
	HttpCode  int         `json:"-"`
	Success   bool        `json:"success"`
	Message   string      `json:"message,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	Error     string      `json:"error,omitempty"`
	ErrorCode string      `json:"error_code,omitempty"`
}

func NewResponse(params ...func(*Response) *Response) Response {
	var resp = Response{
		Success: true,
	}
	for _, param := range params {
		param(&resp)
	}

	return resp
}

func WithHttpCode(httpCode int) func(*Response) *Response {
	return func(r *Response) *Response {
		r.HttpCode = httpCode
		return r
	}
}

func WithSuccess(success bool) func(*Response) *Response {
	return func(r *Response) *Response {
		r.Success = success
		return r
	}
}

func WithMessage(message string) func(*Response) *Response {
	return func(r *Response) *Response {
		r.Message = message
		return r
	}
}

func WithData(data interface{}) func(*Response) *Response {
	return func(r *Response) *Response {
		r.Data = data
		return r
	}
}

func WithError(err error) func(*Response) *Response {
	return func(r *Response) *Response {
		r.Success = false

		myError, ok := err.(errorpkg.Error)
		if !ok {
			myError = errorpkg.ErrorGeneral
		}

		r.Error = myError.Messsage
		r.ErrorCode = myError.Code

		return r
	}
}

func (r Response) Send(ctx echo.Context) error {
	return ctx.JSON(r.HttpCode, r)
}
