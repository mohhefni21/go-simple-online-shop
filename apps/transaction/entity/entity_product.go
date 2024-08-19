package entity

import "mohhefni/go-online-shop/infra/errorpkg"

type ProductJsonEntity struct {
	Id    int    `db:"id" json:"id"`
	SKU   string `db:"sku" json:"sku"`
	Name  string `db:"name" json:"name"`
	Stock uint16 `db:"stock" json:"-"`
	Price int    `db:"price" json:"price"`
}

func (p *ProductJsonEntity) UpdateStockProduct(amount uint16) (err error) {
	if p.Stock < amount {
		return errorpkg.ErrAmountGreaterThanStock
	}

	p.Stock = p.Stock - amount

	return
}
