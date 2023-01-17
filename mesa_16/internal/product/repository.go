package product

import (
	"errors"
	"fmt"

	"github.com/bootcamp-go/Consignas-Go-Web.git/internal/domain"
)
// here we define the interface for the repository and the service layers of the product domain model 
type Repository interface {
	GetAll() []domain.Product
	GetByID(id int) (domain.Product, error)
	SearchPriceGt(price float64) []domain.Product
	Create(p domain.Product) (domain.Product, error)
	Put(p domain.Product) (domain.Product, error)
	PatchName(id int, name string) (domain.Product, error)
	Delete(id int) error
}

type repository struct {
	list []domain.Product
}

// NewRepository crea un nuevo repositorio
func NewRepository(list []domain.Product) Repository {
	return &repository{list}
}

func (r *repository) PatchName(id int, name string) (domain.Product, error) {
	for i, domainVar := range r.list {
		if domainVar.Id == id {
			product := r.list[i]
			product.Name = name
			return product, nil
		}
	}

	return domain.Product{}, fmt.Errorf("Product with id %d not found.", id)
}

func (r *repository) Put(p domain.Product) (domain.Product, error) {
	for i, domainVar := range r.list {
		if domainVar.Id == p.Id {
			r.list[i] = p
			return p, nil
		}
	}

	if !r.validateCodeValue(p.CodeValue) {
		return domain.Product{}, errors.New("code value already exists")
	}

	p.Id = len(r.list) + 1
	r.list = append(r.list, p)
	return p, nil
}

// GetAll devuelve todos los productos
func (r *repository) GetAll() []domain.Product {
	return r.list
}

// GetByID busca un producto por su id
func (r *repository) GetByID(id int) (domain.Product, error) {
	for _, product := range r.list {
		if product.Id == id {
			return product, nil
		}
	}
	return domain.Product{}, errors.New("product not found")

}

// SearchPriceGt busca productos por precio mayor o igual que el precio dado
func (r *repository) SearchPriceGt(price float64) []domain.Product {
	var products []domain.Product
	for _, product := range r.list {
		if product.Price > price {
			products = append(products, product)
		}
	}
	return products
}

// Create agrega un nuevo producto
func (r *repository) Create(p domain.Product) (domain.Product, error) {
	if !r.validateCodeValue(p.CodeValue) {
		return domain.Product{}, errors.New("code value already exists")
	}
	p.Id = len(r.list) + 1 //
	r.list = append(r.list, p)
	return p, nil
}

// validateCodeValue valida que el codigo no exista en la lista de productos
func (r *repository) validateCodeValue(codeValue string) bool {
	for _, product := range r.list {
		if product.CodeValue == codeValue {
			return false
		}
	}
	return true
}
func (r *repository) Delete(id int) error {

	for i, product := range r.list {
		if product.Id == id {
			r.list = append(r.list[:i], r.list[i+1:]...)
			return nil
		}
	}
		return errors.New("code value don't exists")
	

}
