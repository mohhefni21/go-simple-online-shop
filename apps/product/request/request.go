package request

type AddProductPayload struct {
	Name  string `json:"name"`
	Stock int16  `json:"stock"`
	Price int    `json:"price"`
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
