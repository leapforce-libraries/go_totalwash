package totalwash

import (
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type Location struct {
	ExternalID  int64  `json:"externalid"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Street      string `json:"street"`
	Housenumber string `json:"housenumber"`
	PostalCode  string `json:"postalcode"`
	City        string `json:"city"`
}

func (service *Service) GetLocations() (*[]Location, *errortools.Error) {
	locationsResponse := struct {
		Success   bool       `json:"success"`
		Remark    string     `json:"remark"`
		Locations []Location `json:"locations"`
	}{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url("api/transactions/Locations"),
		ResponseModel: &locationsResponse,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &locationsResponse.Locations, nil
}
