package main

import (
	"fmt"
	"mutlicontainer/db"
	_ "github.com/lib/pq"

	 "mutlicontainer/repository"
	 "mutlicontainer/service"

	"github.com/labstack/echo/v4"
)

func main() {
	port := "8080"
	var E = echo.New()

	newDb := db.NewDatabase()
	dbc := newDb.GetDBinstance()

	productrepo := repository.NewProductsRepository(dbc)
	productservice := service.NewProductsService(productrepo)
	//E.GET("/products",productservice.GetDetails)
	E.POST("/product", productservice.CreateProduct)
	// E.PUT("/product/:id", updateproduct)
	// E.DELETE("/product/:id", deleteproduct)
	E.Logger.Print(fmt.Sprintf("Listening on port %s", port))
	address := fmt.Sprintf(":%s", port)
	E.Logger.Fatal(E.Start(address))
}














// type Product struct {
// 	Id   int    `json:"id"`
// 	Name string `json:"name"`
// }
// type Products struct {
// 	Products []Product `json:"products"`
// }

// func createProduct(c echo.Context) error {
// 	u := new(Product)
// 	if err := c.Bind(u); err != nil {
// 		return err
// 	}
// 	sqlStatement := "INSERT INTO products (id,name)VALUES ($1,$2)"
// 	res, err := db.Query(sqlStatement, u.Id, u.Name)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(res)
// 		return c.JSON(http.StatusCreated, u)
// 	}
// 	newproduct := Product{}
// 	newproduct.Id = u.Id
// 	newproduct.Name = u.Name
// 	return c.JSON(http.StatusOK, newproduct)

// }

// func getProduct(c echo.Context) error {
// 	sqlStatement := "SELECT id, name FROM products order by id"
// 	rows, err := db.Query(sqlStatement)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer rows.Close()
// 	result := Products{}

// 	for rows.Next() {
// 		product := Product{}
// 		err2 := rows.Scan(&product.Id, &product.Name)
// 		if err2 != nil {
// 			return err2
// 		}
// 		result.Products = append(result.Products, product)
// 	}
// 	return c.JSON(http.StatusCreated, result)
// }

// func updateproduct(c echo.Context) error {
// 	id := c.Param("id")
// 	u := new(Product)
// 	if err := c.Bind(u); err != nil {
// 		return err
// 	}
// 	sqlStatement := "UPDATE products SET name=$1 WHERE id=$2"
// 	res, err := db.Query(sqlStatement, u.Name, id)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(res)
// 		return c.JSON(http.StatusCreated, u)
// 	}
// 	return c.JSON(http.StatusOK, u)
// }

// func deleteproduct(c echo.Context) error {
// 	id := c.Param("id")
// 	sqlStatement := "DELETE FROM products WHERE id = $1"
// 	res, err := db.Query(sqlStatement, id)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(res)
// 		return c.JSON(http.StatusOK, "Deleted")
// 	}
// 	return c.String(http.StatusOK, id+"Deleted")
// }

// func getDBinstance() sql.DB {
// 	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
// 		"db", 5432, "root", "root", "root")
// 	conn, err := sql.Open("postgres", dsn)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	err = conn.Ping()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return *conn
// }
