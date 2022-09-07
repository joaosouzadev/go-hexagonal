package application

type ProductInputDto struct {
	Name    string `json:"name"`
	Price   int    `json:"price"`
	Active  bool   `json:"active"`
	OnStock bool   `json:"on_stock"`
}

type ProductOutputDto struct {
	Uuid    string `json:"uuid"`
	Name    string `json:"name"`
	Price   int    `json:"price"`
	Active  bool   `json:"active"`
	OnStock bool   `json:"on_stock"`
}

func (dto *ProductInputDto) HydrateFromInput(product *Product) (*Product, error) {
	product.Name = dto.Name
	product.Price = dto.Price
	product.Active = dto.Active
	product.OnStock = dto.OnStock

	_, err := product.IsValid()
	if err != nil {
		return &Product{}, nil
	}
	return product, nil
}

func (dto *ProductOutputDto) HydrateFromEntity(product ProductInterface) {
	dto.Uuid = product.GetUuid()
	dto.Name = product.GetName()
	dto.Price = product.GetPrice()
	dto.Active = product.IsActive()
	dto.OnStock = product.GetOnStock()
}
