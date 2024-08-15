package responsepkg

import (
	"mohhefni/go-online-shop/infra/errorpkg"

	"github.com/labstack/echo/v4"
)

type Response struct {
	HttpCode int         `json:"-"`
	Status   string      `json:"status"`
	Message  string      `json:"message,omitempty"`
	Data     interface{} `json:"data,omitempty"`
	Query    interface{} `json:"query,omitempty"`
	Error    string      `json:"error,omitempty"`
}

func NewResponse(params ...func(*Response) *Response) Response {
	var resp = Response{
		Status: "success",
	}
	for _, param := range params {
		param(&resp)
	}

	return resp
}

func WithStatus(err error) func(*Response) *Response {
	return func(r *Response) *Response {
		var receivedError error
		// Cek error di ErrorMapping
		receivedError, ok := errorpkg.ErrorMapping[err.Error()]
		if !ok {
			// Jika tidak ada gunakan error general
			receivedError = errorpkg.ErrorGeneral
		}

		// Casting ke tipe Error
		myError, ok := receivedError.(errorpkg.Error)
		if !ok {
			// Jika casting gagal gunakan error general lagi
			myError = errorpkg.ErrorGeneral
		}

		r.Status = "fail"
		r.HttpCode = myError.HttpCode
		r.Message = myError.Messsage

		// Jika error adalah ErrorGeneral kirim error asli untuk debugging
		if myError == errorpkg.ErrorGeneral {
			r.Status = "error"
			r.Error = err.Error()
		}

		return r
	}
}

func WithHttpCode(httpCode int) func(*Response) *Response {
	return func(r *Response) *Response {
		r.HttpCode = httpCode
		return r
	}
}

func WithData(data interface{}) func(*Response) *Response {
	return func(r *Response) *Response {
		r.Data = data
		return r
	}
}

func WithQuery(query interface{}) func(*Response) *Response {
	return func(r *Response) *Response {
		r.Query = query
		return r
	}
}

func (r Response) Send(ctx echo.Context) error {
	return ctx.JSON(r.HttpCode, r)
}
