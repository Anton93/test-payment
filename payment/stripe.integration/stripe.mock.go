package integration

import (
	"context"
	"github.com/Anton93/test-payment/model"
	"github.com/kelseyhightower/envconfig"
)

type StripeApiMock struct {
	config StripeConfig
}

func NewMockPayService() (model.ExternalServiceIntegration, error) {
	a := &StripeApiMock{}

	envconfig.MustProcess(StripeConfigPrefix, &a.config)

	// TODO: init
	return a, nil
}

func (a *StripeApiMock) GetServiceKey(_ context.Context) string {
	return PaymentServiceName
}

func (a *StripeApiMock) GetButtonUrl(ctx context.Context, productIdExt string) (*model.WebPaymentButtonExtService, error) {
	// TODO: implement
	return &model.WebPaymentButtonExtService{
		ID:   PaymentServiceID,
		Name: PaymentServiceName,
		WebPaymentButton: model.WebPaymentButton{
			Url:  "https://api.stripe.com/v1/apple_pay/domains",
			Alt:  "Stripe Logo",
			Logo: "https://img.favpng.com/16/21/1/stripe-logo-payment-gateway-business-payment-processor-png-favpng-qbwhc1xBq2R9tKqy0TMYMdgc4.jpg",
		},
	}, nil
}
