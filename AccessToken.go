package totalwash

import (
	"encoding/json"

	errortools "github.com/leapforce-libraries/go_errortools"
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

	skipAccessToken := true
	xWWWFormURLEncoded := true

	requestConfig := oauth2.RequestConfig{
		URL:                service.url("token"),
		BodyModel:          body,
		ResponseModel:      &accessToken,
		SkipAccessToken:    &skipAccessToken,
		XWWWFormURLEncoded: &xWWWFormURLEncoded,
	}

	_, _, e := service.post(&requestConfig)
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
