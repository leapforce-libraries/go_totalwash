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
func (service *Service) get(urlPath string, responseModel interface{}) (*http.Request, *http.Response, *errortools.Error) {
	return service.httpRequest(http.MethodGet, urlPath, nil, responseModel)
}

// generic Post method
//
func (service *Service) post(urlPath string, bodyModel interface{}, responseModel interface{}) (*http.Request, *http.Response, *errortools.Error) {
	return service.httpRequest(http.MethodPost, urlPath, bodyModel, responseModel)
}

// generic Put method
//
func (service *Service) put(urlPath string, bodyModel interface{}, responseModel interface{}) (*http.Request, *http.Response, *errortools.Error) {
	return service.httpRequest(http.MethodPut, urlPath, bodyModel, responseModel)
}

// generic Patch method
//
func (service *Service) patch(urlPath string, bodyModel interface{}, responseModel interface{}) (*http.Request, *http.Response, *errortools.Error) {
	return service.httpRequest(http.MethodPatch, urlPath, bodyModel, responseModel)
}

// generic Delete method
//
func (service *Service) delete(urlPath string, bodyModel interface{}, responseModel interface{}) (*http.Request, *http.Response, *errortools.Error) {
	return service.httpRequest(http.MethodDelete, urlPath, bodyModel, responseModel)
}

func (service *Service) httpRequest(httpMethod string, urlPath string, bodyModel interface{}, responseModel interface{}) (*http.Request, *http.Response, *errortools.Error) {
	url := fmt.Sprintf("https://%s.%s/integration/%s", service.domain, Host, urlPath)
	fmt.Println(url)

	e := new(errortools.Error)

	errorResponse := ErrorResponse{}

	request, response, e := service.oAuth2.HTTP(httpMethod, url, bodyModel, responseModel, &errorResponse)

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
