package integration

import (
	"context"
	"github.com/Anton93/test-payment/model"
	"github.com/kelseyhightower/envconfig"
)

const PaymentServiceID = 1
const PaymentServiceName = "apple_pay"


const ApplePayConfigPrefix = "PAYMENT_APPLE"

type ApplePayConfig struct {

}

type _ApplePayApi struct {
	config ApplePayConfig
}

func NewPayService() (model.ExternalServiceIntegration, error) {
	a := &_ApplePayApi{}
	envconfig.MustProcess(ApplePayConfigPrefix, &a.config)

	// TODO: init
	return a, nil
}

func (a *_ApplePayApi) GetServiceKey(_ context.Context) string {
	return PaymentServiceName
}

func (a *_ApplePayApi) GetButtonUrl(ctx context.Context, productIdExt string) (*model.WebPaymentButtonExtService, error) {
	panic("implement me")
}

