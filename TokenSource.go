package totalwash

import (
	"encoding/json"
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	go_token "github.com/leapforce-libraries/go_oauth2/token"
)

type TokenSource struct {
	token   *go_token.Token
	service *Service
}

func NewTokenSource(service *Service) (*TokenSource, *errortools.Error) {
	if service == nil {
		return nil, errortools.ErrorMessage("Service is a nil pointer")
	}

	return &TokenSource{
		service: service,
	}, nil
}

func (t *TokenSource) Token() *go_token.Token {
	return t.token
}

func (t *TokenSource) NewToken() (*go_token.Token, *errortools.Error) {

	body := struct {
		Username  string `json:"username"`
		Password  string `json:"password"`
		GrantType string `json:"grant_type"`
	}{
		t.service.username,
		t.service.password,
		accessTokenGrantType,
	}

	accessToken := AccessToken{}

	xWWWFormURLEncoded := true

	requestConfig := go_http.RequestConfig{
		Method:             http.MethodPost,
		URL:                t.service.url("token"),
		BodyModel:          body,
		ResponseModel:      &accessToken,
		XWWWFormURLEncoded: &xWWWFormURLEncoded,
	}

	_, _, e := t.service.httpRequestWithoutAccessToken(&requestConfig)
	if e != nil {
		return nil, e
	}

	expiresIn, _ := json.Marshal(accessToken.ExpiresIn)
	expiresInJson := json.RawMessage(expiresIn)

	return &go_token.Token{
		AccessToken: &accessToken.AccessToken,
		ExpiresIn:   &expiresInJson,
		TokenType:   &accessToken.TokenType,
	}, nil
}

func (t *TokenSource) SetToken(token *go_token.Token, save bool) *errortools.Error {
	t.token = token

	if !save {
		return nil
	}

	return t.SaveToken()
}

func (t *TokenSource) RetrieveToken() *errortools.Error {
	return nil
}

func (t *TokenSource) SaveToken() *errortools.Error {
	return nil
}
