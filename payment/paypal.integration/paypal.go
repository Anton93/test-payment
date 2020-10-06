package integration

import (
	"context"
	"errors"
	"github.com/Anton93/test-payment/model"
	"github.com/plutov/paypal"
)

const PaymentServiceID = 3
const PaymentServiceName = "paypal"

type _PaypalApi struct {
	client * paypal.Client
}

type PaypalConfig struct {

}

func Create(clientID string, secret string, APIBase string) (model.ExternalServiceIntegration, error) {
	p := &_PaypalApi{}
	var err error
	p.client, err = paypal.NewClient(clientID, secret, APIBase)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (p *_PaypalApi) GetButtonUrl(ctx context.Context, productIdExt string) (*model.WebPaymentButtonExtService, error) {
	if p.client == nil {
		return nil, errors.New("paypal: client not init")
	}

	// TODO: paypal client request
	// p.clinet.

	return  &model.WebPaymentButtonExtService{
		ID:               PaymentServiceID,
		Name:             PaymentServiceName,
		WebPaymentButton: model.WebPaymentButton{
			Url:  "https://www.paypal.com/webapps/mpp/paypal-popup",
			Alt:  "PayPal Logo",
			Logo: "https://www.paypalobjects.com/webstatic/mktg/logo/pp_cc_mark_37x23.jpg",
		},
	}, nil
}


