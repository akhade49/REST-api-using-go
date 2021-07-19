package model


type ProductDetails struct {
	ProductModels []ProductModelResponse `json:"productModels"`
}

type ProductModelResponse struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
}