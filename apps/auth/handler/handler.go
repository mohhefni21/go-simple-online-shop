package handler

import (
	"mohhefni/go-online-shop/apps/auth/request"
	"mohhefni/go-online-shop/apps/auth/usecase"
	infraecho "mohhefni/go-online-shop/infra/echo"
	"mohhefni/go-online-shop/infra/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type handler struct {
	ucs usecase.Usecase
}

func NewHandler(usecase usecase.Usecase) *handler {
	return &handler{
		ucs: usecase,
	}
}

func (h *handler) Register(c echo.Context) error {
	req := request.RegisterRequestPayload{}

	err := c.Bind(&req)
	if err != nil {
		myError := response.ErrorBadRequest
		return infraecho.NewResponse(
			infraecho.WithMesssage(err.Error()),
			infraecho.WithError(myError),
			infraecho.WithHttpCode(myError.HttpCode),
		).Send(c)
	}

	idUser, err := h.ucs.RegisterUser(c.Request().Context(), req)
	if err != nil {
		myError, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myError = response.ErrorGeneral
		}
		return infraecho.NewResponse(
			infraecho.WithMesssage(err.Error()),
			infraecho.WithError(myError),
			infraecho.WithHttpCode(response.ErrorBadRequest.HttpCode),
		).Send(c)
	}

	return infraecho.NewResponse(
		infraecho.WithHttpCode(http.StatusCreated),
		infraecho.WithData(map[string]interface{}{
			"userId": idUser,
		}),
	).Send(c)
}

func (h *handler) Login(c echo.Context) error {
	req := request.LoginRequestPayload{}

	err := c.Bind(&req)
	if err != nil {
		error := response.ErrorBadRequest
		return infraecho.NewResponse(
			infraecho.WithMesssage(err.Error()),
			infraecho.WithError(error),
			infraecho.WithHttpCode(error.HttpCode),
		).Send(c)
	}

	token, err := h.ucs.LoginUser(c.Request().Context(), req)
	if err != nil {
		error, ok := response.ErrorMapping[err.Error()]
		if !ok {
			error = response.ErrorGeneral
		}
		return infraecho.NewResponse(
			infraecho.WithMesssage(err.Error()),
			infraecho.WithError(error),
			infraecho.WithHttpCode(response.ErrorBadRequest.HttpCode),
		).Send(c)
	}

	return infraecho.NewResponse(
		infraecho.WithHttpCode(http.StatusCreated),
		infraecho.WithData(map[string]interface{}{
			"accessToken": token,
		}),
	).Send(c)
}