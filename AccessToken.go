package totalwash

import (
	"encoding/json"
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	oauth2 "github.com/leapforce-libraries/go_oauth2"
)

// Token stures Token object
//
type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func (service *Service) GetAccessToken() (*oauth2.Token, *errortools.Error) {
	body := struct {
		Username  string `json:"username"`
		Password  string `json:"password"`
		GrantType string `json:"grant_type"`
	}{
		service.username,
		service.password,
		AccessTokenGrantType,
	}

	accessToken := AccessToken{}

	xWWWFormURLEncoded := true

	requestConfig := go_http.RequestConfig{
		URL:                service.url("token"),
		BodyModel:          body,
		ResponseModel:      &accessToken,
		XWWWFormURLEncoded: &xWWWFormURLEncoded,
	}

	_, _, e := service.httpRequest(http.MethodPost, &requestConfig, true)
	if e != nil {
		return nil, e
	}

	expiresIn, _ := json.Marshal(accessToken.ExpiresIn)
	expiresInJson := json.RawMessage(expiresIn)

	token := oauth2.Token{
		AccessToken: &accessToken.AccessToken,
		ExpiresIn:   &expiresInJson,
		TokenType:   &accessToken.TokenType,
	}

	return &token, nil
}
