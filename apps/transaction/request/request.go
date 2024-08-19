package request

type AddTransactionPayload struct {
	ProduckSku   string `json:"product_sku"`
	Amount       uint16 `json:"amount"`
	UserPublicId string `json:"-"`
}
