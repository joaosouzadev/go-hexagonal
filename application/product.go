package application

import (
	"errors"
	uuid "github.com/satori/go.uuid"
)

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetId() string
	GetUuid() string
	GetName() string
	GetPrice() int
	IsActive() bool
	GetOnStock() bool
}

type ProductServiceInterface interface {
	Get(id string) (ProductInterface, error)
	Create(dto ProductInputDto) (ProductInterface, error)
	Enable(product ProductInterface) (ProductInterface, error)
	Disable(product ProductInterface) (ProductInterface, error)
}

type ProductReader interface {
	Get(id string) (ProductInterface, error)
}

type ProductWriter interface {
	Save(product ProductInterface) (ProductInterface, error)
}

type ProductPersistenceInterface interface {
	ProductReader
	ProductWriter
}

type Product struct {
	ID      string `json:"id,omitempty"`
	Uuid    string `json:"uuid"`
	Name    string `json:"name"`
	Price   int    `json:"price"`
	Active  bool   `json:"active"`
	OnStock bool   `json:"on_stock"`
}

func NewProduct(name string, price int, active bool, onStock bool) *Product {
	return &Product{
		Uuid:    uuid.NewV4().String(),
		Name:    name,
		Price:   price,
		Active:  active,
		OnStock: onStock,
	}
}

func (p *Product) IsValid() (bool, error) {
	if len(p.Name) < 2 {
		return false, errors.New("product name must be at least 2 characters long")
	}

	if p.Price < 0 {
		return false, errors.New("product price cannot be less than zero")
	}

	return true, nil
}

func (p *Product) Enable() error {
	if p.Price > 0 {
		p.Active = true
		return nil
	}

	return errors.New("price must be greater than zero to enable a product")
}

func (p *Product) Disable() error {
	if p.Price == 0 {
		p.Active = false
		return nil
	}

	return errors.New("price must be zero to disable a product")
}

func (p *Product) GetId() string {
	return p.ID
}

func (p *Product) GetUuid() string {
	return p.Uuid
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetPrice() int {
	return p.Price
}

func (p *Product) IsActive() bool {
	return p.Active
}

func (p *Product) GetOnStock() bool {
	return p.OnStock
}
