package totalwash

import (
	"encoding/json"
	"fmt"
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	oauth2 "github.com/leapforce-libraries/go_oauth2"
)

const (
	Host                 string = "carwash-cms.com"
	AccessTokenGrantType string = "password"
	DateFormat           string = "2006-01-02T15:04:05"
)

// Service stores Service configuration
//
type Service struct {
	domain   string
	username string
	password string
	oAuth2   *oauth2.OAuth2
}

// methods
//
func NewService(domain string, username string, password string) (*Service, *errortools.Error) {
	service := Service{domain: domain, username: username, password: password}

	tokenFunction := func() (*oauth2.Token, *errortools.Error) {
		return service.GetAccessToken()
	}

	config := oauth2.OAuth2Config{
		NewTokenFunction: &tokenFunction,
	}
	service.oAuth2 = oauth2.NewOAuth(config)
	return &service, nil
}

// generic Get method
//
func (service *Service) get(config *oauth2.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	return service.httpRequest(http.MethodGet, config)
}

// generic Post method
//
func (service *Service) post(config *oauth2.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	return service.httpRequest(http.MethodPost, config)
}

// generic Put method
//
func (service *Service) put(config *oauth2.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	return service.httpRequest(http.MethodPut, config)
}

// generic Patch method
//
func (service *Service) patch(config *oauth2.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	return service.httpRequest(http.MethodPatch, config)
}

// generic Delete method
//
func (service *Service) delete(config *oauth2.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	return service.httpRequest(http.MethodDelete, config)
}

func (service *Service) url(path string) string {
	return fmt.Sprintf("https://%s.%s/integration/%s", service.domain, Host, path)
}

func (service *Service) httpRequest(httpMethod string, config *oauth2.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	e := new(errortools.Error)

	errorResponse := ErrorResponse{}

	request, response, e := service.oAuth2.HTTP(httpMethod, config)
	if e != nil {
		if errorResponse.ErrorDescription != "" {
			e.SetMessage(errorResponse.ErrorDescription)
		}

		b, _ := json.Marshal(errorResponse)
		e.SetExtra("error", string(b))

		return nil, nil, e
	}

	return request, response, e
}
