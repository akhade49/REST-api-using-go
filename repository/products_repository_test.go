package repository

import (
	"database/sql"
	"testing"
	"net/http"
	"net/http/httptest"
	"fmt"
	"os"
	"log"
	"strings"
	"mutlicontainer/model"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/ory/dockertest/v3"
	_ "github.com/lib/pq"
)

func createTable() {
	_, err := dbx.Exec(`
	CREATE TABLE IF NOT EXISTS products 
	(id int,  name varchar(20))
	`)
	if err != nil {
		log.Fatal(err)
	}
}

func deleteTable() {
	_, err := dbx.Exec("DROP TABLE products")
	if err != nil {
		log.Fatal(err)
	}
}
func seedData(){
	dbx.Query("INSERT INTO products (id, name) VALUES (1,'TV');")
	dbx.Query("INSERT INTO products (id, name) VALUES (2,'AC');")
}

//Dockertest approach
var dbx *sql.DB

func TestMain(m *testing.M) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker123: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("postgres", "9.5", []string{"POSTGRES_PASSWORD=password","POSTGRES_DB=root"})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		var err error
		dbx, err = sql.Open("postgres", fmt.Sprintf("postgres://postgres:password@localhost:%s/root?sslmode=disable", resource.GetPort("5432/tcp")))
		if err != nil {
			fmt.Println(err)
			return err
		}
		return dbx.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	code := m.Run()
	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
	os.Exit(code)
}

func TestProductsRepositoryGetDetailsShouldReturnProductDetails(t *testing.T){
	//Arrange
	repo := NewProductsRepository(dbx)
	var ctx echo.Context
	createTable()
	seedData()
	//ACT
	expectedDetails := model.ProductDetails{
		ProductModels: []model.ProductModelResponse{
			{
				Id : 1,
				Name : "TV",
			},
			{
				Id:   2,
				Name: "AC",
			},
		},
	}
	output,_ := repo.GetDetails(ctx)
	//Assert
	assert.Equal(t,expectedDetails,output)
}

func TestProductsRepositoryGetDetailsShouldReturnError(t *testing.T){
	//Arrange
	repo := NewProductsRepository(dbx)
	deleteTable()
	var ctx echo.Context
	expectedDetails := model.ProductDetails{}
	//ACT
	actualDetails,err := repo.GetDetails(ctx)
	//Assert
	assert.Equal(t,expectedDetails,actualDetails)
	assert.Error(t,err)
}

func TestProductsRepositoryCreateProductReturnsNoError(t *testing.T) {
	//Arrange
	repo := NewProductsRepository(dbx)
	createTable()
	userJSON := `{"id":5,"name":"Fridge"}`
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/product", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	//Act
	actualError := repo.CreateProduct(c)
	//Assert
	assert.NoError(t,actualError)
}

func TestProductsRepositoryCreateProductReturnsError(t *testing.T) {
	//Arrange
	repo := NewProductsRepository(dbx)
	deleteTable()
	userJSON := `{"id":5,"name":"Fridge"}`
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/product", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	//Act
	actualError := repo.CreateProduct(c)
	//Assert
	assert.Error(t,actualError)
}

