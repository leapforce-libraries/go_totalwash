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
	apiName              string = "TotalWash"
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

	tokenSource, e := NewTokenSource(&service)
	if e != nil {
		return nil, e
	}

	oAuth2ServiceConfig := oauth2.ServiceConfig{
		TokenSource: tokenSource,
	}
	oauth2Service, e := oauth2.NewService(&oAuth2ServiceConfig)
	if e != nil {
		return nil, e
	}
	service.oAuth2Service = oauth2Service
	return &service, nil
}

func (service *Service) url(path string) string {
	return fmt.Sprintf("https://%s.%s/integration/%s", service.domain, host, path)
}

func (service *Service) httpRequest(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	errorResponse := ErrorResponse{}
	requestConfig.ErrorModel = &errorResponse

	request, response, e := service.oAuth2Service.HttpRequest(requestConfig)
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

func (service *Service) httpRequestWithoutAccessToken(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	errorResponse := ErrorResponse{}
	requestConfig.ErrorModel = &errorResponse

	request, response, e := service.oAuth2Service.HttpRequestWithoutAccessToken(requestConfig)
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

func (service Service) ApiName() string {
	return apiName
}

func (service Service) ApiKey() string {
	return service.username
}

func (service Service) ApiCallCount() int64 {
	return service.oAuth2Service.ApiCallCount()
}

func (service Service) ApiReset() {
	service.oAuth2Service.ApiReset()
}
