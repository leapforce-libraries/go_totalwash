package totalwash

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type GeneratedCouponCodes struct {
	Success bool     `json:"success"`
	Remark  string   `json:"remark"`
	Codes   []string `json:"codes"`
}

func (service *Service) GenerateCouponCodes(campaignID string, count uint, expiry *time.Time) (*GeneratedCouponCodes, *errortools.Error) {
	if count == 0 {
		return nil, nil
	}

	var expirationDate *string = nil

	if expiry != nil {
		expirationDate_ := expiry.Format(dateLayout)
		expirationDate = &expirationDate_
	}

	body := struct {
		CampaignID     string  `json:"campaignid"`
		Count          uint    `json:"count"`
		ExpirationDate *string `json:"expirationdate,omitempty"`
	}{
		campaignID,
		count,
		expirationDate,
	}
	generatedCouponCodes := GeneratedCouponCodes{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodPost,
		Url:           service.url("api/couponcodes/Generate"),
		BodyModel:     body,
		ResponseModel: &generatedCouponCodes,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &generatedCouponCodes, nil
}

type CouponCode struct {
	Id           string
	CampaignId   string
	Code         string    `json:"code"`
	Redeemed     bool      `json:"redeemed"`
	DateModified time.Time `json:"timestamp"`
}

type listCouponCodesResponse struct {
	Success     bool         `json:"success"`
	Remark      string       `json:"remark"`
	CouponCodes []CouponCode `json:"codes"`
}

type ListCouponCodesConfig struct {
	CampaignId string
}

func (service *Service) ListCouponCodes(cfg *ListCouponCodesConfig) (*[]CouponCode, *errortools.Error) {
	values := url.Values{}
	values.Add("CampaignId", cfg.CampaignId)

	response := listCouponCodesResponse{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url(fmt.Sprintf("api/couponcodes/List?%s", values.Encode())),
		ResponseModel: &response,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &response.CouponCodes, nil
}
