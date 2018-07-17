package mollie

import (
	"fmt"
	"gopkg.in/go-playground/validator.v9"
	"net/url"
)

const endpoint = "payments"

var validate = validator.New()

type PaymentRequest struct {
	Amount           map[string]string `json:"amount" validate:"required"`
	Description      string            `json:"description" validate:"required"`
	RedirectURL      string            `json:"redirectUrl" validate:"required"`
	WebhookURL       string            `json:"webhookUrl" validate:"required"`
	Locale           string            `json:"locale" `
	Method           string            `json:"method"`
	MethodParameters interface{}
	Metadata         interface{}       `json:"metadata"`
	SequenceType     string            `json:"sequenceType"`
	CustomerId       string            `json:"customerId"`
	MandateId        string            `json:"mandateId"`
}

type PaymentResponse struct {
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

func (c *Client) GetPayment(id ID, options *PaymentOptions) (*PaymentResponse, error) {

	paymentURL := fmt.Sprintf("%s%s/%s", c.baseURL, endpoint, id)

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

func (c *Client) CancelPayment(id ID) (*PaymentResponse, error) {
	paymentURL := fmt.Sprintf("%s%s/%s", c.baseURL, endpoint, id)

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
	paymentURL := fmt.Sprintf("%s%s", c.baseURL, endpoint)

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
