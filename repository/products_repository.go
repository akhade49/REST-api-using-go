//go:generate mockgen -source products_repository.go -destination mock/products_repository_mock.go -package mock

package repository

import (
	"database/sql"
	"fmt"
	"mutlicontainer/model"

	"github.com/labstack/echo/v4"
)

type productsRepository struct{
	db *sql.DB
}

//constructor  newpersonrepo
func NewProductsRepository(db *sql.DB) ProductsRepository{
	return &productsRepository {
		db: db,
	}
}

type ProductsRepository interface{
		GetDetails(ctx echo.Context) (model.ProductDetails, error)
		CreateProduct(ctx echo.Context) error
}


const (
	selectProductDetailsQuery = `select id,name from products`
	insertProductQuery = `INSERT INTO products (id, name) VALUES ($1,$2)`
)


func (b productsRepository) GetDetails(ctx echo.Context) (model.ProductDetails, error) {
	fmt.Println("Start DB query for product details")
	rows, _ := b.db.Query(selectProductDetailsQuery)

	fmt.Println("Executed db query to get product details")
	defer rows.Close()

	var id int
	var name string
	productModels := make([]model.ProductModelResponse, 0)
	for rows.Next() {
		if err := rows.Scan(&id, &name); err != nil {
			fmt.Println(err.Error())
		}
		productModel := model.ProductModelResponse{
			Id:           		id,
			Name:              name,
		}
		productModels = append(productModels, productModel)
	}


	productDetails := model.ProductDetails{ProductModels: productModels}
	return productDetails, nil
}

func (b productsRepository) CreateProduct(ctx echo.Context) error{
	u := new(model.ProductModelResponse)
	if err := ctx.Bind(u); err != nil {
				return err
		}
	result, err := b.db.Exec(insertProductQuery, u.Id, u.Name)
	if err != nil {
		println("Query not executed")
		return err
	}

	_, err = result.RowsAffected()

	return err
}





