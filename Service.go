package totalwash

import (
	"encoding/json"
	"fmt"
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	oauth2 "github.com/leapforce-libraries/go_oauth2"
)

const (
	apiName              string = "totalwash"
	host                 string = "carwash-cms.com"
	accessTokenGrantType string = "password"
	dateLayout           string = "2006-01-02T15:04:05"
)

type ServiceConfig struct {
	Domain   string
	Username string
	Password string
}

type Service struct {
	domain        string
	username      string
	password      string
	oAuth2Service *oauth2.Service
}

// methods
//
func NewService(serviceConfig *ServiceConfig) (*Service, *errortools.Error) {
	if serviceConfig == nil {
		return nil, errortools.ErrorMessage("ServiceConfig must not be a nil pointer")
	}

	service := Service{
		domain:   serviceConfig.Domain,
		username: serviceConfig.Username,
		password: serviceConfig.Password,
	}

	tokenFunction := func() (*oauth2.Token, *errortools.Error) {
		return service.GetAccessToken()
	}

	oAuth2ServiceConfig := oauth2.ServiceConfig{
		NewTokenFunction: &tokenFunction,
	}
	oauth2Service, e := oauth2.NewService(&oAuth2ServiceConfig)
	if e != nil {
		return nil, e
	}
	service.oAuth2Service = oauth2Service
	return &service, nil
}

// generic Get method
//
func (service *Service) get(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	return service.oAuth2Service.Get(requestConfig)
}

// generic Post method
//
func (service *Service) post(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	return service.oAuth2Service.Post(requestConfig)
}

// generic Put method
//
func (service *Service) put(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	return service.oAuth2Service.Put(requestConfig)
}

// generic Patch method
//
func (service *Service) patch(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	return service.oAuth2Service.Patch(requestConfig)
}

// generic Delete method
//
func (service *Service) delete(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	return service.oAuth2Service.Delete(requestConfig)
}

func (service *Service) url(path string) string {
	return fmt.Sprintf("https://%s.%s/integration/%s", service.domain, host, path)
}

func (service *Service) httpRequest(httpMethod string, requestConfig *go_http.RequestConfig, skipAccessToken bool) (*http.Request, *http.Response, *errortools.Error) {
	e := new(errortools.Error)

	errorResponse := ErrorResponse{}
	requestConfig.ErrorModel = &errorResponse

	request, response, e := service.oAuth2Service.HTTPRequest(httpMethod, requestConfig, skipAccessToken)
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

func (service Service) APIName() string {
	return apiName
}

func (service Service) APIKey() string {
	return service.username
}

func (service Service) APICallCount() int64 {
	return service.oAuth2Service.APICallCount()
}
