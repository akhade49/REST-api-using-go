package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	// 	"time"
	// //	"io/ioutil"
	// 	"encoding/json"
	// 	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"github.com/labstack/echo"
)

type Product struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type Products struct {
	Products []Product `json:"employees"`
}


func main() {
	port := "8080"
	var E = echo.New()
	
	getDBinstance()
	E.GET("/", helloworld)
	E.GET("/product", product)
	E.Logger.Print(fmt.Sprintf("Listening on port %s", port))
	E.Logger.Fatal(E.Start(fmt.Sprintf("localhost:%s", port)))


	// without fecho
	//router := mux.NewRouter()
	// router.HandleFunc("/posts", getProducts).Methods("GET")
  	// //router.HandleFunc("/post", createProduct).Methods("POST")
	//  srv := &http.Server{
	// 	Handler:      router,
	// 	Addr:         ":8080",
	// 	ReadTimeout:  10 * time.Second,
	// 	WriteTimeout: 10 * time.Second,
	// }
	// if err := srv.ListenAndServe(); err != nil {
	// 	log.Fatal(err)
	// }
}


func getDBinstance() sql.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"db", 5432, "root", "root", "root")
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	err = conn.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("db is used=========")
	//defer conn.Close()
	return *conn
}


// without echo
// func getProducts(w http.ResponseWriter, request *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	db := getDBinstance()
//   var products []Product
//   result, err := db.Query("SELECT id, name from products")
//   if err != nil {
//     panic(err.Error())
//   }
//   defer result.Close()
//   for result.Next() {
//     var product Product
//     err := result.Scan(&product.Id, &product.Name)
//     if err != nil {
//       panic(err.Error())
//     }
//     products = append(products, product)
//   }
//   json.NewEncoder(w).Encode(products)
// }





// helloworld and product is of echo

func helloworld(c echo.Context) error {
	// db := getDBinstance()
	// u := new(Product)
	// if err := c.Bind(u); err != nil {
	// 	return err
	// }
	// sqlStatement := "INSERT INTO products (name)VALUES ($1)"
	// res, err := db.Query(sqlStatement, u.Name)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(res)
	// 	return c.JSON(http.StatusCreated, u)
	// }
	return c.String(http.StatusOK, "ok")

}

func product(c echo.Context) error {
	db := getDBinstance()
	var id int
	var name string
	fmt.Println("hi")
	userSql := "SELECT id, name FROM products WHERE id = $1"
	err := db.QueryRow(userSql, 1).Scan(&id, &name)
	if err != nil {
		log.Fatal("Failed to execute query: ", err)
	}
	user := Product{Id: id, Name: name}
	fmt.Printf("Hi %s, welcome back!\n", user.Name)

	return c.String(http.StatusOK, "ok")
	//w.Write([]byte(user.Name))
}

// without echo
func testdb(w http.ResponseWriter, request *http.Request) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"db", 5432, "root", "root", "root")
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connection to pg herr.....")
	err = conn.Ping()
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte("success"))
}


