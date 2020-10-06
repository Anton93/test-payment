package integration

import (
	"context"
	"github.com/Anton93/test-payment/model"
	"github.com/kelseyhightower/envconfig"
)

type GooglePayApiMock struct {
	config GooglePayConfig
}

func NewMockPayService() (model.ExternalServiceIntegration, error) {
	a := &GooglePayApiMock{}

	envconfig.MustProcess(GooglePayConfigPrefix, &a.config)

	// TODO: init
	return a, nil
}

func (g *GooglePayApiMock) GetServiceKey(_ context.Context) string {
	return PaymentServiceName
}

func (g *GooglePayApiMock) GetButtonUrl(ctx context.Context, productIdExt string) (*model.WebPaymentButtonExtService, error) {
	return &model.WebPaymentButtonExtService{
		ID:   PaymentServiceID,
		Name: PaymentServiceName,
		WebPaymentButton: model.WebPaymentButton{
			Url:  "https://www.google.com/url?sa=t&rct=j&q=&esrc=s&source=web&cd=&cad=rja&uact=8&ved=2ahUKEwiQ2KjGup_sAhVPEncKHRdKBHoQFjAAegQIBRAC&url=https%3A%2F%2Fpay.google.com%2Fsend%2Fhome&usg=AOvVaw2SmiuRfJEKuJh-b8meR7nu",
			Alt:  "Google Pay Logo",
			Logo: "https://developers.google.com/pay/api/images/brand-guidelines/google-pay-mark.png",
		},
	}, nil
}

