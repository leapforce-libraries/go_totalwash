package totalwash

import (
	"cloud.google.com/go/civil"
	"fmt"
	"net/http"
	"net/url"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	t "github.com/leapforce-libraries/go_totalwash/types"
)

type Customer struct {
	ExternalId        int               `json:"externalid"`
	ExternalPersonId  int               `json:"externalpersonid"`
	Initials          string            `json:"initials"`
	FirstName         string            `json:"firstname"`
	MiddleName        string            `json:"middlename"`
	Name              string            `json:"name"`
	Street            string            `json:"street"`
	HouseNumber       string            `json:"housenumber"`
	PostalCode        string            `json:"postalcode"`
	City              string            `json:"city"`
	BirthDate         *civil.Date       `json:"birthdate"`
	Email             string            `json:"email"`
	Marketing         bool              `json:"marketing"`
	DateModified      time.Time         `json:"datemodified"`
	Group             string            `json:"group"`
	LastVisit         *t.DateTimeString `json:"lastvisit"`
	LastVisitLocation string            `json:"lastvisitlocation"`
	LoyaltyPoints     float64           `json:"loyaltypoints"`
	Balance           float64           `json:"balance"`
	Subscription      bool              `json:"subscription"`
}

type listCustomersResponse struct {
	Success   bool       `json:"success"`
	Remark    string     `json:"remark"`
	Customers []Customer `json:"customers"`
}

type ListCustomersConfig struct {
	DateTimeModifiedStart time.Time
}

func (service *Service) ListCustomers(cfg *ListCustomersConfig) (*[]Customer, *errortools.Error) {
	values := url.Values{}
	values.Add("DateTimeModifiedStart", cfg.DateTimeModifiedStart.Format("2006-01-02T15:04:05"))

	response := listCustomersResponse{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url(fmt.Sprintf("api/customers/List?%s", values.Encode())),
		ResponseModel: &response,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &response.Customers, nil
}
