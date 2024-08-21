package handler

import (
	"mohhefni/go-online-shop/apps/transaction/request"
	"mohhefni/go-online-shop/apps/transaction/response"
	"mohhefni/go-online-shop/apps/transaction/usecase"
	"mohhefni/go-online-shop/infra/responsepkg"
	"mohhefni/go-online-shop/utility"
	"net/http"

	"github.com/labstack/echo/v4"
)

type handler struct {
	usecase usecase.Usecase
}

func NewHandler(ucs usecase.Usecase) handler {
	return handler{
		usecase: ucs,
	}
}

func (h *handler) PostTransactionHandler(c echo.Context) error {
	req := request.AddTransactionPayload{}
	publicId := c.Get("public_id").(string)

	err := c.Bind(&req)
	if err != nil {
		utility.MakeLogEntry(nil).Warning(err)
		return responsepkg.NewResponse(
			responsepkg.WithStatus(err),
		).Send(c)
	}

	req.UserPublicId = publicId

	err = h.usecase.CreateTransaction(c.Request().Context(), req)
	if err != nil {
		utility.MakeLogEntry(nil).Warning(err)
		return responsepkg.NewResponse(
			responsepkg.WithStatus(err),
		).Send(c)
	}

	return responsepkg.NewResponse(
		responsepkg.WithHttpCode(http.StatusCreated),
	).Send(c)
}

func (h *handler) GetTransactionByUserHandler(c echo.Context) error {
	publicId := c.Get("public_id").(string)

	transactions, err := h.usecase.GetTransactionsHistory(c.Request().Context(), publicId)
	if err != nil {
		utility.MakeLogEntry(nil).Warning(err)
		return responsepkg.NewResponse(
			responsepkg.WithStatus(err),
		).Send(c)
	}

	transactionsList := response.NewListTransactionHistroyResponse(transactions)

	return responsepkg.NewResponse(
		responsepkg.WithHttpCode(http.StatusOK),
		responsepkg.WithData(map[string]interface{}{
			"transactions": transactionsList,
		}),
	).Send(c)
}
