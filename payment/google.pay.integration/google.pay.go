package integration

import (
	"context"
	"github.com/Anton93/test-payment/model"
	"github.com/kelseyhightower/envconfig"
)

const PaymentServiceID = 2
const PaymentServiceName = "google_pay"

const GooglePayConfigPrefix = "PAYMENT_GOOGLE"

type GooglePayConfig struct {

}

type GooglePayApi struct {
	config GooglePayConfig
}

func NewPayService() (model.ExternalServiceIntegration, error) {
	a := &GooglePayApi{}

	envconfig.MustProcess(GooglePayConfigPrefix, &a.config)

	// TODO: init
	return a, nil
}

func (g *GooglePayApi) GetServiceKey(_ context.Context) string {
	return PaymentServiceName
}

func (g *GooglePayApi) GetButtonUrl(ctx context.Context, productIdExt string) (*model.WebPaymentButtonExtService, error) {
	panic("implement me")
}

