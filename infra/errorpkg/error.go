package errorpkg

import (
	"errors"
	"net/http"
)

var (
	ErrNotFound        = errors.New("not found")
	ErrUnauthorized    = errors.New("unauthorized")
	ErrForbiddenAccess = errors.New("forbidden access")
)

var (
	// Auth
	ErrEmailRequired         = errors.New("email is required")
	ErrEmailInvalid          = errors.New("email is invalid")
	ErrPasswordRequired      = errors.New("password is required")
	ErrPasswordInvalidLength = errors.New("password must have minimum character")
	ErrEmailAlreadyUsed      = errors.New("email already used")
	ErrPasswordNotMatch      = errors.New("password not match")

	// Product
	ErrProductRequired = errors.New("product is required")
	ErrProductInvalid  = errors.New("product must have minimum 3 character")
	ErrStockInvalid    = errors.New("stock must be greater than 0")
	ErrPriceInvalid    = errors.New("price must be greater than 0")

	// Transaction
	ErrAmountInvalid          = errors.New("invalid amount")
	ErrAmountGreaterThanStock = errors.New("amount greater than stock")
)

type Error struct {
	Messsage string
	HttpCode int
}

func (e Error) Error() string {
	return e.Messsage
}

func NewError(message string, httpCode int) Error {
	return Error{
		Messsage: message,
		HttpCode: httpCode,
	}
}

var (
	ErrorGeneral         = NewError("internal server error", http.StatusInternalServerError)
	ErrorBadRequest      = NewError("bad request", http.StatusBadRequest)
	ErrorNotFound        = NewError(ErrNotFound.Error(), http.StatusNotFound)
	ErrorUnauthorized    = NewError(ErrUnauthorized.Error(), http.StatusUnauthorized)
	ErrorForbiddenAccess = NewError(ErrForbiddenAccess.Error(), http.StatusForbidden)
)

var (
	// Auth
	ErrorEmailRequired         = NewError(ErrEmailRequired.Error(), http.StatusBadRequest)
	ErrorEmailInvalid          = NewError(ErrEmailInvalid.Error(), http.StatusBadRequest)
	ErrorPasswordRequired      = NewError(ErrPasswordRequired.Error(), http.StatusBadRequest)
	ErrorPasswordInvalidLength = NewError(ErrPasswordInvalidLength.Error(), http.StatusBadRequest)

	ErrorPasswordNotMatch = NewError(ErrPasswordNotMatch.Error(), http.StatusUnauthorized)
	ErrorEmailAlreadyUsed = NewError(ErrEmailAlreadyUsed.Error(), http.StatusConflict)

	// product
	ErrorProductRequired = NewError(ErrProductRequired.Error(), http.StatusBadRequest)
	ErrorProductInvalid  = NewError(ErrProductInvalid.Error(), http.StatusBadRequest)
	ErrorStockInvalid    = NewError(ErrStockInvalid.Error(), http.StatusBadRequest)
	ErrorPriceInvalid    = NewError(ErrPriceInvalid.Error(), http.StatusBadRequest)

	// Transaction
	ErrorAmountInvalid          = NewError(ErrAmountInvalid.Error(), http.StatusBadRequest)
	ErrorAmountGreaterThanStock = NewError(ErrAmountInvalid.Error(), http.StatusBadRequest)
)

var (
	ErrorMapping = map[string]Error{
		ErrEmailRequired.Error():          ErrorEmailRequired,
		ErrEmailInvalid.Error():           ErrorEmailInvalid,
		ErrPasswordRequired.Error():       ErrorPasswordRequired,
		ErrEmailAlreadyUsed.Error():       ErrorEmailAlreadyUsed,
		ErrPasswordInvalidLength.Error():  ErrorPasswordInvalidLength,
		ErrPasswordNotMatch.Error():       ErrorPasswordNotMatch,
		ErrProductRequired.Error():        ErrorProductRequired,
		ErrProductInvalid.Error():         ErrorProductInvalid,
		ErrPriceInvalid.Error():           ErrorPriceInvalid,
		ErrStockInvalid.Error():           ErrorStockInvalid,
		ErrNotFound.Error():               ErrorNotFound,
		ErrAmountInvalid.Error():          ErrorAmountInvalid,
		ErrAmountGreaterThanStock.Error(): ErrorAmountGreaterThanStock,
	}
)
