package handler

import (
	"mohhefni/go-online-shop/apps/auth/request"
	"mohhefni/go-online-shop/apps/auth/usecase"
	"mohhefni/go-online-shop/infra/responsepkg"
	"mohhefni/go-online-shop/utility"
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

// PostRegisterHandler godoc
// @Router       /auth/register [post]
// @Summary Register new user
// @Description Register with new user with the provided details
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body request.RegisterRequestPayload true "User registration details"
// @Success      201 "Created - User registered successfully"
// @Failure      400 "Bad request - Invalid input"
// @Failure      409 "Conflict - user already exists"
// @Failure      500 "Internal server error"
func (h *handler) PostRegisterHandler(c echo.Context) error {
	req := request.RegisterRequestPayload{}

	err := c.Bind(&req)
	if err != nil {
		utility.MakeLogEntry(nil).Warning(err)
		return responsepkg.NewResponse(
			responsepkg.WithStatus(err),
		).Send(c)
	}

	idUser, err := h.ucs.RegisterUser(c.Request().Context(), req)
	if err != nil {
		utility.MakeLogEntry(nil).Warning(err)
		return responsepkg.NewResponse(
			responsepkg.WithStatus(err),
		).Send(c)
	}

	return responsepkg.NewResponse(
		responsepkg.WithHttpCode(http.StatusCreated),
		responsepkg.WithData(map[string]interface{}{
			"userId": idUser,
		}),
	).Send(c)
}

// PostLoginHandler godoc
// @Router       /auth/login [post]
// @Summary      Login user
// @Description  Authenticate a user and return a JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body request.LoginRequestPayload true "User login details"
// @Success      201 "Created - User logged successfully"
// @Failure      400 "Bad request - Invalid input"
// @Failure      401 "Unauthorized - Invalid or missing token"
// @Failure      404 "Not Found - Resource not found"
// @Failure      500 "Internal server error"
func (h *handler) PostLoginHandler(c echo.Context) error {
	req := request.LoginRequestPayload{}

	err := c.Bind(&req)
	if err != nil {
		utility.MakeLogEntry(nil).Warning(err)
		return responsepkg.NewResponse(
			responsepkg.WithStatus(err),
		).Send(c)
	}

	token, err := h.ucs.LoginUser(c.Request().Context(), req)
	if err != nil {
		utility.MakeLogEntry(nil).Warning(err)
		return responsepkg.NewResponse(
			responsepkg.WithStatus(err),
		).Send(c)
	}

	return responsepkg.NewResponse(
		responsepkg.WithHttpCode(http.StatusCreated),
		responsepkg.WithData(map[string]interface{}{
			"accessToken": token,
		}),
	).Send(c)
}
