package model

import "context"

// WebIntegrationRequest.
type GetWebPaymentButtonsRequest struct {
	ProductID string `json:"product_id"`
}

type WebPaymentButton struct {
	Url  string `json:"url"`
	Alt  string `json:"alt"`
	Logo string `json:"logo"`
}

type WebPaymentButtonExtService struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	WebPaymentButton
}

type GetWebPaymentButtonsResponse struct {
	Button []*WebPaymentButtonExtService `json:"button"`
}

type PaymentService interface {
	GetWebPaymentButtons(ctx context.Context, request *GetWebPaymentButtonsRequest) (response *GetWebPaymentButtonsResponse, err error)
}

type ExternalServiceIntegration interface {
	GetButtonUrl(ctx context.Context, productIdExt string) (*WebPaymentButtonExtService, error)
}
