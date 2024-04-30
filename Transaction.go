package totalwash

import (
	"cloud.google.com/go/civil"
	"fmt"
	"net/http"
	"net/url"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type TransactionItem struct {
	ProductCode            string  `json:"productcode"`
	ProductDescription     string  `json:"productdescription"`
	Quantity               float64 `json:"quantity"`
	Amount                 float64 `json:"amount"`
	OriginalAmount         float64 `json:"originalamount"`
	LicenceplateNumber     string  `json:"licenceplatenumber"`
	RateDescription        string  `json:"ratedescription"`
	ProductKindId          int64   `json:"productkindid"`
	ProductKindDescription string  `json:"productkinddescription"`
}

type TransactionPayment struct {
	PaymentDescription string  `json:"paymentdescription"`
	Amount             float64 `json:"amount"`
	AmountValue        float64 `json:"amountvalue"`
}

type TransactionDiscount struct {
	CampaignCode        string  `json:"campaigncode"`
	CampaignDescription string  `json:"campaigndescription"`
	Amount              float64 `json:"amount"`
	Value               float64 `json:"value"`
	DiscountKind        string  `json:"discountkind"`
	Code                string  `json:"code"`
	Identification      string  `json:"identification"`
}

type Transaction struct {
	ExternalId               int                   `json:"externalid"`
	ExternalLocationId       int                   `json:"externallocationid"`
	ExternalUserId           int                   `json:"externaluserid"`
	ExternalCustomerId       int                   `json:"externalcustomerid"`
	ExternalCustomerPersonId int                   `json:"externalcustomerpersonid"`
	LocationDescription      string                `json:"locationdescription"`
	UserName                 string                `json:"username"`
	DateTime                 time.Time             `json:"datetime"`
	ReceiptNumber            int                   `json:"receiptnumber"`
	LicenceplateNumber       string                `json:"licenceplatenumber"`
	CardNumber               string                `json:"cardnumber"`
	DeviceId                 string                `json:"deviceid"`
	TransactionItems         []TransactionItem     `json:"transactionitems"`
	TransactionPayments      []TransactionPayment  `json:"transactionpayments"`
	TransactionDiscounts     []TransactionDiscount `json:"transactiondiscounts"`
}

type listTransactionsResponse struct {
	Success      bool          `json:"success"`
	Remark       string        `json:"remark"`
	Transactions []Transaction `json:"transactions"`
}

type ListTransactionsConfig struct {
	Date civil.Date
}

func (service *Service) ListTransactions(cfg *ListTransactionsConfig) (*[]Transaction, *errortools.Error) {
	values := url.Values{}
	values.Add("Date", cfg.Date.String())

	response := listTransactionsResponse{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url(fmt.Sprintf("api/transactions/List?%s", values.Encode())),
		ResponseModel: &response,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &response.Transactions, nil
}
