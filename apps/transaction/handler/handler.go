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

// PostTransactionHandler godoc
// @Router       /transactions/checkout [post]
// @Summary      Checkout transation
// @Description  Process a checkout transaction, reducing stock levels and creating an order record.
// @Tags         transactions
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Bearer token"
// @Param        request body request.AddTransactionPayload true "Add transaction details"
// @Success      201 "Created - Transaction added successfully"
// @Failure      400 "Bad request - Invalid input"
// @Failure      401 "Unauthorized - Invalid or missing token"
// @Failure      403 "Forbidden access - access not allowed"
// @Failure      404 "Not Found - Resource not found"
// @Failure      500 "Internal server error"
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

// GetTransactionByUserHandler godoc
// @Router       /transactions/history [get]
// @Summary      History transation
// @Description  Retrieve a list of history transactions. Returns a collection of transactions attributes.
// @Tags         transactions
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Bearer token"
// @Success      200 "Ok - Successfully retrieved transaction histories"
// @Failure      500 "Internal server error"
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
