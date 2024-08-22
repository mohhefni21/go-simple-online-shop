package request

type AddProductPayload struct {
	Name  string `json:"name" example:"sabun"`
	Stock int16  `json:"stock" example:"25"`
	Price int    `json:"price" example:"2500"`
}

type GetProductsRequestPayload struct {
	Cursor int `query:"cursor" json:"cursor"`
	Size   int `query:"size" json:"size"`
}

func (g *GetProductsRequestPayload) DefaultValuePagination() GetProductsRequestPayload {
	if g.Cursor < 0 {
		g.Cursor = 0
	}

	if g.Size <= 0 {
		g.Cursor = 10
	}

	return *g
}
