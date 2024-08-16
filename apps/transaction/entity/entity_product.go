package entity

type ProductJsonEntity struct {
	Id    int    `db:"id" json:"id"`
	SKU   string `db:"sku" json:"sku"`
	Name  string `db:"name" json:"name"`
	Price int    `db:"price" json:"price"`
}
