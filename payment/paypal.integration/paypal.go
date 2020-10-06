package integration

import (
	"context"
	"errors"
	"github.com/Anton93/test-payment/model"
	"github.com/kelseyhightower/envconfig"
	"github.com/plutov/paypal/v3"
)

const PaymentServiceID = 3
const PaymentServiceName = "paypal"

const PaypalConfigPrefix = "PAYMENT_PAYPAL"

type PaypalConfig struct {
	ClientID string `envconfig:"CLIENT_ID"`
	Secret   string `envconfig:"SECRET"`
	APIBase  string `envconfig:"BASE"`
}

type _PaypalApi struct {
	config PaypalConfig
	client *paypal.Client
}

func NewPayService() (model.ExternalServiceIntegration, error) {
	p := &_PaypalApi{}

	envconfig.MustProcess(PaypalConfigPrefix, &p.config)

	var err error
	p.client, err = paypal.NewClient(p.config.ClientID, p.config.Secret, p.config.APIBase)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (p *_PaypalApi) GetServiceKey(_ context.Context) string {
	return PaymentServiceName
}

func (p *_PaypalApi) GetButtonUrl(ctx context.Context, productIdExt string) (*model.WebPaymentButtonExtService, error) {
	if p.client == nil {
		return nil, errors.New("paypal: client not init")
	}

	// TODO: paypal client request
	// p.clinet.

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
