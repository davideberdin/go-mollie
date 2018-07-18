package mollie

import (
	"fmt"
	"gopkg.in/go-playground/validator.v9"
	"net/url"
	"time"
)

const endpoint = "payments"

var validate = validator.New()

type Address struct {
	StreetAndNumber string `json:"streetAndNumber,omitempty"`
	PostalCode      string `json:"postalCode,omitempty"`
	City            string `json:"city,omitempty"`
	Region          string `json:"region,omitempty"`
	Country         string `json:"country,omitempty"`
}

type PaymentRequest struct {
	Amount            map[string]string `json:"amount" validate:"required"`
	Description       string            `json:"description" validate:"required"`
	RedirectURL       string            `json:"redirectUrl" validate:"required"`
	WebhookURL        string            `json:"webhookUrl" validate:"required"`
	Method            string            `json:"method" validate:"required"`
	Locale            string            `json:"locale,omitempty"`
	Metadata          interface{}       `json:"metadata,omitempty"`
	SequenceType      string            `json:"sequenceType,omitempty"`
	CustomerId        string            `json:"customerId,omitempty"`
	MandateId         string            `json:"mandateId,omitempty"`
	BillingEmail      string            `json:"billingEmail,omitempty"`
	DueDate           string            `json:"dueDate,omitempty"`
	BillingAddress    Address           `json:"billingAddress,omitempty"`
	ShippingAddress   Address           `json:"shippingAddress,omitempty"`
	VoucherNumber     string            `json:"voucherNumber,omitempty"`
	VoucherPin        string            `json:"voucherPin,omitempty"`
	Issuer            string            `json:"issuer,omitempty"`
	CustomerReference string            `json:"customerReference,omitempty"`
}

type PaymentResponse struct {
	Resource         string            `json:"resource"`
	ID               string            `json:"id"`
	Mode             string            `json:"mode"`
	CreatedAt        time.Time         `json:"createdAt"`
	Amount           map[string]string `json:"amount"`
	Description      string            `json:"description"`
	Method           interface{}       `json:"method"`
	Metadata         interface{}       `json:"metadata"`
	Status           string            `json:"status"`
	IsCancelable     bool              `json:"isCancelable"`
	ExpiresAt        time.Time         `json:"expiresAt"`
	Details          interface{}       `json:"details"`
	ProfileID        string            `json:"profileId"`
	SettlementAmount interface{}       `json:"settlementAmount"`
	SequenceType     string            `json:"sequenceType"`
	RedirectURL      string            `json:"redirectUrl"`
	WebhookURL       string            `json:"webhookUrl"`
	Links            interface{}       `json:"_links"`
}

func (c *Client) CreatePayment(p *PaymentRequest) (*PaymentResponse, error) {

	if err := validate.Struct(p); err != nil {
		return nil, err
	}

	paymentURL := fmt.Sprintf("%s%s", c.baseURL, endpoint)

	var r PaymentResponse
	err := c.post(paymentURL, p, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

type PaymentOptions struct {
	IncludeQrCode    string // value: "details.qrCode"
	EmbedRefunds     string // value: "refunds"
	EmbedChargebacks string // value: "chargebacks"
	From             string // value: "from"
	Limit            string // value: "limit"
}

func (c *Client) GetPayment(id string, options *PaymentOptions) (*PaymentResponse, error) {

	paymentURL := fmt.Sprintf("%s/%s", endpoint, id)

	values := url.Values{}
	if options != nil {
		if options.IncludeQrCode != "" {
			values.Set("include", options.IncludeQrCode)
		}
		if options.EmbedRefunds != "" {
			values.Set("embed", options.EmbedRefunds)
		}
		if options.EmbedChargebacks != "" {
			values.Set("embed", options.EmbedChargebacks)
		}
	}
	if query := values.Encode(); query != "" {
		paymentURL += "?" + query
	}

	var r PaymentResponse
	err := c.get(paymentURL, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (c *Client) CancelPayment(id string) (*PaymentResponse, error) {
	paymentURL := fmt.Sprintf("%s/%s", endpoint, id)

	var r PaymentResponse
	err := c.get(paymentURL, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

type PaymentsResponse struct {
	Count    uint                         `json:"count"`
	Embedded map[string][]PaymentResponse `json:"_embedded"`
	Links    map[string]string            `json:"_links"`
}

func (c *Client) ListPayments(options *PaymentOptions) (*PaymentsResponse, error) {
	paymentURL := fmt.Sprintf("%s", endpoint)

	values := url.Values{}
	if options != nil {
		if options.IncludeQrCode != "" {
			values.Set("include", options.IncludeQrCode)
		}
		if options.EmbedRefunds != "" {
			values.Set("embed", options.EmbedRefunds)
		}
		if options.EmbedChargebacks != "" {
			values.Set("embed", options.EmbedChargebacks)
		}
	}
	if query := values.Encode(); query != "" {
		paymentURL += "?" + query
	}

	var r PaymentsResponse
	err := c.get(paymentURL, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
