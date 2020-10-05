package payment

import "context"

// WebIntegrationRequest.
type WebIntegrationsRequest struct {
	ProductID string `json:"product_id"`
}

type WebIntegration struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
	Alt  string `json:"alt"`
	Logo string `json:"logo"`
}

type WebIntegrationsResponse struct {
	Button []*WebIntegration `json:"button"`
}

type Service interface {
	WebIntegrations(ctx context.Context, request *WebIntegrationsRequest) (response *WebIntegrationsResponse, err error)
}


type ExternalServiceIntegration interface {
	
}
