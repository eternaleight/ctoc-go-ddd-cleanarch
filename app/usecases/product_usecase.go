package usecases

import (
	"github.com/eternaleight/go-backend/domain/models"
	"github.com/eternaleight/go-backend/infra/stores"
)

type ProductInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

type ProductUsecasesInterface interface {
	CreateProduct(input ProductInput) (*models.Product, error)
	ListProducts() ([]models.Product, error)
	GetProductByID(id uint) (*models.Product, error)
	UpdateProduct(id uint, input ProductInput) (*models.Product, error)
	DeleteProduct(id uint) error
}

type ProductUsecases struct {
	ProductStore stores.ProductStoreInterface
}

// ProductUsecasesの新しいインスタンスを初期化
func NewProductUsecases(productStore stores.ProductStoreInterface) *ProductUsecases {
	return &ProductUsecases{
		ProductStore: productStore,
	}
}

func (u *ProductUsecases) CreateProduct(input ProductInput) (*models.Product, error) {
	product := models.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
	}

	err := u.ProductStore.CreateProduct(&product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (u *ProductUsecases) ListProducts() ([]models.Product, error) {
	products, err := u.ProductStore.ListProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (u *ProductUsecases) GetProductByID(id uint) (*models.Product, error) {
	product, err := u.ProductStore.GetProductByID(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (u *ProductUsecases) UpdateProduct(id uint, input ProductInput) (*models.Product, error) {
	product := models.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
	}

	err := u.ProductStore.UpdateProduct(id, &product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (u *ProductUsecases) DeleteProduct(id uint) error {
	err := u.ProductStore.DeleteProduct(id)
	if err != nil {
		return err
	}
	return nil
}
