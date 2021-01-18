package totalwash

import (
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
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

	expirationDate := new(string)
	expirationDate = nil

	if expiry != nil {
		expirationDate_ := expiry.Format(DateFormat)
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

	_, _, e := service.post("api/couponcodes/Generate", body, &generatedCouponCodes)
	if e != nil {
		return nil, e
	}

	return &generatedCouponCodes, nil
}
