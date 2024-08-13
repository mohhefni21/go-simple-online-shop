package handler

import (
	"mohhefni/go-online-shop/apps/auth/request"
	"mohhefni/go-online-shop/apps/auth/usecase"
	"mohhefni/go-online-shop/infra/errorpkg"
	"mohhefni/go-online-shop/infra/responsepkg"
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
		myError := errorpkg.ErrorBadRequest
		return responsepkg.NewResponse(
			responsepkg.WithMessage(err.Error()),
			responsepkg.WithError(myError),
			responsepkg.WithHttpCode(myError.HttpCode),
		).Send(c)
	}

	idUser, err := h.ucs.RegisterUser(c.Request().Context(), req)
	if err != nil {
		myError, ok := errorpkg.ErrorMapping[err.Error()]
		if !ok {
			myError = errorpkg.ErrorGeneral
		}
		return responsepkg.NewResponse(
			responsepkg.WithMessage(err.Error()),
			responsepkg.WithError(myError),
			responsepkg.WithHttpCode(errorpkg.ErrorBadRequest.HttpCode),
		).Send(c)
	}

	return responsepkg.NewResponse(
		responsepkg.WithHttpCode(http.StatusCreated),
		responsepkg.WithData(map[string]interface{}{
			"userId": idUser,
		}),
	).Send(c)
}

func (h *handler) Login(c echo.Context) error {
	req := request.LoginRequestPayload{}

	err := c.Bind(&req)
	if err != nil {
		error := errorpkg.ErrorBadRequest
		return responsepkg.NewResponse(
			responsepkg.WithMessage(err.Error()),
			responsepkg.WithError(error),
			responsepkg.WithHttpCode(error.HttpCode),
		).Send(c)
	}

	token, err := h.ucs.LoginUser(c.Request().Context(), req)
	if err != nil {
		error, ok := errorpkg.ErrorMapping[err.Error()]
		if !ok {
			error = errorpkg.ErrorGeneral
		}
		return responsepkg.NewResponse(
			responsepkg.WithMessage(err.Error()),
			responsepkg.WithError(error),
			responsepkg.WithHttpCode(errorpkg.ErrorBadRequest.HttpCode),
		).Send(c)
	}

	return responsepkg.NewResponse(
		responsepkg.WithHttpCode(http.StatusCreated),
		responsepkg.WithData(map[string]interface{}{
			"accessToken": token,
		}),
	).Send(c)
}
