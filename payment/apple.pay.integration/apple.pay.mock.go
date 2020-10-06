package integration

import (
	"context"
	"github.com/Anton93/test-payment/model"
	"github.com/kelseyhightower/envconfig"
)

type _ApplePayApiMock struct {
	config ApplePayConfig
}

func NewMockPayService() (model.ExternalServiceIntegration, error) {
	a := &_ApplePayApiMock{}
	envconfig.MustProcess(ApplePayConfigPrefix, &a.config)

	// TODO: init
	return a, nil
}

func (a *_ApplePayApiMock) GetServiceKey(_ context.Context) string {
	return PaymentServiceName
}

func (a *_ApplePayApiMock) GetButtonUrl(ctx context.Context, productIdExt string) (*model.WebPaymentButtonExtService, error) {

	return &model.WebPaymentButtonExtService{
		ID:   PaymentServiceID,
		Name: PaymentServiceName,
		WebPaymentButton: model.WebPaymentButton{
			Url:  "https://www.apple.com/legal/intellectual-property/trademark/appletmlist.html",
			Alt:  "Apple Logo",
			Logo: "https://developer.apple.com/design/human-interface-guidelines/apple-pay/images/ApplePay_black_yes.png",
		},
	}, nil
}

