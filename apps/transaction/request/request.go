package request

type AddTransactionPayload struct {
	ProduckSku string `json:"product_sku"`
	Amount     uint8  `json:"amount"`
	Email      string `json:"-"`
}
