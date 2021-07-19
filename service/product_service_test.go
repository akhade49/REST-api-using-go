package service

import (
	// "net/http"
	// "net/http/httptest"

	//"mutlicontainer/model"
	"testing"
	"errors"
	"mutlicontainer/mock"
	"mutlicontainer/model"

	gomock "github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// type BikesServiceTestSuite struct {
// 	mockRepository mock.MockProductsRepository
// 	mockCtrl       gomock.Controller
// 	productsService        ProductsService
// 	context        echo.Context
// }


// func (suite *BikesServiceTestSuite) SetupTest() {
// 	suite.mockCtrl = gomock.NewController(suite.T())
// 	suite.mockRepository = mocks.NewMockBikesRepository(suite.mockCtrl)
// 	suite.service = NewBikesService(suite.mockRepository)
// 	suite.context = context.Background()
// }

func TestGetDetailsShouldReturnProductDetails(t *testing.T) {
	//Arrange
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockRepository := mock.NewMockProductsRepository(controller)
	productservice := NewProductsService(mockRepository)
	var ctx echo.Context
	//Act
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
	mockRepository.EXPECT().GetDetails(ctx).Return(expectedDetails,nil)
	actualDetails,_ := productservice.GetDetails(ctx)
	//Assert
	assert.Equal(t,expectedDetails,actualDetails)
}

func TestGetDetailsShouldReturnError(t *testing.T) {
	//Arrange
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockRepository := mock.NewMockProductsRepository(controller)
	productservice := NewProductsService(mockRepository)
	var ctx echo.Context
	expectedDetails := model.ProductDetails{}
	expectedError := errors.New("Repostitory call failed")
	mockRepository.EXPECT().GetDetails(ctx).Return(expectedDetails,expectedError)
	//Act
	actualDetails,actualError := productservice.GetDetails(ctx)
	//Assert
	assert.Equal(t,expectedDetails,actualDetails)
	assert.Equal(t,expectedError,actualError)
}


func TestCreateProductShouldCreateProduct(t *testing.T) {
	//Arrange
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockRepository := mock.NewMockProductsRepository(controller)
	productservice := NewProductsService(mockRepository)
	var ctx echo.Context
	//Act
	var expectedError error
	mockRepository.EXPECT().CreateProduct(ctx).Return(nil)
	actualError := productservice.CreateProduct(ctx)
	//Assert
	assert.Equal(t,expectedError,actualError)
}

func TestCreateProductShouldReturnError(t *testing.T) {
	//Arrange
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockRepository := mock.NewMockProductsRepository(controller)
	productservice := NewProductsService(mockRepository)
	var ctx echo.Context
	//Act
	expectedError := errors.New("Repostitory call failed")
	mockRepository.EXPECT().CreateProduct(ctx).Return(expectedError)
	actualError := productservice.CreateProduct(ctx)
	//Assert
	assert.Equal(t,expectedError,actualError)
}



// func (suite *ProductsServiceTestSuite) TestGetDetailsShouldReturnProductDetails() {
// 	expectedDetails := model.ProductModelResponse{
// 		Id: 1,
// 		Name: "TV",
// 	}
// 	db := db.GetDBinstance()

// 	productrepo := repository.NewProductsRepository(db)
// 	productservice := service.NewProductsService(productrepo)


// 	suite.mockRepository.EXPECT().GetDetails(suite.context).Return(expectedDetails, nil)
// 	details, err := suite.service.GetDetails(suite.context)

// 	suite.Equal(expectedDetails, details)
// 	suite.Nil(err)
// }

// func TestGetDetailsShouldReturnErrorIfRepositoryReturnsError() {
// 	expectedErr := errors.New("repo error")
// 	suite.mockRepository.EXPECT().GetDetails(suite.context).Return(model.ProductModelResponse{}, expectedErr)

// 	details, err := suite.service.GetDetails(suite.context)

// 	suite.Empty(details)
// 	suite.Equal(expectedErr, err)
// }

