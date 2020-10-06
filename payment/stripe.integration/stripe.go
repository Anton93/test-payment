package integration

import (
	"context"
	"github.com/Anton93/test-payment/model"
	"github.com/kelseyhightower/envconfig"
)

const PaymentServiceID = 4
const PaymentServiceName = "stripe"

const StripeConfigPrefix = "PAYMENT_STRIPE"

type StripeConfig struct {
}

type StripeApi struct {
	config StripeConfig
}

func NewPayService() (model.ExternalServiceIntegration, error) {
	a := &StripeApi{}

	envconfig.MustProcess(StripeConfigPrefix, &a.config)

	// TODO: init
	return a, nil
}

func (a *StripeApi) GetServiceKey(_ context.Context) string {
	return PaymentServiceName
}

func (a *StripeApi) GetButtonUrl(ctx context.Context, productIdExt string) (*model.WebPaymentButtonExtService, error) {
	// TODO: implement
	panic("implement me")
}
