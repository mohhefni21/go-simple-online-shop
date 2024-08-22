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

// PostProductHandler godoc
// @Router       /products [post]
// @Summary      Add product
// @Description  Add products and return id products
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Bearer token"
// @Param        request body request.AddProductPayload true "Add product details"
// @Success      201 "Created - Product added successfully"
// @Failure      400 "Bad request - Invalid input"
// @Failure      401 "Unauthorized - Invalid or missing token"
// @Failure      403 "Forbidden access - access not allowed"
// @Failure      500 "Internal server error"
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

// GetAllProductsHandler godoc
// @Router       /products [get]
// @Summary      Get List of Products
// @Description  Retrieve a list of products with optional pagination. Returns a collection of product attributes.
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        cursor query int false "Cursor for pagination"
// @Param        size   query int false "Number of items per page"
// @Success      200 "Ok - Successfully retrieved all products"
// @Failure      500 "Internal server error"
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

// GetDetailProductHandler godoc
// @Router       /products/sku/{sku} [get]
// @Summary      Get Product Details
// @Description  Retrieve detailed information about a specific product using its SKU. Returns attributes of the product.
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        sku path string true "Sku products"
// @Success      200 "Ok - Successfully retrieved the product details"
// @Failure      404 "Not Found - Resource not found"
// @Failure      500 "Internal server error"
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
