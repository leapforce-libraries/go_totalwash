package totalwash

import (
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type Product struct {
	ExternalLocationID     int64             `json:"externallocationid"`
	LocationDescription    string            `json:"locationdescription"`
	ProductCode            string            `json:"productcode"`
	ProductDescription     string            `json:"productdescription"`
	ProductKindDescription string            `json:"productkinddescription"`
	Categories             []ProductCategory `json:"categories"`
}

type ProductCategory struct {
	ID                  int64  `json:"id"`
	CategoryDescription string `json:"categorydescription"`
}

func (service *Service) GetProducts() (*[]Product, *errortools.Error) {
	productsResponse := struct {
		Success  bool      `json:"success"`
		Remark   string    `json:"remark"`
		Products []Product `json:"products"`
	}{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url("api/transactions/Products"),
		ResponseModel: &productsResponse,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &productsResponse.Products, nil
}
