package handler

import (
	"mohhefni/go-online-shop/apps/product/request"
	"mohhefni/go-online-shop/apps/product/response"
	"mohhefni/go-online-shop/apps/product/usecase"
	"mohhefni/go-online-shop/infra/errorpkg"
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

func (h *handler) PostProductHandler(c echo.Context) error {
	req := request.AddProductPayload{}

	err := c.Bind(&req)
	if err != nil {
		utility.MakeLogEntry(nil).Warning(err)
		return responsepkg.NewResponse(
			responsepkg.WithStatus(err),
		).Send(c)
	}

	productSku, err := h.ucs.CreateProduct(c.Request().Context(), req)
	if err != nil {
		utility.MakeLogEntry(nil).Warning(err)
		return responsepkg.NewResponse(
			responsepkg.WithStatus(err),
		).Send(c)
	}

	return responsepkg.NewResponse(
		responsepkg.WithHttpCode(http.StatusCreated),
		responsepkg.WithData(map[string]interface{}{
			"productSku": productSku,
		}),
	).Send(c)
}

func (h *handler) GetAllProductsHandler(c echo.Context) error {
	req := request.GetProductsRequestPayload{}

	err := c.Bind(&req)
	if err != nil {
		utility.MakeLogEntry(nil).Warning(err)
		return responsepkg.NewResponse(
			responsepkg.WithStatus(err),
		).Send(c)
	}

	products, err := h.ucs.GetProducts(c.Request().Context(), req)
	if err != nil {
		utility.MakeLogEntry(nil).Warning(err)
		return responsepkg.NewResponse(
			responsepkg.WithStatus(err),
		).Send(c)
	}

	productResponse := response.NewListALlProductResponse(products)

	return responsepkg.NewResponse(
		responsepkg.WithHttpCode(http.StatusOK),
		responsepkg.WithData(map[string]interface{}{
			"products": productResponse,
		}),
		responsepkg.WithQuery(req),
	).Send(c)

}

func (h *handler) GetDetailProductHandler(c echo.Context) error {
	params := c.Param("sku")
	if params == "" {
		err := errorpkg.ErrorNotFound
		utility.MakeLogEntry(nil).Warning(err)
		return responsepkg.NewResponse(
			responsepkg.WithStatus(err),
		).Send(c)
	}

	products, err := h.ucs.GetDetailProduct(c.Request().Context(), params)
	if err != nil {
		utility.MakeLogEntry(nil).Warning(err)
		return responsepkg.NewResponse(
			responsepkg.WithStatus(err),
		).Send(c)
	}

	productResponse := response.GetDetailProductResponse(products)

	return responsepkg.NewResponse(
		responsepkg.WithHttpCode(http.StatusOK),
		responsepkg.WithData(map[string]interface{}{
			"products": productResponse,
		}),
	).Send(c)

}
