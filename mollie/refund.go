package mollie

import (
	"time"
	"fmt"
	"net/url"
)

const refundsEndpoint = "refunds"

type RefundRequest struct {
	Amount      map[string]string `json:"amount" validate:"required"`
	Description string            `json:"description" validate:"required"`
}

type RefundResponse struct {
	Resource    string            `json:"resource"`
	ID          string            `json:"id"`
	Amount      map[string]string `json:"amount"`
	Status      string            `json:"status"`
	CreatedAt   time.Time         `json:"createdAt"`
	Description string            `json:"description"`
	PaymentID   string            `json:"paymentId"`
	Links       interface{}       `json:"_links"`
}

func (c *Client) CreateRefund(r *RefundRequest, paymentId string) (*RefundResponse, error) {
	if err := validate.Struct(r); err != nil {
		return nil, err
	}

	refundURL := fmt.Sprintf("%s/%s/%s", paymentsEndpoint, paymentId, refundsEndpoint)

	var p RefundResponse
	err := c.post(refundURL, r, &p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (c *Client) GetRefund(paymentId, refundId string) (*RefundResponse, error) {

	refundURL := fmt.Sprintf("%s/%s/%s/%s", paymentsEndpoint, paymentId, refundsEndpoint, refundId)

	// TODO: Missing the Embed -> don't understand how it works

	var r RefundResponse
	err := c.get(refundURL, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (c *Client) CancelRefund(paymentId, refundId string) (*RefundResponse, error) {
	refundURL := fmt.Sprintf("%s/%s/%s/%s", paymentsEndpoint, paymentId, refundsEndpoint, refundId)

	var r RefundResponse
	err := c.delete(refundURL, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

type RefundOptions struct {
	From  string
	Limit string
}

type RefundsResponse struct {
	Count int `json:"count"`
	Embedded struct {
		Refunds []RefundResponse `json:"refunds"`
	} `json:"_embedded"`
	Links map[string]interface{} `json:"_links"`
}

func (c *Client) ListAllRefunds(options *RefundOptions) (*RefundsResponse, error) {
	refundURL := fmt.Sprintf("%s", refundsEndpoint)

	values := url.Values{}
	if options != nil {
		if options.From != "" {
			values.Set("from", options.From)
		}
		if options.Limit != "" {
			values.Set("limit", options.Limit)
		}
	}
	if query := values.Encode(); query != "" {
		refundURL += "?" + query
	}

	var r RefundsResponse
	err := c.get(refundURL, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (c *Client) ListRefundsOfPayment(options *RefundOptions, paymentId string) (*RefundsResponse, error) {
	refundURL := fmt.Sprintf("%s/%s/%s", paymentsEndpoint, paymentId, refundsEndpoint)

	values := url.Values{}
	if options != nil {
		if options.From != "" {
			values.Set("from", options.From)
		}
		if options.Limit != "" {
			values.Set("limit", options.Limit)
		}
	}
	if query := values.Encode(); query != "" {
		refundURL += "?" + query
	}

	var r RefundsResponse
	err := c.get(refundURL, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
