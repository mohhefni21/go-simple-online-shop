package request

type AddTransactionPayload struct {
	ProduckSku   string `json:"product_sku" example:"3f369638-de78-4c6e-8e99-2f5507d346c7"`
	Amount       uint16 `json:"amount" example:"7"`
	UserPublicId string `json:"-"`
}
