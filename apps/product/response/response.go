package response

import (
	"mohhefni/go-online-shop/apps/product/entity"
	"time"
)

type GetALlProductsResponse struct {
	Id    int    `json:"id"`
	SKU   string `json:"sku"`
	Name  string `json:"name"`
	Stock int16  `json:"stock"`
	Price int    `json:"price"`
}

func NewListALlProductResponse(products []entity.ProductEntity) []GetALlProductsResponse {
	var productList = []GetALlProductsResponse{}

	for _, product := range products {
		productList = append(productList, GetALlProductsResponse{
			Id:    product.Id,
			SKU:   product.SKU,
			Name:  product.Name,
			Stock: product.Stock,
			Price: product.Price,
		})
	}

	return productList
}

type GetDetailProductResponse struct {
	Id        int       `json:"id"`
	SKU       string    `json:"sku"`
	Name      string    `json:"name"`
	Stock     int16     `json:"stock"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewDetailProductResponse(products entity.ProductEntity) GetDetailProductResponse {
	return GetDetailProductResponse{
		Id:        products.Id,
		SKU:       products.SKU,
		Name:      products.Name,
		Stock:     products.Stock,
		Price:     products.Price,
		CreatedAt: products.CreatedAt,
		UpdatedAt: products.UpdatedAt,
	}
}
