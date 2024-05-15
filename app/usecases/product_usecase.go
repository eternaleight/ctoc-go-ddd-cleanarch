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

func CreateProduct(store stores.ProductStoreInterface, input ProductInput) (*models.Product, error) {
	product := models.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
	}

	err := store.CreateProduct(&product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func ListProducts(store stores.ProductStoreInterface) ([]models.Product, error) {
	products, err := store.ListProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func GetProductByID(store stores.ProductStoreInterface, id uint) (*models.Product, error) {
	product, err := store.GetProductByID(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func UpdateProduct(store stores.ProductStoreInterface, id uint, input ProductInput) (*models.Product, error) {
	product := models.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
	}

	err := store.UpdateProduct(id, &product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func DeleteProduct(store stores.ProductStoreInterface, id uint) error {
	err := store.DeleteProduct(id)
	if err != nil {
		return err
	}
	return nil
}
