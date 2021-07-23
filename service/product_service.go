//go:generate mockgen -source product_service.go -destination mock/product_service_mock.go -package mock
package service

import (
	"fmt"
	"mutlicontainer/model"
	"mutlicontainer/repository"

	"github.com/labstack/echo/v4"
)


type productsService struct {
	repository repository.ProductsRepository
}

type ProductsService interface {
	GetDetails(ctx echo.Context) (model.ProductDetails, error)
	CreateProduct(ctx echo.Context) error
}

func NewProductsService(productsRepository repository.ProductsRepository) ProductsService {
	return productsService{productsRepository}
}

func (s productsService) GetDetails(ctx echo.Context) (model.ProductDetails, error) {
	details, err := s.repository.GetDetails(ctx)
	if err != nil {
		return model.ProductDetails{}, err
	}

	return details, nil
}

func (s productsService) CreateProduct(ctx echo.Context)  error {
	err := s.repository.CreateProduct(ctx)
	if err!= nil {
		fmt.Println("something went wrong")
		return err
	}
	return nil
}