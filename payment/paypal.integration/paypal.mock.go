package integration

import (
	"context"
	"github.com/Anton93/test-payment/model"
	"github.com/kelseyhightower/envconfig"
)


type _PaypalApiMock struct {
	config PaypalConfig
}

func NewMockPayService() (model.ExternalServiceIntegration, error) {
	p := &_PaypalApiMock{}

	envconfig.MustProcess(PaypalConfigPrefix, &p.config)

	return p, nil
}

func (p *_PaypalApiMock) GetServiceKey(_ context.Context) string {
	return PaymentServiceName
}

func (p *_PaypalApiMock) GetButtonUrl(ctx context.Context, productIdExt string) (*model.WebPaymentButtonExtService, error) {

	return &model.WebPaymentButtonExtService{
		ID:   PaymentServiceID,
		Name: PaymentServiceName,
		WebPaymentButton: model.WebPaymentButton{
			Url:  "https://www.paypal.com/webapps/mpp/paypal-popup",
			Alt:  "PayPal Logo",
			Logo: "https://www.paypalobjects.com/webstatic/mktg/logo/pp_cc_mark_37x23.jpg",
		},
	}, nil
}
