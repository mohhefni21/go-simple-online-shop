package entity

import (
	"mohhefni/go-online-shop/apps/product/request"
	"mohhefni/go-online-shop/infra/errorpkg"
	"time"

	"github.com/google/uuid"
)

type ProductEntity struct {
	Id        int       `db:"id"`
	SKU       string    `db:"sku"`
	Name      string    `db:"name"`
	Stock     int16     `db:"stock"`
	Price     int       `db:"price"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type ProductPaginationEntity struct {
	Cursor int
	Size   int
}

func NewFromGetProductsRequest(req request.GetProductsRequestPayload) ProductPaginationEntity {
	req = req.DefaultValuePagination()
	return ProductPaginationEntity{
		Cursor: req.Cursor,
		Size:   req.Size,
	}
}

func NewFromAddProductRequest(req request.AddProductPayload) ProductEntity {
	return ProductEntity{
		SKU:       uuid.NewString(),
		Name:      req.Name,
		Stock:     req.Stock,
		Price:     req.Price,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (p *ProductEntity) Validate() (err error) {
	err = p.ValidateName()
	if err != nil {
		return
	}

	err = p.ValidateStock()
	if err != nil {
		return
	}

	err = p.ValidatePrice()
	if err != nil {
		return
	}

	return
}

func (p *ProductEntity) ValidateName() (err error) {
	if p.Name == "" {
		return errorpkg.ErrProductRequired
	}

	if len(p.Name) < 3 {
		return errorpkg.ErrProductInvalid
	}

	return
}

func (p *ProductEntity) ValidateStock() (err error) {
	if p.Stock <= 0 {
		return errorpkg.ErrStockInvalid
	}

	return
}

func (p *ProductEntity) ValidatePrice() (err error) {
	if p.Price <= 0 {
		return errorpkg.ErrPriceInvalid
	}

	return
}
