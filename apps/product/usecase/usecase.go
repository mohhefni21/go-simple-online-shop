package usecase

import (
	"context"
	"mohhefni/go-online-shop/apps/product/entity"
	"mohhefni/go-online-shop/apps/product/repository"
	"mohhefni/go-online-shop/apps/product/request"
	"mohhefni/go-online-shop/infra/errorpkg"
)

type Usecase interface {
	CreateProduct(ctx context.Context, req request.AddProductPayload) (skuProduct string, err error)
	GetProducts(ctx context.Context, req request.GetProductsRequestPayload) (products []entity.ProductEntity, err error)
	GetDetailProduct(ctx context.Context, sku string) (product entity.ProductEntity, err error)
}

type usecase struct {
	repo repository.Repository
}

func NewUsecase(repository repository.Repository) Usecase {
	return &usecase{
		repo: repository,
	}
}

func (u *usecase) CreateProduct(ctx context.Context, req request.AddProductPayload) (skuProduct string, err error) {
	productEntity := entity.NewFromAddProductRequest(req)
	err = productEntity.Validate()
	if err != nil {
		return
	}

	skuProduct, err = u.repo.AddProduct(ctx, productEntity)
	if err != nil {
		return
	}

	return
}

func (u *usecase) GetProducts(ctx context.Context, req request.GetProductsRequestPayload) (products []entity.ProductEntity, err error) {
	pagination := entity.NewFromGetProductsRequest(req)

	products, err = u.repo.GetAllProducts(ctx, pagination)
	if err != nil {
		if err == errorpkg.ErrNotFound {
			return []entity.ProductEntity{}, nil
		}

		return []entity.ProductEntity{}, err
	}

	if len(products) == 0 {
		return []entity.ProductEntity{}, nil
	}

	return
}

func (u *usecase) GetDetailProduct(ctx context.Context, sku string) (product entity.ProductEntity, err error) {
	product, err = u.repo.GetDetailProductBySku(ctx, sku)
	if err != nil {
		return
	}

	return
}
